package gopher

import (
	"log"
	"reflect"
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

// 407h Trapping Rain Water II
func Test407(t *testing.T) {
	log.Print("4 ?= ", trapRainWater([][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}))
	log.Print("10 ?= ", trapRainWater([][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}))
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

// 1792m Maximum Average Pass Ratio
func Test1792(t *testing.T) {
	for _, c := range []struct {
		r             float64
		classes       [][]int
		extraStudents int
	}{
		{0.78333, [][]int{{1, 2}, {3, 5}, {2, 2}}, 2},
		{0.53484, [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4},
	} {
		t.Run("", func(t *testing.T) {
			if maxAverageRatio(c.classes, c.extraStudents)-c.r >= 0.00001 {
				t.Fail()
			}
		})
	}
}

// 1942m The Number of the Smallest Unoccupied Chair
func Test1942(t *testing.T) {
	log.Print("1 ?= ", smallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
	log.Print("2 ?= ", smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
	log.Print("2 ?= ", smallestChair([][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}, 6))
}

// 2182m Construct String With Repeat Limit
func Test2182(t *testing.T) {
	log.Print("zzcccac ?= ", repeatLimitedString("cczazcc", 3))
	log.Print("bbabaa ?= ", repeatLimitedString("ababab", 2))
}

// 2406m Divide Intervals Into Minimum Number of Groups
func Test2406(t *testing.T) {
	log.Print("3 ?= ", minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}))
	log.Print("1 ?= ", minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}))
}

// 2530m Maximal Score After Applying K Operations
func Test2530(t *testing.T) {
	log.Print("50 ?= ", maxKelements([]int{10, 10, 10, 10, 10}, 5))
	log.Print("17 ?= ", maxKelements([]int{1, 10, 3, 3, 3}, 3))
}

// 2558 Take Gifts From the Richest Pile
func Test2558(t *testing.T) {
	log.Print("29 ?= ", pickGifts([]int{25, 64, 9, 4, 100}, 4))
	log.Print("4 ?= ", pickGifts([]int{1, 1, 1, 1}, 4))
}

func Test2818(t *testing.T) {
	// 1 <= N <= 10^5

	for _, c := range []struct {
		rst, k int
		nums   []int
	}{
		{81, 2, []int{8, 3, 9, 3, 8}},
		{4788, 3, []int{19, 12, 14, 6, 10, 18}},

		{256720975, 6, []int{3289, 2832, 14858, 22011}},
		{630596200, 27, []int{6, 1, 13, 10, 1, 17, 6}},
	} {
		log.Print("** ", c.nums, c.k)
		if c.rst != maximumScore(c.nums, c.k) {
			t.FailNow()
		}
		log.Printf(":: %d", c.rst)
	}
}

// 2940m Find Building Where Alice and Bob Can Meet
func Test2940(t *testing.T) {
	log.Print("[2 5 -1 5 2] ?= ", leftmostBuildingQueries([]int{6, 4, 8, 5, 2, 7}, [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}))
	log.Print("[7 6 -1 4 6] ?= ", leftmostBuildingQueries([]int{5, 3, 8, 2, 6, 1, 4, 6}, [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}))
}

// 3066m Minimum Operations to Exceed Threshold Value II
func Test3066(t *testing.T) {
	log.Print("2 ?= ", minOperationsII([]int{2, 11, 10, 1, 3}, 10))
	log.Print("4 ?= ", minOperationsII([]int{1, 1, 2, 4, 9}, 20))
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

// 3264 Final Array State After K Multiplication Operations I
func Test3264(t *testing.T) {
	for _, c := range []struct {
		r          []int
		nums       []int
		k          int
		multiplier int
	}{
		{[]int{8, 4, 6, 5, 6}, []int{2, 1, 3, 5, 6}, 5, 2},
		{[]int{16, 8}, []int{1, 2}, 3, 4},
	} {
		t.Run("", func(t *testing.T) {
			if !reflect.DeepEqual(c.r, getFinalState(c.nums, c.k, c.multiplier)) {
				t.Fail()
			}
		})
	}
}
