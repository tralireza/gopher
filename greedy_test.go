package gopher

import (
	"log"
	"testing"
)

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