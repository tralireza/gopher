package gopher

import (
	"log"
	"slices"
	"testing"
)

// 239h Sliding Window Maximum
func Test239(t *testing.T) {
	log.Print("[3 3 5 5 6 7] ?= ", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	log.Print("[1 -1] ?= ", maxSlidingWindow([]int{1, -1}, 1))
}

// 1508m Range Sum of Sorted Subarray Sums
func Test1508(t *testing.T) {
	// 1 <= left, right <= n*(n+1)/2
	WithBinSearch := func(nums []int, n, left, right int) int {
		minSum := slices.Min(nums)
		maxSum := nums[0]
		for _, n := range nums[1:] {
			maxSum += n
		}

		log.Printf("BinSearch (space, sub-array sums): [ %d ... %d ]", minSum, maxSum)

		SumOfK := func(k int) int {
			Count := func(targetSum int) (int, int) {
				totalSum, count := 0, 0
				for l := 0; l < len(nums); l++ {
					curSum := 0
					for r := l; r < len(nums); r++ {
						curSum += nums[r]
						if curSum <= targetSum {
							count++
							totalSum += curSum
						}
					}
				}
				return count, totalSum
			}

			l, r := minSum, maxSum
			for l <= r {
				m := l + (r-l)>>1

				i, t := Count(m)

				if i == k {
					return t
				} else if i > k {
					r = m - 1
				} else {
					l = m + 1
				}
			}

			return 0
		}

		return SumOfK(right) - SumOfK(left-1)
	}

	for _, f := range []func([]int, int, int, int) int{rangeSum, WithBinSearch} {
		log.Print("13 ?= ", f([]int{1, 2, 3, 4}, 4, 1, 5))
		log.Print("6 ?= ", f([]int{1, 2, 3, 4}, 4, 3, 4))
		log.Print("50 ?= ", f([]int{1, 2, 3, 4}, 4, 1, 10))
		log.Print("---")
	}
}

// 3256h Maximum Value Sum by Placing Three Rooks I
func Test3256(t *testing.T) {
	// 3 <= Rows, Cols <= 100

	log.Print("4 ?= ", maximumValueSum([][]int{{-3, 1, 1, 1}, {-3, 1, -3, 1}, {-3, 2, 1, 1}}))
	log.Print("15 ?= ", maximumValueSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	log.Print("3 ?= ", maximumValueSum([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

// 3257h Maximum Value Sum by Placing Three Rooks II
func Test3257(t *testing.T) {
	// 3 <= Rows, Cols <= 500

	log.Print("4 ?= ", maximumValueSum2([][]int{{-3, 1, 1, 1}, {-3, 1, -3, 1}, {-3, 2, 1, 1}}))
	log.Print("15 ?= ", maximumValueSum2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	log.Print("3 ?= ", maximumValueSum2([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}
