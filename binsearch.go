package gopher

import (
	"log"
	"math"
	"slices"
)

// 3224m Minimum Array Changes to Make Difference Equal
func minChanges(nums []int, k int) int {
	M := map[int]int{}
	Diffs := make([]int, 0, len(nums)/2)

	l, r := 0, len(nums)-1
	for l < r {
		A, a := nums[l], nums[r]
		if a > A {
			A, a = a, A
		}

		M[A-a]++

		// maximum difference of "pair" elements that can be fixed by one operation
		// ... by setting either: a to 0 or A to k
		Diffs = append(Diffs, max(A, k-a))

		l++
		r--
	}

	log.Print("Difference Frequency -> ", M)

	slices.Sort(Diffs)
	log.Print("(One Operation) Maximum Difference -> ", Diffs)

	minOps := math.MaxInt
	for x, f := range M {
		l, r := 0, len(Diffs)-1
		for l < r {
			m := l + (r-l)>>1
			if Diffs[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		minOps = min(minOps, len(nums)/2-f+l)
	}
	return minOps
}
