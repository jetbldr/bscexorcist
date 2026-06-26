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
// the asset bought in the front-run and the asset sold back in the back-run. A genuine
// attacker unwinds almost exactly what it acquired (residual ~0; real captures sit
// within a few percent), so a small tolerance separates true sandwiches from unrelated
// trades that merely share a pool and direction.
const sandwichTolPercent = 15

// maxLegsForConservation is the leg count up to which the exact two-sided subset search
// (conserves) runs. Front and back are disjoint subsets, so it is O(n*2^n) overall (~1M
// big.Int ops at n=16). Above it, conservesByTotals takes over (scalable, slightly weaker);
// above maxLegsHardCap the result is left indeterminate (not flagged) as a work bound.
const maxLegsForConservation = 16

// maxLegsHardCap bounds total work. The scalable conservesByTotals path is O(n^2) and
// handles pools up to this many legs; this sits around the most swap transactions that can
// fit in one (half-gaslimit) block, so an economically viable bundle stays under it and is
// always conservation-checked. Above it the result is left indeterminate — and indeterminate
// means NOT a sandwich, never a forced flag, so a large bundle cannot trip a false positive
// merely by its leg count. The residual (a conserving sandwich padded past the cap is missed)
// requires hundreds of same-pool transactions, whose gas and self-inflicted slippage dwarf
// any sandwich profit.
const maxLegsHardCap = 512

// overCapFrontSubsetCap bounds the front-run subset enumeration on the scalable path. A
// front-run split into more transactions than this on one pool is not realistic; beyond it
// the scalable path matches single front legs only.
const overCapFrontSubsetCap = 12

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
// sandwich. It requires both:
//
//  1. the directional shape Buy-Buy-Sell or Sell-Sell-Buy (a necessary precondition), and
//  2. asset conservation around a bracketed victim — a front-run whose acquired asset is
//     sold back by a later back-run, with the victim trading the same direction in between.
//
// Condition 2 is what distinguishes a real sandwich from unrelated trades that merely
// share a pool and direction (e.g. one searcher's incidental swap plus another searcher's
// multi-leg arbitrage). It is sender-agnostic: the bought amount equals the sold-back
// amount regardless of which addresses sign the front-run and back-run, so it still
// catches sandwiches that spread their legs across multiple addresses.
func isSandwich(legs []swapLeg) bool {
	n := len(legs)
	if n < 3 {
		return false
	}
	dirs := make([]bool, n)
	for i, l := range legs {
		dirs[i] = l.isToken0To1
	}
	if !hasDirectionalPattern(dirs) {
		return false
	}

	// A pool is single-protocol; if its amounts are not reliable for value reconciliation
	// (DODO, FourMeme) fall back to direction-only flagging, preserving prior behavior.
	for _, l := range legs {
		if !l.reliable {
			return true
		}
	}
	if n > maxLegsHardCap {
		// Too many legs to analyze affordably. Leave it indeterminate rather than flag:
		// forcing a sandwich verdict here would be a direction-only false positive — exactly
		// what this gate removes — and would let any complex same-pool bundle trip it.
		return false
	}
	if n > maxLegsForConservation {
		// Too many legs for the exact subset search. Use the scalable conservation
		// approximation rather than reverting to direction-only flagging, which would
		// refire the false positive this gate exists to prevent.
		return conservesByTotals(legs, true) || conservesByTotals(legs, false)
	}

	// A buy side of token0->token1 makes token1 the asset (Buy-Buy-Sell); the reverse
	// makes token0 the asset (Sell-Sell-Buy). Try both orientations.
	return conserves(legs, true) || conserves(legs, false)
}

// conserves checks the asset-conservation condition for one orientation. buySide selects
// the front/victim direction; the asset token is token1 when buySide is true, else token0.
//
// The signed net asset change (dT1 or dT0), not its absolute value, decides whether a leg
// can participate: a front-run must have actually acquired the asset (delta > 0) and a
// back-run must have actually disposed of it (delta < 0). A transaction whose majority
// swap direction disagrees with its net asset flow — e.g. two tiny buys outvoting one
// large sell so the leg folds to "buy" while it net-sold — would otherwise have its
// magnitude matched against a real sell and fabricate a sandwich.
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
		// front-run candidates execute before the victim and must have acquired the asset;
		// back-run candidates execute after it and must have disposed of the asset.
		var front, back []*big.Int
		for i := 0; i < vi; i++ {
			if legs[i].isToken0To1 == buySide {
				if d := assetDelta(legs[i]); d.Sign() > 0 {
					front = append(front, d)
				}
			}
		}
		for i := vi + 1; i < len(legs); i++ {
			if legs[i].isToken0To1 != buySide {
				if d := assetDelta(legs[i]); d.Sign() < 0 {
					back = append(back, new(big.Int).Neg(d))
				}
			}
		}
		if len(front) == 0 || len(back) == 0 {
			continue
		}
		if subsetSumsMatch(front, back) {
			return true
		}
	}
	return false
}

