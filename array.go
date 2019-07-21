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
