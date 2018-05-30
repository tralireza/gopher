package gopher

import (
	"fmt"
	"log"
	"slices"
	"strconv"
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

// 68h Text Justification
func fullJustify(words []string, maxWidth int) []string {
	Pack := func() [][]string {
		L := [][]string{}
		llen, line := 0, []string{}
		for _, w := range words {
			if len(w) <= maxWidth-llen-len(line) {
				line = append(line, w)
				llen += len(w)
			} else {
				L = append(L, line)
				llen, line = len(w), []string{w}
			}
		}
		if len(line) > 0 {
			L = append(L, line)
		}
		return L
	}

	L := Pack()
	log.Print(L)

	J := []string{}
	for i, line := range L {
		if i == len(L)-1 || len(line) == 1 { // last line or one word line are only left-justified
			llen := 0
			for _, w := range line {
				llen += len(w)
			}
			for llen < maxWidth-len(line)+1 { // left-justify last word of line
				line[len(line)-1] += " "
				llen++
			}
		} else { // middle lines -> fully justified
			llen := 0
			for _, w := range line {
				llen += len(w)
			}
			p := 0
			for llen < maxWidth-len(line)+1 {
				line[p] += " "
				llen++
				p++
				if p == len(line)-1 {
					p = 0
				}
			}
		}
		J = append(J, strings.Join(line, " "))
	}
	return J
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
				if l < 2 || l > 3 {
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

// 566 Reshape the Matrix
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	R := [][]int{}
	i := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if i%c == 0 {
				R = append(R, make([]int, c))
			}
			R[i/c][i%c] = mat[x][y]
			i++
		}
	}
	return R
}

// 592m Fraction Addition and Subtraction
func fractionAddition(expression string) string {
	gcd := func(a, b int) int {
		if b > a {
			a, b = b, a
		}
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	i := 0

	e := expression
	Value := func() int {
		v := 0
		for ; i < len(e) && e[i] <= '9' && e[i] >= '0'; i++ {
			v = 10*v + int(expression[i]-'0')
		}
		return v
	}

	n, d := 1, 1
	if e[0] == '-' {
		d *= -1
		i++
	}

	for i < len(e) {
		switch e[i] {
		case '+':
			i++ // Skip +
			N := Value()
			i++ // Skip /
			D := Value()

			n, d = n*D+N*d, d*D
			r := gcd(n, d)
			n, d = n/r, d/r

		case '-':
			i++ // Skip -
			N := Value()
			i++ // Skip /
			D := Value()

			n, d = n*D-N*d, d*D
			r := gcd(n, d)
			n, d = n/r, d/r

		default: // first n/d
			n *= Value()
			i++
			d *= Value()
		}
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	expOut := strconv.Itoa(abs(n)) + "/" + strconv.Itoa(abs(d))
	if n*d < 0 {
		return "-" + expOut
	}
	return expOut
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

// 860 Lemonade Change
func lemonadeChange(bills []int) bool {
	f, t := 0, 0 // fives, tens

	for _, b := range bills {
		switch b {
		case 5:
			f++
		case 10:
			if f > 0 {
				f--
			} else {
				return false
			}
			t++
		case 20:
			if f > 0 && t > 0 {
				f--
				t--
			} else if f > 2 {
				f -= 3
			} else {
				return false
			}
		}
	}
	return true
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

// 1945 Sum of Digits of String After Convert
func getLucky(s string, k int) int {
	t := ""
	for i := 0; i < len(s); i++ {
		t += fmt.Sprintf("%d", s[i]-'a'+1)
	}

	for range k {
		n := 0
		for i := 0; i < len(t); i++ {
			n += int(t[i] - '0')
		}
		t = fmt.Sprintf("%d", n)
	}

	n := 0
	fmt.Sscanf(t, "%d", &n)
	return n
}

// 2022 Convert 1D Array Into 2D Array
func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	R := [][]int{}
	for i, v := range original {
		if i%n == 0 {
			R = append(R, make([]int, n))
		}
		R[i/n][i%n] = v
	}
	return R
}

// 2028m Find Missing Observations
func missingRolls(rolls []int, mean int, n int) []int {
	tSum := (len(rolls) + n) * mean
	for _, v := range rolls {
		tSum -= v
	}

	if tSum < 1*n || 6*n < tSum {
		return []int{} // Impossible
	}

	R := make([]int, n)
	for i := range R {
		R[i] = tSum / n
	}

	i := 0
	for range tSum % n {
		R[i]++
		i++
	}

	return R
}

// 2419m Longest Subarray With Maximum Bitwise AND
func longestSubarray(nums []int) int {
	xVal, longest, cur := 0, 0, 0

	for _, n := range nums {
		if xVal < n {
			xVal = n
			longest, cur = 1, 1
		} else if xVal == n {
			cur++
			longest = max(cur, longest)
		} else {
			cur = 0
		}
	}

	return longest
}

// 2696 Minimum String Length After Removing Substrings
func minLength(s string) int {
	Q := []rune{}

	for _, c := range s {
		if len(Q) > 0 &&
			(Q[len(Q)-1] == 'A' && c == 'B' || Q[len(Q)-1] == 'C' && c == 'D') {
			Q = Q[:len(Q)-1]
			continue
		}
		Q = append(Q, c)
	}

	return len(Q)
}

// 3001m Minimum Moves to Capture the Queen
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	Cross := func(dirs [][]int, tgX, tgY, bkX, bkY int) int {
		for _, d := range dirs {
			x, y := e+d[0], f+d[1]
			for x < 9 && x > 0 && y < 9 && y > 0 && !(x == bkX && y == bkY) { // && ... and other piece is not in the way
				if x == tgX && y == tgY {
					return 1 // Rook|Bishop hits...
				}
				x, y = x+d[0], y+d[1]
			}
		}
		return -1
	}

	if Cross([][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}, a, b, c, d) == 1 { // Rook
		return 1
	}
	if Cross([][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}, c, d, a, b) == 1 { // Bishop
		return 1
	}
	return 2
}

// 3270 Find the Key of the Numbers
func generateKey(num1 int, num2 int, num3 int) int {
	k := 0

	rdx := 1
	for range 4 {
		d := min(num1%10, num2%10, num3%10)
		k += d * rdx
		rdx *= 10

		num1 /= 10
		num2 /= 10
		num3 /= 10
	}

	return k
}

// 3274 Check if Two Chessboard Squares Have the Same Color
func checkTwoChessboards(coordinate1, coordinate2 string) bool {
	Chess := [8][8]int{}
	for c := 1; c < 8; c++ {
		Chess[0][c] = Chess[0][c-1] ^ 1
	}
	for r := 1; r < 8; r++ {
		for c := 0; c < 8; c++ {
			Chess[r][c] = Chess[r-1][c] ^ 1
		}
	}

	return Chess[7-coordinate1[1]+'1'][coordinate1[0]-'a'] ==
		Chess[7-coordinate2[1]+'1'][coordinate2[0]-'a']
}

// 3304 Find the K-th Character in String Game I
func kthCharacter(k int) byte {
	s := []byte{'a'}
	for len(s) < k {
		for i := range len(s) {
			s = append(s, 'a'+(s[i]-'a'+1)%26)
		}
	}
	return s[k-1]
}
