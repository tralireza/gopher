package gopher

import "log"

// 1991 Find the Middle Index in Array
func findMiddleIndex(nums []int) int {
	rSum := 0
	for _, n := range nums {
		rSum += n
	}

	lSum := 0
	for i, n := range nums {
		rSum -= n
		if rSum == lSum {
			return i
		}
		lSum += n
	}
	return -1
}

// 2134m Minimum Swaps to Group All 1's Together II
func minSwaps(nums []int) int {
	ones := 0
	for _, n := range nums {
		ones += n
	}

	circular := make([]int, len(nums)*2)
	copy(circular, nums)
	copy(circular[len(nums):], nums)

	log.Print(nums, " -> ", circular)

	// Prefix Sum for zeros
	pSum := make([]int, 2*len(nums)+1)
	for i := range circular {
		pSum[i+1] = pSum[i]
		if circular[i] == 0 {
			pSum[i+1]++
		}
	}

	ops := len(nums) - ones
	for r := ones - 1; r < len(circular); r++ {
		ops = min(pSum[r+1]-pSum[r-ones+1], ops)
	}
	return ops
}

// 2574 Left and Right Sum Difference
func leftRightDifference(nums []int) []int {
	lSum, rSum := 0, 0
	for _, n := range nums {
		rSum += n
	}

	R := []int{}
	for _, n := range nums {
		rSum -= n
		r := rSum - lSum
		if r < 0 {
			r *= -1
		}
		R = append(R, r)
		lSum += n
	}
	return R
}
