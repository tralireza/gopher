package gopher

import (
	"log"
	"slices"
)

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

	rCalls, Mem := 0, map[[2]int]int{}
	ops := n

	var CopyPaste func(l, lp int) int
	CopyPaste = func(l, lp int) int {
		if l >= n {
			if l == n {
				return 0
			}
			return n
		}

		rCalls++
		if v, ok := Mem[[2]int{l, lp}]; ok {
			return v
		}

		p := CopyPaste(l+lp, lp) // just Paste
		cp := CopyPaste(l+l, l)  // Copy & Paste

		Mem[[2]int{l, lp}] = min(1+p, 2+cp)
		return Mem[[2]int{l, lp}]
	}

	ops = 1 + CopyPaste(1, 1)
	log.Print(" -> ", rCalls, " # rCalls")

	return ops
}
