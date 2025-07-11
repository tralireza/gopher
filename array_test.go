package gopher

import (
	"log"
	"reflect"
	"slices"
	"testing"
)

func Test73(t *testing.T) {
	for _, c := range []struct {
		rst, matrix [][]int
	}{
		{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
		{[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}},
	} {
		setZeroes(c.matrix)
		if !reflect.DeepEqual(c.rst, c.matrix) {
			t.FailNow()
		}
	}
}

// 485 Max Consecutive Ones
func Test485(t *testing.T) {
	log.Print("3 ?= ", findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	log.Print("2 ?= ", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

func Test661(t *testing.T) {
	for _, c := range []struct {
		rst, img [][]int
	}{
		{
			[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		}, {
			[][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
			[][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
		},
	} {
		log.Print("* ", c.img)
		if !reflect.DeepEqual(c.rst, imageSmoother(c.img)) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test766(t *testing.T) {
	for _, c := range []struct {
		rst    bool
		matrix [][]int
	}{
		{true, [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}},
		{false, [][]int{{1, 2}, {2, 2}}},
	} {
		log.Print("* ", c.matrix)
		if c.rst != isToeplitzMatrix(c.matrix) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test798(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{3, []int{2, 3, 1, 4, 0}},
		{0, []int{1, 3, 0, 2, 4}},
	} {
		log.Print("* ", c.nums)
		if c.rst != bestRotation(c.nums) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test821(t *testing.T) {
	for _, c := range []struct {
		rst []int
		s   string
		chr byte
	}{
		{[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}, "loveleetcode", 'e'},
		{[]int{3, 2, 1, 0}, "aaab", 'b'},
	} {
		log.Printf("* %q %q", c.s, c.chr)
		if !reflect.DeepEqual(c.rst, shortestToChar(c.s, c.chr)) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test830(t *testing.T) {
	for _, c := range []struct {
		rst [][]int
		s   string
	}{
		{[][]int{{3, 6}}, "abbxxxxzzy"},
		{[][]int{}, "abc"},
		{[][]int{{3, 5}, {6, 9}, {12, 14}}, "abcdddeeeeaabbbcd"},
	} {
		log.Printf("* %q", c.s)
		if !reflect.DeepEqual(c.rst, largeGroupPositions(c.s)) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1394(t *testing.T) {
	for _, c := range []struct {
		rst int
		arr []int
	}{
		{2, []int{2, 2, 3, 4}},
		{3, []int{1, 2, 2, 3, 3, 3}},
		{-1, []int{2, 2, 2, 3, 3}},
	} {
		log.Print("* ", c.arr)
		if c.rst != findLucky(c.arr) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

// 1437 Check If All 1's Are at Least Length K Places Away
func Test1437(t *testing.T) {
	log.Print("true ?= ", kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	log.Print("false ?= ", kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

func Test1534(t *testing.T) {
	for _, c := range []struct {
		rst, a, b, c int
		arr          []int
	}{
		{4, 7, 2, 3, []int{3, 0, 1, 1, 9, 7}},
		{0, 0, 0, 1, []int{1, 1, 2, 2, 3}},
	} {
		if c.rst != countGoodTriplets(c.arr, c.a, c.b, c.c) {
			t.FailNow()
		}
	}
}

func Test1550(t *testing.T) {
	for _, c := range []struct {
		rst bool
		arr []int
	}{
		{false, []int{2, 6, 4, 1}},
		{true, []int{1, 2, 34, 3, 4, 5, 7, 23, 12}},

		{true, []int{1, 3, 5}},  // edge cases
		{false, []int{1, 2, 3}}, // edge cases

		{false, []int{1}}, // boundary
	} {
		if c.rst != threeConsecutiveOdds(c.arr) {
			t.FailNow()
		}
	}
}

// 1752 Check If Array Is Sorted and Rotated
func Test1752(t *testing.T) {
	log.Print("true ?= ", check([]int{3, 4, 5, 1, 2}))
	log.Print("false ?= ", check([]int{2, 1, 3, 4}))
	log.Print("true ?= ", check([]int{1, 2, 3}))
}

func Test1920(t *testing.T) {
	for _, c := range []struct {
		rst, nums []int
	}{
		{[]int{0, 1, 2, 4, 5, 3}, []int{0, 2, 1, 5, 3, 4}},
		{[]int{4, 5, 0, 1, 2, 3}, []int{5, 0, 1, 2, 3, 4}},
	} {
		if !reflect.DeepEqual(c.rst, buildArray(c.nums)) {
			t.FailNow()
		}
	}
}

func Test2200(t *testing.T) {
	for _, c := range []struct {
		rst, nums []int
		key, k    int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{3, 4, 9, 1, 3, 9, 5}, 9, 1},
		{[]int{0, 1, 2, 3, 4}, []int{2, 2, 2, 2, 2}, 2, 2},
	} {
		log.Print("* ", c.nums, c.key, c.k)
		if !reflect.DeepEqual(c.rst, findKDistantIndices(c.nums, c.key, c.k)) {
			t.FailNow()
		}
	}
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

func Test2094(t *testing.T) {
	for _, c := range []struct {
		rst    []int
		digits []int
	}{
		{[]int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}, []int{2, 1, 3, 0}},
		{[]int{222, 228, 282, 288, 822, 828, 882}, []int{2, 2, 8, 8, 2}},
		{[]int{}, []int{3, 7, 5}},
	} {
		if !reflect.DeepEqual(c.rst, findEvenNumbers(c.digits)) {
			t.FailNow()
		}
	}
}

func Test2145(t *testing.T) {
	Optimized := func(differences []int, lower, upper int) int {
		x, n := 0, 0
		v := 0
		for _, d := range differences {
			v += d
			x, n = max(v, x), min(v, n)
			if x-n > upper-lower {
				return 0
			}
		}

		return upper - lower - (x - n) + 1
	}

	for _, c := range []struct {
		rst          int
		differences  []int
		lower, upper int
	}{
		{2, []int{1, -3, 4}, 1, 6},
		{4, []int{3, -4, 5, 1, -2}, -4, 5},
		{0, []int{4, -7, 2}, 3, 6},
	} {
		if c.rst != numberOfArrays(c.differences, c.lower, c.upper) {
			t.FailNow()
		}
		log.Printf(":: %d ~ %d", c.rst, Optimized(c.differences, c.lower, c.upper))
	}
}

func Test2176(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
		k    int
	}{
		{4, []int{3, 1, 2, 2, 2, 1, 3}, 2},
		{0, []int{1, 2, 3, 4}, 1},
	} {
		if c.rst != countPairs_Divisible(c.nums, c.k) {
			t.FailNow()
		}
	}
}

func Test2302(t *testing.T) {
	for _, c := range []struct {
		rst  int64
		nums []int
		k    int64
	}{
		{int64(6), []int{2, 1, 4, 3, 5}, int64(10)},
		{int64(5), []int{1, 1, 1}, int64(5)},
	} {
		if c.rst != countSubarrays_KScore(c.nums, c.k) {
			t.FailNow()
		}
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
	// Maximize: (N_i - N_j) * N_k for any i < j < k

	Optimized := func(nums []int) int64 {
		leftMax := nums[0]

		rightMaxs := make([]int, len(nums))
		rightMaxs[len(nums)-1] = nums[len(nums)-1]
		for r := len(nums) - 2; r >= 2; r-- {
			rightMaxs[r] = max(nums[r], rightMaxs[r+1])
		}

		log.Print("-> ", leftMax, rightMaxs)

		xVal := int64(0)
		for i, n := range nums[1 : len(nums)-2] {
			xVal = max(xVal, int64(leftMax-n)*int64(rightMaxs[i+1]))
			leftMax = max(n, leftMax)
		}

		return xVal
	}

	SpaceOptimized := func(nums []int) int64 {
		xVal := int64(0)
		leftMax, diffMax := 0, 0

		for _, n := range nums {
			xVal = max(int64(diffMax)*int64(n), xVal)

			diffMax = max(leftMax-n, diffMax)
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
		for _, f := range []func([]int) int64{maximumTripletValue, Optimized, SpaceOptimized} {
			if c.rst != f(c.nums) {
				t.FailNow()
			}
		}
		log.Printf(":: %d   <- %v", c.rst, c.nums)
	}
}

func Test2874(t *testing.T) {

	for _, c := range []struct {
		rst  int64
		nums []int
	}{
		{77, []int{12, 6, 1, 2, 7}},
		{133, []int{1, 10, 3, 4, 19}},
		{0, []int{1, 2, 3}},
	} {
		if c.rst != maximumTripletValueII(c.nums) {
			t.FailNow()
		}
		log.Printf(":: %d   <- %v", c.rst, c.nums)
	}
}

func Test2942(t *testing.T) {
	for _, c := range []struct {
		rst   []int
		words []string
		x     byte
	}{
		{[]int{0, 1}, []string{"leet", "code"}, 'e'},
		{[]int{0, 2}, []string{"abc", "bcd", "aaaa", "cbc"}, 'a'},
		{[]int{}, []string{"abc", "bcd", "aaaa", "cbc"}, 'z'},
	} {
		if !reflect.DeepEqual(c.rst, findWordsContaining(c.words, c.x)) {
			t.FailNow()
		}
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

func Test3355(t *testing.T) {
	for _, c := range []struct {
		rst     bool
		nums    []int
		queries [][]int
	}{
		{true, []int{1, 0, 1}, [][]int{{0, 2}}},
		{false, []int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}},
	} {
		if c.rst != isZeroArray(c.nums, c.queries) {
			t.FailNow()
		}
	}
}

func Test3392(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{1, []int{1, 2, 1, 4, 1}},
		{0, []int{1, 1, 1}},
	} {
		if c.rst != countSubarrays_Length3(c.nums) {
			t.FailNow()
		}
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

func Test3423(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{3, []int{1, 2, 4}},
		{5, []int{-5, -10, -5}},
	} {
		if c.rst != maxAdjacentDistance(c.nums) {
			t.FailNow()
		}
	}
}
