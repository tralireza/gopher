package gopher

import (
	"log"
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

// 135m Candy
func Test135(t *testing.T) {
	log.Print("5 ?= ", candy([]int{1, 0, 2}))
	log.Print("4 ?= ", candy([]int{1, 2, 2}))
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

// 918m Maximum Sum Circular Subarray
func Test918(t *testing.T) {
	log.Print("3 ?= ", maxSubarraySumCircular([]int{1, -2, 3, -2}))
	log.Print("10 ?= ", maxSubarraySumCircular([]int{5, -3, 5}))
	log.Print("-2 ?= ", maxSubarraySumCircular([]int{-3, -2, -3}))
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

// 2202m Maximize the Topmost Element After K Moves
func Test2202(t *testing.T) {
	log.Print("5 ?= ", maximumTop([]int{5, 2, 2, 4, 0, 6}, 4))
	log.Print("-1 ?= ", maximumTop([]int{2}, 1))
}

// 2280m Minimum Lines to Represent a Line Chart
func Test2280(t *testing.T) {
	log.Print("3 ?= ", minimumLines([][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}))
	log.Print("1 ?= ", minimumLines([][]int{{3, 4}, {1, 2}, {7, 8}, {2, 3}}))
}

// 2938m Separate Black and White Balls
func Test2938(t *testing.T) {
	log.Print("1 ?= ", minimumSteps("101"))
	log.Print("2 ?= ", minimumSteps("100"))
	log.Print("0 ?= ", minimumSteps("0111"))
}