// conservesByTotals is the scalable fallback used when a pool has more legs than the exact
// subset search can afford. For each candidate victim it matches the front-run against the
// total of the back-run after the victim. The back-run total absorbs a back-run split
// across transactions, and a bounded subset of the front legs absorbs a split front-run;
// only when there are more front legs than overCapFrontSubsetCap does it fall back to
// matching single front legs. Unlike direction-only flagging it never reports a bundle
// whose amounts do not conserve. Sign discipline matches conserves: front and victim must
// have net-acquired the asset, the back-run must have net-disposed of it.
//
// Residual (disproportionate to close cheaply, see maxLegsHardCap guard): on a pool with
// more than overCapFrontSubsetCap front legs before the victim, a split front-run can be
// missed; and padding sells after the victim inflates the back total past tolerance. Both
// require many same-pool legs in one bundle, which is costly for an attacker to construct.
func conservesByTotals(legs []swapLeg, buySide bool) bool {
	assetDelta := func(l swapLeg) *big.Int {
		if buySide {
			return l.dT1
		}
		return l.dT0
	}
	for vi := range legs {
		if legs[vi].isToken0To1 != buySide || assetDelta(legs[vi]).Sign() <= 0 {
			continue
		}
		backTotal := new(big.Int)
		for i := vi + 1; i < len(legs); i++ {
			if legs[i].isToken0To1 != buySide {
				if d := assetDelta(legs[i]); d.Sign() < 0 {
					backTotal.Sub(backTotal, d) // accumulate the positive magnitude
				}
			}
		}
		if backTotal.Sign() == 0 {
			continue
		}
		var front []*big.Int
		for i := 0; i < vi; i++ {
			if legs[i].isToken0To1 == buySide {
				if d := assetDelta(legs[i]); d.Sign() > 0 {
					front = append(front, d)
				}
			}
		}
		if len(front) == 0 {
			continue
		}
		if len(front) <= overCapFrontSubsetCap {
			for _, sf := range subsetSums(front) {
				if withinTol(sf, backTotal) {
					return true
				}
			}
		} else {
			for _, f := range front {
				if withinTol(f, backTotal) {
					return true
				}
			}
		}
	}
	return false
}

// subsetSumsMatch reports whether some nonempty subset of a and some nonempty subset of b
// have sums within tolerance of each other. This tolerates a front-run or back-run that is
// split across several transactions.
func subsetSumsMatch(a, b []*big.Int) bool {
	sa := subsetSums(a)
	sb := subsetSums(b)
	for _, x := range sa {
		for _, y := range sb {
			if withinTol(x, y) {
				return true
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

// subsetSums returns the sums of all 2^len(vals)-1 nonempty subsets of vals.
func subsetSums(vals []*big.Int) []*big.Int {
	sums := make([]*big.Int, 0, (1<<len(vals))-1)
	for mask := 1; mask < (1 << len(vals)); mask++ {
		s := new(big.Int)
		for i, v := range vals {
			if mask&(1<<i) != 0 {
				s.Add(s, v)
			}
		}
		sums = append(sums, s)
	}
	return sums
}

// hasDirectionalPattern checks if swap directions contain a sandwich-shaped subsequence.
func hasDirectionalPattern(directions []bool) bool {
	n := len(directions)
	if n < 3 {
		return false
	}

	// Look for Buy-Buy-Sell or Sell-Sell-Buy patterns
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				// Buy-Buy-Sell pattern
				if directions[i] && directions[j] && !directions[k] {
					return true
				}
				// Sell-Sell-Buy pattern
				if !directions[i] && !directions[j] && directions[k] {
					return true
				}
			}
		}
	}

	return false
}
