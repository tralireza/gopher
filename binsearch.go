package gopher

import (
	"log"
	"math"
	"slices"
)

// 274m H-Index
func hIndex(citations []int) int {
	slices.Sort(citations)

	Check := func(m int) int {
		x := 0
		for i := 0; i < len(citations); i++ {
			if citations[i] >= m {
				x++
			}
		}
		return x
	}

	l, r := 0, len(citations)
	var h int
	for l <= r {
		m := l + (r-l)>>1

		v := Check(m)

		log.Print(l, m, r, " :: ", v)

		if v >= m {
			l = m + 1
			h = m
		} else {
			r = m - 1
		}
	}
	return h
}

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
