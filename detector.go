// Package bscexorcist provides sandwich attack detection for BSC transaction bundles.
package bscexorcist

import (
	"fmt"

	"github.com/48Club/bscexorcist/protocols"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DetectSandwichForBundle analyzes a bundle of transaction logs to identify potential sandwich attacks.
// Returns an error if a sandwich pattern is detected in any pool within the bundle.
func DetectSandwichForBundle(bundleLogs [][]*types.Log) error {
	if len(bundleLogs) < 3 {
		return nil
	}

	poolDirections := make(map[common.Address][]bool)
	for _, txLogs := range bundleLogs {
		// Collapse all swaps a single transaction makes on the same pool into one
		// net-direction entry. A transaction that routes through the same pool
		// several times (an aggregator split or a multi-hop arbitrage) would
		// otherwise contribute multiple same-direction legs, fabricating a
		// "Buy-Buy" / "Sell-Sell" prefix that misclassifies an honest backrun as a
		// sandwich. A genuine sandwich keeps its front-run and back-run in separate
		// transactions (the victim's transaction sits between them), so collapsing
		// per transaction never hides a real attack.
		for poolID, isToken0To1 := range netDirectionsPerPool(txLogs) {
			poolDirections[poolID] = append(poolDirections[poolID], isToken0To1)
		}
	}

	for pool, directions := range poolDirections {
		if hasSandwichPattern(directions) {
			return fmt.Errorf("sandwich attack detected on pool: %s", pool.Hex())
		}
	}

	return nil
}

// netDirectionsPerPool reduces all swaps in a single transaction to one net
// direction per pool. For each pool it tallies how many swaps go token0->token1
// versus the reverse and reports the majority direction; pools with an equal split
// carry no net directional signal and are omitted.
//
// Only IsToken0To1 is consulted. AmountIn/AmountOut are deliberately avoided: their
// meaning is not uniform across adapters (e.g. DODO indexes them by token0/token1
// rather than input/output, and FourMeme returns zero for both), so any amount-based
// netting would silently misread or drop those protocols.
func netDirectionsPerPool(txLogs []*types.Log) map[common.Address]bool {
	type tally struct{ token0To1, token1To0 int }
	counts := make(map[common.Address]*tally)
	for _, swap := range protocols.ParseSwapEvents(txLogs) {
		poolID := swap.PairID()
		if counts[poolID] == nil {
			counts[poolID] = &tally{}
		}
		if swap.IsToken0To1() {
			counts[poolID].token0To1++
		} else {
			counts[poolID].token1To0++
		}
	}

	directions := make(map[common.Address]bool, len(counts))
	for poolID, t := range counts {
		switch {
		case t.token0To1 > t.token1To0:
			directions[poolID] = true
		case t.token1To0 > t.token0To1:
			directions[poolID] = false
		}
	}
	return directions
}

// hasSandwichPattern checks if swap directions form a sandwich attack pattern.
func hasSandwichPattern(directions []bool) bool {
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
