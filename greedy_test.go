package gopher

import (
	"log"
	"reflect"
	"slices"
	"testing"
)

// 11m Container With Most Water
func Test11(t *testing.T) {
	log.Print("49 ?= ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	log.Print("1 ?= ", maxArea([]int{1, 1}))
}

// 15m 3Sum
func Test15(t *testing.T) {
	log.Print("[[-1 -1 2] [-1 0 1]] ?= ", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// 36m Valid Sudoku
func Test36(t *testing.T) {
	log.Print("true ?= ", isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}

// 53m Maximum Subarray
func Test53(t *testing.T) {
	log.Print("6 ?= ", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	log.Print("1 ?= ", maxSubArray([]int{1}))
	log.Print("23 ?= ", maxSubArray([]int{5, 4, -1, 7, 8}))
}

// 134m Gas Station
func Test134(t *testing.T) {
	log.Print("3 ?= ", canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	log.Print("-1 ?= ", canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}

func Test135(t *testing.T) {
	for _, c := range []struct {
		rst     int
		ratings []int
	}{
		{5, []int{1, 0, 2}},
		{4, []int{1, 2, 2}},
	} {
		if c.rst != candy(c.ratings) {
			t.FailNow()
		}
	}
}

// 149h Max Points on a Line
func Test149(t *testing.T) {
	log.Print("3 ?= ", maxPointsOnLine([][]int{{1, 1}, {2, 2}, {3, 3}}))
	log.Print("4 ?= ", maxPointsOnLine([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}

// 167m Two Sum II - Input Array Is Sorted
func Test167(t *testing.T) {
	log.Print("[1 2] ?= ", twoSum([]int{2, 7, 11, 15}, 9))
	log.Print("[1 3] ?= ", twoSum([]int{2, 3, 4}, 6))
	log.Print("[1 2] ?= ", twoSum([]int{-1, 0}, -1))
}

// 670m Maximum Swap
func Test670(t *testing.T) {
	BruteForce := func(num int) int {
		xVal := num

		D := []int{}
		for num > 0 {
			D = append(D, num%10)
			num /= 10
		}
		slices.Reverse(D)
		log.Print(" -> ", D)

		for i := 0; i < len(D)-1; i++ {
			for j := i + 1; j < len(D); j++ {
				if D[i] < D[j] {
					D[i], D[j] = D[j], D[i]
					v := 0
					for _, d := range D {
						v = 10*v + d
					}
					xVal = max(v, xVal)
					D[i], D[j] = D[j], D[i]
				}
			}
		}

		return xVal
	}

	for _, fn := range []func(int) int{BruteForce, maximumSwap} {
		log.Print("7236 ?= ", fn(2736))
		log.Print("9973 ?= ", fn(9973))
		log.Print("--")
	}
}

func Test781(t *testing.T) {
	// 1 <= N <= 1000, 0 <= N_i < 1000

	for _, c := range []struct {
		rst     int
		answers []int
	}{
		{5, []int{1, 1, 2}},
		{11, []int{10, 10, 10}},
	} {
		if c.rst != numRabbits(c.answers) {
			t.FailNow()
		}
	}
}

// 918m Maximum Sum Circular Subarray
func Test918(t *testing.T) {
	log.Print("3 ?= ", maxSubarraySumCircular([]int{1, -2, 3, -2}))
	log.Print("10 ?= ", maxSubarraySumCircular([]int{5, -3, 5}))
	log.Print("-2 ?= ", maxSubarraySumCircular([]int{-3, -2, -3}))
}

func Test1007(t *testing.T) {
	for _, c := range []struct {
		rst           int
		tops, bottoms []int
	}{
		{2, []int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}},
		{-1, []int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}},
	} {
		if c.rst != minDominoRotations(c.tops, c.bottoms) {
			t.FailNow()
		}
	}
}

// 1605m Find Valid Matrix Given Row and Column Sums
func Test1605(t *testing.T) {
	Optimized := func(rowSum []int, colSum []int) [][]int {
		M := make([][]int, len(rowSum))
		for r := range M {
			M[r] = make([]int, len(colSum))

		}

		r, c := 0, 0
		for r < len(rowSum) && c < len(colSum) {
			M[r][c] = min(rowSum[r], colSum[c])

			rowSum[r] -= M[r][c]
			colSum[c] -= M[r][c]

			if rowSum[r] == 0 {
				r++
			}
			if colSum[c] == 0 {
				c++
			}
		}

		return M
	}

	for _, f := range []func([]int, []int) [][]int{restoreMatrix, Optimized} {
		log.Print("[[3 0] [1 7]] ?= ", f([]int{3, 8}, []int{4, 7}))
		log.Print("[[0 5 0] [6 1 0] [2 0 8]] ?= ", f([]int{5, 7, 10}, []int{8, 6, 8}))
		log.Print("--")
	}
}

func Test2131(t *testing.T) {
	for _, c := range []struct {
		rst   int
		words []string
	}{
		{6, []string{"lc", "cl", "gg"}},
		{8, []string{"ab", "ty", "yt", "lc", "cl", "ab"}},
		{2, []string{"cc", "ll", "xx"}},
	} {
		if c.rst != longestPalindrome(c.words) {
			t.FailNow()
		}
	}
}

// 2202m Maximize the Topmost Element After K Moves
func Test2202(t *testing.T) {
	log.Print("5 ?= ", maximumTop([]int{5, 2, 2, 4, 0, 6}, 4))
	log.Print("-1 ?= ", maximumTop([]int{2}, 1))
}

// 2280m Minimum Lines to Represent a Line Chart
func Test2280(t *testing.T) {
	for _, c := range []struct {
		rst         int
		stockPrices [][]int
	}{
		{3, [][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}},
		{1, [][]int{{3, 4}, {1, 2}, {7, 8}, {2, 3}}},
	} {
		log.Print("* ", c.stockPrices)
		if c.rst != minimumLines(c.stockPrices) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test2434(t *testing.T) {
	for _, c := range []struct {
		rst, s string
	}{
		{"azz", "zza"},
		{"abc", "bac"},
		{"addb", "bdda"},

		{"aabbcuztdc", "cdatazbubc"},
		{"eekstrlpmomwzqummz", "mmuqezwmomeplrtskz"},
	} {
		log.Print("* ", c.s)
		if c.rst != robotWithString(c.s) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test2900(t *testing.T) {
	for _, c := range []struct {
		rst, words []string
		groups     []int
	}{
		{[]string{"e", "b"}, []string{"e", "a", "b"}, []int{0, 0, 1}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, []int{1, 0, 1, 1}},

		{
			[]string{"iy", "gh", "e", "pc", "t", "c", "r", "l", "y", "ok", "w", "jq", "p", "tt", "mg", "vt", "to", "q", "fs", "nh", "o", "i", "d"},
			[]string{"iy", "gh", "e", "pc", "a", "j", "t", "g", "c", "n", "r", "v", "m", "ub", "l", "y", "ok", "w", "z", "gg", "jq", "p", "tt", "x", "mg", "vt", "to", "k", "q", "fs", "nh", "o", "i", "d", "b", "u", "kd", "s"},
			[]int{1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1},
		},
	} {
		if !reflect.DeepEqual(c.rst, getLongestSubsequence(c.words, c.groups)) {
			t.FailNow()
		}
	}
}

func Test2918(t *testing.T) {
	for _, c := range []struct {
		rst          int64
		nums1, nums2 []int
	}{
		{12, []int{3, 2, 0, 1, 0}, []int{6, 5, 0}},
		{-1, []int{2, 0, 2, 0}, []int{1, 4}},

		{139, []int{0, 16, 28, 12, 10, 15, 25, 24, 6, 0, 0}, []int{20, 15, 19, 5, 6, 29, 25, 8, 12}},
	} {
		if c.rst != minSum(c.nums1, c.nums2) {
			t.FailNow()
		}
	}
}

// 2938m Separate Black and White Balls
func Test2938(t *testing.T) {
	for _, c := range []struct {
		rst int64
		s   string
	}{
		{int64(1), "101"},
		{int64(2), "100"},
		{int64(0), "0111"},
	} {
		if c.rst != minimumSteps(c.s) {
			t.FailNow()
		}
	}
}

func Test3170(t *testing.T) {
	for _, c := range []struct {
		rst, s string
	}{
		{"aab", "aaba*"},
		{"abc", "abc"},
	} {
		log.Print("* ", c.s)
		if c.rst != clearStars(c.s) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}
