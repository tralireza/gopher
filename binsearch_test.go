package gopher

import (
	"log"
	"testing"
)

// 153m Find Minimum in Rotated Sorted Array
func Test153(t *testing.T) {
	log.Print("1 ?= ", findMin([]int{3, 4, 5, 1, 2}))
	log.Print("0 ?= ", findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	log.Print("11 ?= ", findMin([]int{11, 13, 15, 17}))
}

// 154h Find Minimum in Rotated Sorted Array II
func Test154(t *testing.T) {
	log.Print("1 ?= ", findMinII([]int{1, 3, 5}))
	log.Print("0 ?= ", findMinII([]int{2, 2, 2, 0, 1}))
}

// 274m H-Index
func Test274(t *testing.T) {
	log.Print("3 ?= ", hIndex([]int{3, 0, 6, 1, 5}))
	log.Print("1 ?= ", hIndex([]int{1, 3, 1}))
	log.Print("1 ?= ", hIndex([]int{1}))
	log.Print("0 ?= ", hIndex([]int{0}))
}

// 492 Construct the Rectangle
func Test492(t *testing.T) {
	log.Print("[2 2] ?= ", constructRectangle(4))
	log.Print("[37 1] ?= ", constructRectangle(37))
	log.Print("[427 286] ?= ", constructRectangle(122122))
}

// 564h Find the Closest Palindrome
func Test564(t *testing.T) {
	// 1 <= N <= 10^18-1

	log.Print("121 ?= ", nearestPalindromic("123"))
	log.Print("0 ?= ", nearestPalindromic("1"))
	log.Print("99799 ?= ", nearestPalindromic("99800"))
}

// 704 Binary Search
func Test704(t *testing.T) {
	for _, c := range []struct {
		rst, target int
		nums        []int
	}{
		{4, 9, []int{-1, 0, 3, 5, 9, 12}},
		{-1, 2, []int{-1, 0, 3, 5, 9, 12}},
		{0, -1, []int{-1, 0, 3, 5, 9, 12}},
	} {
		rst, target, nums := c.rst, c.target, c.nums
		log.Printf("%d ?= %d", rst, search(nums, target))
		if rst != search(nums, target) {
			t.FailNow()
		}
	}
}

// 1351 Count Negative Numbers in a Sorted Matrix
func Test1351(t *testing.T) {
	for _, c := range []struct {
		rst  int
		grid [][]int
	}{
		{8, [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}},
		{0, [][]int{{3, 2}, {1, 0}}},
	} {
		rst, grid := c.rst, c.grid
		log.Printf("%d ?= %d", rst, countNegatives(grid))
		if rst != countNegatives(grid) {
			t.FailNow()
		}
	}
}

// 1760m Minimum Limit of Balls in a Bag
func Test1760(t *testing.T) {
	log.Print("3 ?= ", minimumSize([]int{9}, 2))
	log.Print("2 ?= ", minimumSize([]int{2, 4, 8, 2}, 4))
	log.Print("7 ?= ", minimumSize([]int{7, 17}, 2))
}

// 1894m Find the Student that Will Replace the Chalk
func Test1894(t *testing.T) {
	// left-most BinSearch
	lBS := func(A []int, k int) int {
		l, r := 0, len(A)
		for l < r {
			m := l + (r-l)>>1
			if A[m] < k { // l <= m < r
				l = m + 1 // Keep: A[l-1] < k
			} else {
				r = m // Keep: A[r] >= k
			}
		}
		return l
	}

	// right-most BinSearch
	rBS := func(A []int, k int) int {
		l, r := 0, len(A)
		for l < r {
			m := l + (r-l)>>1 // l <= m < r
			if A[m] > k {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	A := []int{2, 3, 3, 3, 4, 5, 7, 7, 8}
	log.Print("      0 1 2 3 4 5 6 7 8")
	log.Print("A :: ", A)
	for _, k := range []int{1, 2, 3, 6, 7, 8, 9} {
		log.Print(k, "?   ==L=> ", lBS(A, k), lBS(A, k+1), "   ==R=> ", rBS(A, k))
	}

	log.Print("0 ?= ", chalkReplacer([]int{5, 1, 5}, 22))
	log.Print("1 ?= ", chalkReplacer([]int{3, 4, 1, 2}, 25))
}

// 2529 Maximum Count of Positive Integer and Negative Integer
func Test2529(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{3, []int{-2, -1, -1, 1, 2, 3}},
		{3, []int{-3, -2, -1, 0, 0, 1, 2}},
		{4, []int{5, 20, 66, 1314}},

		{3, []int{-3, -2, -1}},
		{3, []int{-3, -2, -1, 0}},
	} {
		rst, nums := c.rst, c.nums
		log.Printf("%d ?= %d", rst, maximumCount(nums))
		if rst != maximumCount(nums) {
			t.FailNow()
		}
	}
}

// 3224m Minimum Array Changes to Make Difference Equal
func Test3224(t *testing.T) {
	// 0 <= Array[i] <= k <= 10^5

	log.Print("2 ?= ", minChanges([]int{1, 0, 1, 2, 4, 3}, 4))
	log.Print("2 ?= ", minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6))
}

// 3296m Minimum Number of Seconds to Make Mountain Height Zero
func Test3296(t *testing.T) {
	log.Print("3 ?= ", minNumberOfSeconds(4, []int{2, 1, 1}))
	log.Print("12 ?= ", minNumberOfSeconds(10, []int{3, 2, 2, 4}))
	log.Print("15 ?= ", minNumberOfSeconds(5, []int{1}))
}

// 3356m Zero Array Transformation II
func Test3356(t *testing.T) {
	LineSweep := func(nums []int, queries [][]int) int {
		Diffs := make([]int, len(nums)+1)

		k, tSum := 0, 0
		for i := range nums {
			for tSum+Diffs[i] < nums[i] {
				k++
				if k > len(queries) {
					return -1
				}

				qry := queries[k-1]
				if i <= qry[1] {
					Diffs[max(i, qry[0])] += qry[2]
					Diffs[qry[1]+1] -= qry[2]
				}
			}

			tSum += Diffs[i]
		}

		return k
	}

	for _, c := range []struct {
		rst     int
		nums    []int
		queries [][]int
	}{
		{2, []int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}},
		{-1, []int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}},
	} {
		rst, nums, queries := c.rst, c.nums, c.queries
		for _, f := range []func([]int, [][]int) int{minZeroArray, LineSweep} {
			if rst != f(nums, queries) {
				t.FailNow()
			}
			log.Printf(":: %v <- %d", rst, nums)
		}
	}
}
