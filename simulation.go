package gopher

import (
	"log"
	"slices"
	"strings"
)

// 6m Zigzag Conversion
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	Z := make([][]byte, numRows)

	i, r := 0, 0
	for i < len(s) {
		for _, dir := range []int{1, -1} {
			for range numRows - 1 {
				if i < len(s) {
					Z[r] = append(Z[r], s[i])
				}
				i++
				r += dir
			}
		}
	}

	S := []string{}
	for _, z := range Z {
		S = append(S, string(z))
	}
	return strings.Join(S, "")
}

// 289m Game of Life
func gameOfLife(board [][]int) {
	log.Print(" <- ", board)

	Rows, Cols := len(board), len(board[0])

	lCells := func(r, c int) int {
		l := 0
		for _, x := range []int{r - 1, r, r + 1} {
			for _, y := range []int{c - 1, c, c + 1} {
				if x == r && y == c {
					continue
				}

				if x >= 0 && x < Rows && y >= 0 && y < Cols {
					if board[x][y] == 1 || board[x][y] == -9 {
						l++
					}
				}
			}
		}
		return l
	}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			l := lCells(r, c)
			switch board[r][c] {
			case 0:
				if l == 3 {
					board[r][c] = -1 // new Live
				}
			case 1:
				if l < 2 && l > 3 {
					board[r][c] = -9 // new Dead
				}
			}
		}
	}

	log.Print(board)

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if board[r][c] == -1 {
				board[r][c] = 1
			}
			if board[r][c] == -9 {
				board[r][c] = 0
			}
		}
	}

	log.Print(" -> ", board)
}

// 840m Magic Squares In Grid
func numMagicSquaresInside(grid [][]int) int {
	t := 0

	Check := func(r, c int) bool {
		M := make([]bool, 9+1) // 1..9 uniqueness

		vSums := [3]int{} // vertical sums, ie: columns
		for _, r := range []int{r - 1, r, r + 1} {
			hSum := 0 // horizontal sum, ie: row
			for j, c := range []int{c - 1, c, c + 1} {
				v := grid[r][c]
				if v < 1 || v > 9 || M[v] {
					return false
				}
				M[v] = true

				hSum += v
				vSums[j] += v
			}
			if hSum != 15 {
				return false
			}
		}
		for _, vSum := range vSums {
			if vSum != 15 {
				return false
			}
		}

		dSums := [2]int{} // diagonal sums, ie: \ /
		for _, d := range []int{-1, 0, 1} {
			dSums[0] += grid[r+d][c+d]
			dSums[1] += grid[r-d][c+d]
		}
		for _, dSum := range dSums {
			if dSum != 15 {
				return false
			}
		}

		return true
	}

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			if grid[r][c] == 5 && Check(r, c) {
				t++
			}
		}
	}

	return t
}

// 885m Spiral Matrix III
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	M := [][]int{{rStart, cStart}}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0 // 0: >, 1: v, 2: <, 3: ^

	r, c := rStart, cStart
	steps := 1
	for len(M) < rows*cols {
		for range 2 {
			for range steps {
				r += dirs[dir][0]
				c += dirs[dir][1]

				if r < rows && c < cols && 0 <= r && 0 <= c {
					M = append(M, []int{r, c})
				}
			}
			dir++
		}
		dir %= 4
		steps++
	}

	return M
}

// 1380 Lucky Numbers in a Matrix
func luckyNumbers(matrix [][]int) []int {
	// 1 <= matrix[i][j] <= 10^5
	Mr, Xc := make([]int, len(matrix)), make([]int, len(matrix[0]))

	for r := 0; r < len(matrix); r++ {
		Mr[r] = slices.Min(matrix[r])
	}

	for c := 0; c < len(matrix[0]); c++ {
		Xc[c] = matrix[0][c]
		for r := 1; r < len(matrix); r++ {
			Xc[c] = max(matrix[r][c], Xc[c])
		}
	}

	log.Print(Mr)
	log.Print(Xc)

	R := []int{}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == Mr[r] && matrix[r][c] == Xc[c] {
				R = append(R, matrix[r][c])
			}
		}
	}
	return R
}
