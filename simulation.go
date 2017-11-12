package gopher

import (
	"log"
	"slices"
)

// 885m Spiral Matrix III
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	M := [][]int{{rStart, cStart}}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0 // 0: ->, 1: v, 2: <-, 3: ^

	r, c := rStart, cStart
	rD, cD := 1, 1
	for len(M) < rows*cols {
		switch dir {
		case 0, 2: // -> | <-
			for range rD {
				c += dirs[dir][1]
				if 0 <= c && c < cols && 0 <= r && r < rows {
					M = append(M, []int{r, c})
				}
			}
			rD++

		default: // v | ^
			for range cD {
				r += dirs[dir][0]
				if 0 <= c && c < cols && 0 <= r && r < rows {
					M = append(M, []int{r, c})
				}
			}
			cD++
		}

		dir++
		dir %= 4
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
