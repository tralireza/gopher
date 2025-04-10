// Array
package gopher

import (
	"log"
	"slices"
)

// 485 Max Consecutive Ones
func findMaxConsecutiveOnes(nums []int) int {
	tCur, tMax := 0, 0
	for _, n := range nums {
		if tCur+n > tCur {
			tCur++
		} else {
			tCur = 0
		}

		tMax = max(tCur, tMax)
	}

	return tMax
}

// 1437 Check If All 1's Are at Least Length K Places Away
func kLengthApart(nums []int, k int) bool {
	dist := k
	for _, n := range nums {
		switch n {
		case 1:
			if dist < k {
				return false
			}
			dist = 0
		case 0:
			dist++
		}
	}

	return true
}

// 1752 Check If Array Is Sorted and Rotated
func check(nums []int) bool {
	inversions := 0
	if nums[0] < nums[len(nums)-1] {
		inversions++
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			inversions++
		}
	}
	return inversions <= 1
}

// 2033m Minimum Operations to Make a Uni-Value Grid
func minOperations_UniValue(grid [][]int, x int) int {
	nums := make([]int, 0, len(grid)*len(grid[0]))
	for r := range grid {
		nums = append(nums, grid[r]...)
	}

	slices.Sort(nums)
	median := nums[len(nums)/2]

	log.Print("-> ", nums, median)

	Abs := func(v int) int {
		if v >= 0 {
			return v
		}
		return -v
	}

	ops, r := 0, median%x
	for _, n := range nums {
		if n%x != r {
			return -1
		}

		ops += Abs(n-median) / x
	}

	return ops
}

// 2780m Minimum Index of a Valid Split
func minimumIndex(nums []int) int {
	F := map[int]int{}
	for _, n := range nums {
		F[n]++
	}

	dominant, frq := 0, 0
	for n, f := range F {
		if f > frq {
			dominant, frq = n, f
		}
	}

	log.Print("-> ", dominant, frq)

	f := 0
	for i, n := range nums {
		if n == dominant {
			f++
		}

		if f*2 > (i+1) && (frq-f)*2 > len(nums)-1-i {
			return i
		}
	}

	return -1
}

// 3169m Count Days Without Meetings
func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(x, y []int) int {
		if x[0] == y[0] {
			return x[1] - y[1]
		}
		return x[0] - y[0]
	})

	log.Print("-> ", meetings)

	t := 0

	lDay := 0
	for _, meeting := range meetings {
		start, finish := meeting[0], meeting[1]
		if start > lDay {
			t += start - lDay - 1
		}

		lDay = max(lDay, finish)
	}
	t += days - lDay

	return t
}

// 3394m Check if Grid can be Cut into Sections
func checkValidCuts(n int, rectangles [][]int) bool {
	Check := func(offset int) bool {
		slices.SortFunc(rectangles, func(x, y []int) int { return x[offset] - y[offset] })

		gaps, end := 0, rectangles[0][offset+2]
		for _, rectangle := range rectangles[1:] {
			if end <= rectangle[offset] {
				gaps++
			}
			end = max(rectangle[offset+2], end)
		}

		return gaps >= 2
	}
	return Check(0) || Check(1)
}
