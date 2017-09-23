package gopher

// 1605m Find Valid Matrix Given Row and Column Sums
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	M := make([][]int, len(rowSum))
	for r := range M {
		M[r] = make([]int, len(colSum))
	}

	for r := 0; r < len(rowSum); r++ {
		for c := 0; c < len(colSum); c++ {
			mVal := rowSum[r]
			if colSum[c] < mVal {
				mVal = colSum[c]
			}

			M[r][c] = mVal

			rowSum[r] -= mVal
			colSum[c] -= mVal
		}
	}

	return M
}
