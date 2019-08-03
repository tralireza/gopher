package gopher

import (
	"log"
	"slices"
	"testing"
)

// 485 Max Consecutive Ones
func Test485(t *testing.T) {
	log.Print("3 ?= ", findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	log.Print("2 ?= ", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

// 1437 Check If All 1's Are at Least Length K Places Away
func Test1437(t *testing.T) {
	log.Print("true ?= ", kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	log.Print("false ?= ", kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

// 1752 Check If Array Is Sorted and Rotated
func Test1752(t *testing.T) {
	log.Print("true ?= ", check([]int{3, 4, 5, 1, 2}))
	log.Print("false ?= ", check([]int{2, 1, 3, 4}))
	log.Print("true ?= ", check([]int{1, 2, 3}))
}

func Test2033(t *testing.T) {
	TwoPointers := func(grid [][]int, x int) int {
		nums := make([]int, 0, len(grid)*len(grid[0]))
		for r := range grid {
			for c := range grid[r] {
				if grid[r][c]%x != grid[0][0]%x {
					return -1
				}
				nums = append(nums, grid[r][c])
			}
		}

		slices.Sort(nums)
		log.Print("-> ", nums)

		ops := 0

		prefix, suffix := 0, len(nums)-1
		for prefix < suffix {
			if prefix < len(nums)-1-suffix {
				ops += (nums[prefix+1] - nums[prefix]) / x * (prefix + 1)
				prefix++
			} else {
				ops += (nums[suffix] - nums[suffix-1]) / x * (len(nums) - suffix)
				suffix--
			}
		}

		return ops
	}

	for _, c := range []struct {
		rst  int
		grid [][]int
		x    int
	}{
		{4, [][]int{{2, 4}, {6, 8}}, 2},
		{5, [][]int{{1, 5}, {2, 3}}, 1},
		{-1, [][]int{{1, 2}, {3, 4}}, 2},
	} {
		rst, grid, x := c.rst, c.grid, c.x
		for _, f := range []func([][]int, int) int{minOperations_UniValue, TwoPointers} {
			if rst != f(grid, x) {
				t.FailNow()
			}
		}
		log.Printf(":: %d <- %v / %d", rst, grid, x)
	}
}

func Test2780(t *testing.T) {
	// 1 <= N_i <= 10^9

	BoyerMoore := func(nums []int) int {
		majority := nums[0]

		count := 0
		for _, n := range nums {
			if n == majority {
				count++
			} else {
				count--
			}

			if count == 0 {
				majority, count = n, 1
			}
		}

		return majority
	}

	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{2, []int{1, 2, 2, 2}},
		{4, []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}},
		{-1, []int{3, 3, 3, 3, 7, 2, 2}},
	} {
		if c.rst != minimumIndex(c.nums) {
			t.FailNow()
		}
		log.Printf(":: %d <- %v | Boyer-Moore: %d", c.rst, c.nums, BoyerMoore(c.nums))
	}
}

func Test2873(t *testing.T) {
	Optimized := func(nums []int) int64 {
		leftMax := nums[0]

		rightMaxs := make([]int, len(nums))
		rightMaxs[len(nums)-1] = nums[len(nums)-1]
		for r := len(nums) - 2; r >= 2; r-- {
			rightMaxs[r] = max(nums[r], rightMaxs[r+1])
		}

		xVal := int64(0)
		for i, n := range nums[1 : len(nums)-2] {
			xVal = max(xVal, int64(leftMax-n)*int64(rightMaxs[i+1]))
			leftMax = max(n, leftMax)
		}

		return xVal
	}

	for _, c := range []struct {
		rst  int64
		nums []int
	}{
		{77, []int{12, 6, 1, 2, 7}},
		{133, []int{1, 10, 3, 4, 19}},
		{0, []int{1, 2, 3}},
	} {
		for _, f := range []func([]int) int64{maximumTripletValue, Optimized} {
			if c.rst != f(c.nums) {
				t.FailNow()
			}
		}
		log.Printf(":: %d   <- %v", c.rst, c.nums)
	}
}

func Test3169(t *testing.T) {
	for _, c := range []struct {
		rst, days int
		meetings  [][]int
	}{
		{2, 10, [][]int{{5, 7}, {1, 3}, {9, 10}}},
		{1, 5, [][]int{{2, 4}, {1, 3}}},
		{0, 6, [][]int{{1, 6}}},

		{1, 8, [][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}},
	} {
		rst, days, meetings := c.rst, c.days, c.meetings
		if rst != countDays(days, meetings) {
			t.FailNow()
		}
		log.Printf(":: %d <- %d / %v", rst, days, meetings)
	}
}

func Test3394(t *testing.T) {
	for _, c := range []struct {
		rst        bool
		n          int
		rectangles [][]int
	}{
		{true, 5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}},
		{true, 4, [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}},
		{false, 4, [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}},

		{false, 3, [][]int{{0, 0, 1, 3}, {1, 0, 2, 2}, {2, 0, 3, 2}, {1, 2, 3, 3}}},
	} {
		rst, n, rectangles := c.rst, c.n, c.rectangles
		if rst != checkValidCuts(n, rectangles) {
			t.FailNow()
		}
		log.Printf(":: %t <- %v", rst, rectangles)
	}
}
