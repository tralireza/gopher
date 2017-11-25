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
