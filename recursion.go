package gopher

import "slices"

// 40m Combination Sum II
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	R := [][]int{}

	var v []int
	var Search func(start, cSum int)
	Search = func(start, cSum int) {
		if cSum == target {
			R = append(R, append([]int{}, v...))
			return
		}

		// Prune
		if cSum > target {
			return
		}
		if start == len(candidates) {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] { // Group & Prune
				continue
			}
			v = append(v, candidates[i])
			Search(i+1, cSum+candidates[i])
			v = v[:len(v)-1]
		}
	}

	Search(0, 0)

	return R
}

// 650m 2 Keys Keyboard
func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	ops := n

	var CopyPaste func(l, lp, curOps int)
	CopyPaste = func(l, lp, curOps int) {
		if l >= n {
			if l == n {
				ops = min(ops, curOps)
			}
			return
		}

		CopyPaste(l+lp, lp, curOps+1) // just Paste
		CopyPaste(l+l, l, curOps+2)   // Copy & Paste
	}

	CopyPaste(1, 1, 1)

	return ops
}
