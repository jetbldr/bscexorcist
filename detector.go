// Package bscexorcist provides sandwich attack detection for BSC transaction bundles.
package bscexorcist

import (
	"fmt"
	"math/big"

	"github.com/48Club/bscexorcist/protocols"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// sandwichTolPercent is the allowed mismatch, in percent of the larger side, between
// the asset acquired by the front-run and the asset disposed by the back-run. A genuine
// attacker unwinds almost exactly what it acquired (residual ~0; real captures sit
// within a few percent), so a small tolerance separates true sandwiches from unrelated
// trades that merely share a pool and direction.
const sandwichTolPercent = 15

// swapLeg is one transaction's net activity on a single pool: its net swap direction and
// the trader's signed net receipt of token0/token1 (positive = received).
type swapLeg struct {
	isToken0To1 bool
	dT0, dT1    *big.Int
	reliable    bool
}

// DetectSandwichForBundle analyzes a bundle of transaction logs to identify potential sandwich attacks.
// Returns an error if a sandwich pattern is detected in any pool within the bundle.
func DetectSandwichForBundle(bundleLogs [][]*types.Log) error {
	if len(bundleLogs) < 3 {
		return nil
	}

	poolLegs := make(map[common.Address][]swapLeg)
	for _, txLogs := range bundleLogs {
		// Collapse all swaps a single transaction makes on the same pool into one
		// net-direction leg. A transaction that routes through the same pool several
		// times (an aggregator split or a multi-hop arbitrage) would otherwise
		// contribute multiple same-direction legs, fabricating a "Buy-Buy" / "Sell-Sell"
		// prefix. A genuine sandwich keeps its front-run and back-run in separate
		// transactions (the victim's transaction sits between them), so collapsing per
		// transaction never hides a real attack.
		for poolID, leg := range netLegsPerPool(txLogs) {
			poolLegs[poolID] = append(poolLegs[poolID], leg)
		}
	}

	for pool, legs := range poolLegs {
		if isSandwich(legs) {
			return fmt.Errorf("sandwich attack detected on pool: %s", pool.Hex())
		}
	}

	return nil
}

// netLegsPerPool reduces all swaps in a single transaction to one leg per pool: the net
// swap direction plus the trader's signed net token0/token1 receipt. For each pool it
// tallies token0->token1 versus the reverse and reports the majority direction; pools
// with an equal split carry no net directional signal and are omitted.
//
// The signed receipts are reconstructed from Abs(AmountIn)/Abs(AmountOut), which are true
// input/output magnitudes for the protocols that mark themselves reliable. A leg is
// reliable only if every swap on the pool is reliable (a pool is single-protocol, so this
// is all-or-nothing in practice).
func netLegsPerPool(txLogs []*types.Log) map[common.Address]swapLeg {
	type tally struct {
		token0To1, token1To0 int
		dT0, dT1             *big.Int
		reliable             bool
	}
	counts := make(map[common.Address]*tally)
	for _, swap := range protocols.ParseSwapEvents(txLogs) {
		poolID := swap.PairID()
		t := counts[poolID]
		if t == nil {
			t = &tally{dT0: new(big.Int), dT1: new(big.Int), reliable: true}
			counts[poolID] = t
		}
		in := new(big.Int).Abs(swap.AmountIn())
		out := new(big.Int).Abs(swap.AmountOut())
		if swap.IsToken0To1() {
			// trader pays token0 (input), receives token1 (output)
			t.token0To1++
			t.dT0.Sub(t.dT0, in)
			t.dT1.Add(t.dT1, out)
		} else {
			// trader pays token1, receives token0
			t.token1To0++
			t.dT0.Add(t.dT0, out)
			t.dT1.Sub(t.dT1, in)
		}
		t.reliable = t.reliable && swap.AmountsReliable()
	}

	legs := make(map[common.Address]swapLeg, len(counts))
	for poolID, t := range counts {
		switch {
		case t.token0To1 > t.token1To0:
			legs[poolID] = swapLeg{isToken0To1: true, dT0: t.dT0, dT1: t.dT1, reliable: t.reliable}
		case t.token1To0 > t.token0To1:
			legs[poolID] = swapLeg{isToken0To1: false, dT0: t.dT0, dT1: t.dT1, reliable: t.reliable}
		}
	}
	return legs
}

// isSandwich reports whether a pool's per-transaction legs (in execution order) form a
// sandwich, distinguishing a real attack from unrelated trades that merely share a pool
// and direction (e.g. one searcher's incidental swap plus another's multi-leg arbitrage).
//
// For reliable-amount pools the verdict is asset conservation around a bracketed victim: a
// front-run whose acquired asset is sold back by a later back-run, with the victim trading
// the same direction in between. This is sender-agnostic — the bought amount equals the
// sold-back amount regardless of which addresses sign the legs — so it still catches
// sandwiches spread across multiple addresses. A conserving bracket is a front leg and
// victim in one direction followed by a back leg in the other, so it already entails the
// Buy-Buy-Sell / Sell-Sell-Buy shape; the separate directional pre-check is therefore
// redundant here and is skipped (it would only ever agree).
//
// DODO/FourMeme amounts cannot be reconciled (indexed by token, or zero), so there is
// nothing to conserve; those pools fall back to flagging on the swap directions alone,
// preserving prior behavior.
func isSandwich(legs []swapLeg) bool {
	if len(legs) < 3 {
		return false
	}

	for _, l := range legs {
		if !l.reliable {
			dirs := make([]bool, len(legs))
			for i, leg := range legs {
				dirs[i] = leg.isToken0To1
			}
			return hasDirectionalPattern(dirs)
		}
	}

	// A buy side of token0->token1 makes token1 the asset (Buy-Buy-Sell); the reverse
	// makes token0 the asset (Sell-Sell-Buy). Try both orientations.
	return conserves(legs, true) || conserves(legs, false)
}

// conserves checks the asset-conservation condition for one orientation. buySide selects
// the front/victim direction; the asset token is token1 when buySide is true, else token0.
//
// For each candidate victim leg it builds the front-run as the single contiguous run of
// same-direction legs nearest the victim (walking backward), and the back-run as the
// contiguous run of opposite-direction legs nearest it (walking forward), then flags when
// some front run total matches some back run total within tolerance. Accumulating inward
// from the victim and trying each partial sum absorbs a front-run or back-run split across
// several adjacent transactions, while a nearer partial sum winning excludes an incidental
// same-direction trade that sits further out.
//
// The run is genuinely contiguous: an opposite-direction (or folded, see below) leg that
// falls *inside* the run ends it, so legs on opposite sides of an interruption are never
// stitched into one total — e.g. buy 40, sell 10, buy 60 yields front candidates 40 and
// 60, never 100. An unrelated trade may, however, sit between the run and the victim (a
// leading gap is skipped): real bundles do interleave an unrelated swap between the
// front-run and the victim, yet keep the front-run itself in consecutive transactions.
//
// The signed net asset change (dT1 or dT0), not its absolute value, decides whether a leg
// belongs to a run: a front-run leg must have actually acquired the asset (delta > 0) and
// a back-run leg must have actually disposed of it (delta < 0). A transaction whose
// majority swap direction disagrees with its net asset flow — e.g. two tiny buys outvoting
// one large sell so the leg folds to "buy" while it net-sold — does not qualify and ends
// the run just like an opposite-direction leg.
func conserves(legs []swapLeg, buySide bool) bool {
	assetDelta := func(l swapLeg) *big.Int {
		if buySide {
			return l.dT1
		}
		return l.dT0
	}

	for vi := range legs {
		// A sandwiched victim trades the same direction as the front-run and net-acquires
		// the asset (it ends up holding more of it). Skip legs that fail either: a wrong
		// direction, or a net asset flow that contradicts the folded majority direction.
		if legs[vi].isToken0To1 != buySide || assetDelta(legs[vi]).Sign() <= 0 {
			continue
		}
		// front-run partial sums: the contiguous run of legs that acquired the asset in the
		// front direction, nearest the victim first (ascending). A leading gap of unrelated
		// legs is skipped; the first non-qualifying leg after the run starts ends it.
		var front []*big.Int
		acc := new(big.Int)
		for i := vi - 1; i >= 0; i-- {
			if legs[i].isToken0To1 == buySide && assetDelta(legs[i]).Sign() > 0 {
				acc = new(big.Int).Add(acc, assetDelta(legs[i]))
				front = append(front, acc)
			} else if len(front) > 0 {
				break // interruption inside the run
			}
		}
		// back-run partial sums: the contiguous run of legs that disposed the asset in the
		// opposite direction, nearest the victim first (ascending), same gap/break rule.
		var back []*big.Int
		acc = new(big.Int)
		for i := vi + 1; i < len(legs); i++ {
			if legs[i].isToken0To1 != buySide && assetDelta(legs[i]).Sign() < 0 {
				acc = new(big.Int).Sub(acc, assetDelta(legs[i])) // accumulate the positive magnitude
				back = append(back, acc)
			} else if len(back) > 0 {
				break // interruption inside the run
			}
		}
		// Both runs are ascending, so a within-tolerance pair (if any) is found by a
		// two-pointer sweep: advancing the smaller side can only discard a value that is
		// already too small to match anything remaining on the other side.
		i, j := 0, 0
		for i < len(front) && j < len(back) {
			if withinTol(front[i], back[j]) {
				return true
			}
			if front[i].Cmp(back[j]) < 0 {
				i++
			} else {
				j++
			}
		}
	}
	return false
}

// withinTol reports whether non-negative x and y are within sandwichTolPercent of each
// other, measured against the larger of the two. Two zero values do not match (no volume).
func withinTol(x, y *big.Int) bool {
	max := x
	if y.Cmp(max) > 0 {
		max = y
	}
	if max.Sign() == 0 {
		return false
	}
	diff := new(big.Int).Abs(new(big.Int).Sub(x, y))
	// diff/max <= sandwichTolPercent/100  <=>  diff*100 <= sandwichTolPercent*max
	return new(big.Int).Mul(diff, big.NewInt(100)).Cmp(new(big.Int).Mul(big.NewInt(sandwichTolPercent), max)) <= 0
}

// hasDirectionalPattern reports whether swap directions contain a Buy-Buy-Sell (T,T,F) or
// Sell-Sell-Buy (F,F,T) subsequence. A single left-to-right pass suffices: track how many
// buys and sells have been seen so far; a sell seen after at least two buys completes
// Buy-Buy-Sell, and a buy seen after at least two sells completes Sell-Sell-Buy.
func hasDirectionalPattern(directions []bool) bool {
	buys, sells := 0, 0
	for _, isBuy := range directions {
		if isBuy {
			if sells >= 2 {
				return true // Sell-Sell-Buy
			}
			buys++
		} else {
			if buys >= 2 {
				return true // Buy-Buy-Sell
			}
			sells++
		}
	}
	return false
}
