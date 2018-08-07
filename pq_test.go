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

// 373m Find K Pairs with Smallest Sums
func Test373(t *testing.T) {
	// 1 <= length N1, N2 <= 10^5
	// 1 <= k <= 10^4

	log.Print("[[1 2] [1 4] [1 6]] ?= ", kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	log.Print("[[1 1] [1 1]] ?= ", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	log.Print("--")
	log.Print(" ?= ", kSmallestPairs([]int{1, 2, 4, 5, 6}, []int{3, 5, 7, 9}, 20))
}

// 632h Smallest Range Covering Elements from K Lists
func Test632(t *testing.T) {
	// -10^5 <= N_ij <= 10^5, 1 <= N.length <= 3500, 1 <=N_i.length <= 50

	BruteForce := func(nums [][]int) []int {
		Idx := make([]int, len(nums))
		start, end := -100_000, 100_000

		for {
			curMin, curMax := 100_000, -100_000
			var curI int // minimum value index

			for i := range len(nums) {
				v := nums[i][Idx[i]]
				if v < curMin {
					curMin = v
					curI = i
				}
				if v > curMax {
					curMax = v
				}
			}

			if curMax-curMin < end-start {
				start, end = curMin, curMax
			}

			Idx[curI]++
			if Idx[curI] == len(nums[curI]) { // one list (interval) is complete
				return []int{start, end}
			}
		}

		return []int{}
	}

	for _, fn := range []func([][]int) []int{BruteForce, smallestRange} {
		log.Print("[20 24] ?= ", fn([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}))
		log.Print("[1 1] ?= ", fn([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}))
		log.Print("--")
	}
}

// 1405m Longest Happy String
func Test1405(t *testing.T) {
	log.Printf(`"ccaccbcc" ?= %q`, longestDiverseString(1, 1, 7))
	log.Printf(`"aabaa" ?= %q`, longestDiverseString(7, 1, 0))
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

// 1942m The Number of the Smallest Unoccupied Chair
func Test1942(t *testing.T) {
	log.Print("1 ?= ", smallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
	log.Print("2 ?= ", smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
	log.Print("2 ?= ", smallestChair([][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}, 6))
}

// 2406m Divide Intervals Into Minimum Number of Groups
func Test2406(t *testing.T) {
	log.Print("3 ?= ", minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}))
	log.Print("1 ?= ", minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}))

// 2530m Maximal Score After Applying K Operations
func Test2530(t *testing.T) {
	log.Print("50 ?= ", maxKelements([]int{10, 10, 10, 10, 10}, 5))
	log.Print("17 ?= ", maxKelements([]int{1, 10, 3, 3, 3}, 3))
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
