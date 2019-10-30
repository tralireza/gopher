package gopher

import (
	"fmt"
	"log"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

// 17m Letter Combinations of a Phone Number
func letterCombinations(digits string) []string {
	R := []string{}
	if len(digits) == 0 {
		return R
	}

	Mem := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var r []byte

	var BT func(int)
	BT = func(start int) {
		if start == len(digits) {
			R = append(R, string(r))
			return
		}

		d := digits[start] - '0'
		for i := start; i < len(Mem[d]); i++ {
			r = append(r, Mem[d][i])
			BT(start + 1)
			r = r[:len(r)-1]
		}
	}

	BT(0)

	return R
}

// 40m Combination Sum II
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	R := [][]int{}

	var v []int
	var Search func(start, cSum int)
	Search = func(start, cSum int) {
		if cSum == target {
			R = append(R, append([]int{}, v...))
			return
		}

		// Prune
		if cSum > target {
			return
		}
		if start == len(candidates) {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] { // Group & Prune
				continue
			}
			v = append(v, candidates[i])
			Search(i+1, cSum+candidates[i])
			v = v[:len(v)-1]
		}
	}

	Search(0, 0)

	return R
}

// 46m Permutations
func permute(nums []int) [][]int {
	R := [][]int{}

	var r []int
	var BT func(int)
	BT = func(start int) {
		if start == len(nums) {
			R = append(R, append([]int{}, r...))
			return
		}

		for i := start; i < len(nums); i++ {
			r[start], r[i] = r[i], r[start]
			BT(start + 1)
			r[start], r[i] = r[i], r[start]
		}
	}

	r = append(r, nums...)
	BT(0)

	return R
}

// 50m Pow(x,n)
func myPow(x float64, n int) float64 {
	power := float64(1)

	e := n
	if n < 0 {
		e = -e
	}

	for e > 0 {
		if e&1 == 1 {
			power *= x
		}
		x *= x
		e >>= 1
	}

	if n < 0 {
		return 1.0 / power
	}
	return power
}

// 52h N-Queens II
func totalNQueens(n int, f func(i int, B [][]byte)) int {
	B := make([][]byte, n)
	for r := range B {
		B[r] = slices.Repeat([]byte{'~'}, n)
	}

	Dir := []int{1, -1, -1, 1, 1} // diagonal movement...

	t := 0

	var Queen func(int)
	Queen = func(r int) {
		if r == n {
			t++
			log.Printf("*** -> %s", B)
			f(t, B)
			return
		}

		for c := 0; c < n; c++ {
			B[r][c] = 'Q' // place a Queen!

			q := 0
			for x := 0; x < n; x++ {
				if B[x][c] == 'Q' {
					q++
				}
			}
			for d := range 4 { // diagonal...
				r, c := r, c
				for range n - 1 {
					r += Dir[d]
					c += Dir[d+1]
					if r >= 0 && r < n && c >= 0 && c < n && B[r][c] == 'Q' {
						q++
					}
				}
			}

			if q == 1 { // only 1 Queen roaming this realm!
				Queen(r + 1) // go to next Row
			}

			B[r][c] = '~' // BackTrack :: remove the Queen
		}
	}

	Queen(0) // Row: 0

	return t
}

// 77m Combinations
func combine(n int, k int) [][]int {
	R := [][]int{}
	var r []int

	var Choose func(int)
	Choose = func(start int) {
		if len(r) == k {
			R = append(R, append([]int{}, r...))
			return
		}

		for i := start; i <= n; i++ {
			r = append(r, i)
			Choose(i + 1)
			r = r[:len(r)-1]
		}
	}

	Choose(1)

	return R
}

// 224h Basic Calculator
func calculate(s string) int {
	i := 0
	s = strings.Replace(s, " ", "", -1) // Noise-reduction!

	Value := func() int {
		v := 0
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			v = 10*v + int(s[i]-'0')
		}
		return v
	}

	var Calc func() int
	Calc = func() int {
		var v int
		for i < len(s) {
			switch s[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				v = Value()

			case '(':
				i++
				v = Calc()
			case ')':
				i++
				return v

			case '+':
				i++
				switch s[i] {
				case '(':
					i++
					v += Calc()
				default:
					v += Value()
				}
			case '-':
				i++
				switch s[i] {
				case '(':
					i++
					v -= Calc()
				default:
					v -= Value()
				}
			}
		}
		return v
	}

	return Calc()
}

// 282h Expression Add Operators
func addOperators(num string, target int) []string {
	R := []string{}

	var Search func(start, prvOpr, curVal int, expr string)
	Search = func(start, prvOpr, curVal int, expr string) {
		if start == len(num) {
			if curVal == target {
				R = append(R, expr)
			}
			return
		}

		for i := start; i < len(num); i++ {
			if num[start] == '0' && i > start {
				break
			}

			v, _ := strconv.Atoi(num[start : i+1])
			if i == 0 {
				Search(i+1, v, curVal+v, expr+num[start:i+1])
			} else {
				Search(i+1, v, curVal+v, expr+"+"+num[start:i+1])
				Search(i+1, -v, curVal-v, expr+"-"+num[start:i+1])
				Search(i+1, prvOpr*v, curVal-prvOpr+prvOpr*v, expr+"*"+num[start:i+1])
			}
		}
	}

	Search(0, 0, 0, "")

	return R
}

// 301h Remove Invalid Parentheses
func removeInvalidParentheses(s string) []string {
	Picks := make([]bool, len(s))
	M := map[int]map[string]struct{}{}

	var Search func(start, opens, closes int)
	Search = func(start, opens, closes int) {
		if start == len(s) {
			bfr := []byte{}
			for i := 0; i < len(s); i++ {
				if Picks[i] {
					bfr = append(bfr, s[i])
				}
			}

			Valid := func(bfr []byte) bool {
				S := make([]byte, 0, len(bfr))
				for i := 0; i < len(bfr); i++ {
					switch bfr[i] {
					case '(':
						S = append(S, '(')
					case ')':
						if len(S) == 0 {
							return false
						}
						S = S[:len(S)-1]
					}
				}
				return len(S) == 0
			}

			FastValid := func(bfr []byte) bool {
				counterStack := 0
				for i := 0; i < len(bfr); i++ {
					switch bfr[i] {
					case '(':
						counterStack++
					case ')':
						counterStack--
						if counterStack < 0 {
							return false
						}
					}
				}
				return counterStack == 0
			}

			if Valid(bfr) && FastValid(bfr) {
				if _, ok := M[len(bfr)]; !ok {
					M[len(bfr)] = map[string]struct{}{}
				}
				M[len(bfr)][string(bfr)] = struct{}{}
			}

			return
		}

		switch s[start] {
		case '(':
			Search(start+1, opens, closes)

			opens++
			if opens <= len(s)-start+closes {
				Picks[start] = true
				Search(start+1, opens, closes)
				Picks[start] = false
			}

		case ')':
			Search(start+1, opens, closes)

			closes++
			if closes <= opens {
				Picks[start] = true
				Search(start+1, opens, closes)
				Picks[start] = false
			}

		default:
			Picks[start] = true
			Search(start+1, opens, closes)
		}
	}

	Search(0, 0, 0)

	log.Printf("-> %v", M)

	lMax := 0
	for l := range M {
		lMax = max(l, lMax)
	}

	return slices.Collect(maps.Keys(M[lMax]))
}

// 386m Lexicographical Numbers
func lexicalOrder(n int) []int {
	R := []int{}

	v := 1
	for range n {
		R = append(R, v)

		if v*10 <= n {
			v *= 10
		} else {
			for v%10 == 9 || v >= n {
				v /= 10
			}
			v++
		}
	}

	return R
}

// 427m Construct Quad Tree
type Node427 struct {
	Val, IsLeaf                                bool
	TopLeft, TopRight, BottomLeft, BottomRight *Node427
}

func construct(grid [][]int) *Node427 {
	N := len(grid)
	if N == 1 {
		return &Node427{Val: grid[0][0] == 1, IsLeaf: true}
	}

	n := &Node427{}

	for i, r := range []int{0, N / 2} {
		for j, c := range []int{0, N / 2} {
			G := [][]int{}
			for i := range N / 2 {
				G = append(G, grid[r+i][c:c+N/2])
			}

			switch i<<1 + j {
			case 0:
				n.TopLeft = construct(G)
			case 1:
				n.TopRight = construct(G)
			case 2:
				n.BottomLeft = construct(G)
			case 3:
				n.BottomRight = construct(G)
			}
		}
	}

	Children := []*Node427{n.TopLeft, n.TopRight, n.BottomLeft, n.BottomRight}
	n.IsLeaf = true
	for _, c := range Children {
		n.IsLeaf = n.IsLeaf && c.IsLeaf
	}
	if n.IsLeaf { // Set|Reset Val for this Leaf Node ...
		n.Val = true
		for _, c := range Children {
			n.Val = n.Val && c.Val
		}
		if n.Val { // :: all Child.Val: true
			return &Node427{IsLeaf: true, Val: true}
		}

		n.Val = false
		for _, c := range Children {
			n.Val = n.Val || c.Val
		}
		if !n.Val { // :: all Child.Val: false
			return &Node427{IsLeaf: true}
		}
	}

	n.IsLeaf, n.Val = false, true
	return n
}

// 650m 2 Keys Keyboard
func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	rCalls, Mem := 0, map[[2]int]int{}
	ops := n

	var CopyPaste func(l, lp int) int
	CopyPaste = func(l, lp int) int {
		if l >= n {
			if l == n {
				return 0
			}
			return n
		}

		rCalls++
		if v, ok := Mem[[2]int{l, lp}]; ok {
			return v
		}

		p := CopyPaste(l+lp, lp) // just Paste
		cp := CopyPaste(l+l, l)  // Copy & Paste

		Mem[[2]int{l, lp}] = min(1+p, 2+cp)
		return Mem[[2]int{l, lp}]
	}

	ops = 1 + CopyPaste(1, 1)
	log.Print(" -> ", rCalls, " # rCalls")

	return ops
}

// 1079m Letter Tile Possibilities
func numTilePossibilities(tiles string) int {
	M := map[string]struct{}{}

	Used := make([]bool, len(tiles))

	var GenCs func(tile []byte)
	GenCs = func(tile []byte) {
		if _, ok := M[string(tile)]; !ok {
			M[string(tile)] = struct{}{}
		}

		for p := range len(tiles) {
			if !Used[p] {
				Used[p] = true
				tile = append(tile, tiles[p])

				GenCs(tile)

				Used[p] = false
				tile = tile[:len(tile)-1]
			}
		}
	}

	GenCs([]byte{})

	log.Print(" -> ", M)

	return len(M) - 1
}

// 1106h Parsing a Boolean Expression
func parseBoolExpr(expression string) bool {
	p := 0

	var Parse func() bool
	Parse = func() bool {
		log.Printf(" -> %q", expression[p])

		switch expression[p] {
		case 't':
			p++
			return true
		case 'f':
			p++
			return false
		case '!': // !(expr)
			p += 2
			v := !Parse()
			p++
			return v
		default: // &(expr[,expr]), |(expr[,expr])
			andOr := expression[p] == '&'
			Vals := []bool{}

			p += 2
			for expression[p] != ')' {
				switch expression[p] {
				case ',':
					p++
				default:
					Vals = append(Vals, Parse())
				}
			}
			p++

			v := Vals[0]
			for _, b := range Vals[1:] {
				if andOr {
					v = v && b
				} else {
					v = v || b
				}
			}
			return v
		}
	}

	return Parse()
}

// 1140m Stone Games II
func stoneGameII(piles []int) int {
	// A_score + B_score = Total_score
	// A_score - B_score = MaxDelta
	// A_score -> (Total_score + MaxDelta)/2

	Mem := map[[2]int]int{}

	var MaxDelta func(start, M int) int
	MaxDelta = func(start, M int) int {
		if start == len(piles) {
			return 0
		}

		if v, ok := Mem[[2]int{start, M}]; ok {
			return v
		}

		xd := math.MinInt
		p := 0
		for i := 1; i <= 2*M && start+i-1 < len(piles); i++ {
			p += piles[start+i-1]
			xd = max(p-MaxDelta(start+i, max(M, i)), xd)
		}

		Mem[[2]int{start, M}] = xd
		return xd
	}

	xdelta := MaxDelta(0, 1)
	p := 0
	for _, v := range piles {
		p += v
	}

	log.Print(p, xdelta, Mem)
	return (p + xdelta) / 2
}

// 1545m Find Kth Bit in Nth Binary String
func findKthBit(n, k int) byte {
	if n == 1 {
		return '0'
	}

	lStr := 1<<n - 1
	if k-1 == lStr/2 {
		return '1'
	} else if k-1 < lStr/2 {
		return findKthBit(n-1, k)
	}

	if findKthBit(n-1, lStr-(k-1)) == '0' {
		return '1'
	}
	return '0'
}

// 1593m Split a String Into the Max Number of Unique Substrings
func maxUniqueSplit(s string) int {
	Subs := map[string]struct{}{}

	var W func(int) int
	W = func(start int) int {
		if start == len(s) {
			log.Printf(" -> %d %v", len(Subs), Subs)
			return 0
		}

		xlen := 0
		for end := start + 1; end <= len(s); end++ {
			if _, ok := Subs[s[start:end]]; !ok {
				Subs[s[start:end]] = struct{}{}
				xlen = max(W(end)+1, xlen)
				delete(Subs, s[start:end])
			}
		}
		return xlen
	}

	return W(0)
}

// 1718m Construct the Lexicographically Largest Valid Sequence
func constructDistancedSequence(n int) []int {
	R := make([]int, 2*n-1)
	M := make([]bool, n+1)

	var Search func(int) bool
	Search = func(start int) bool {
		if start == len(R) {
			return true
		}

		if R[start] != 0 {
			return Search(start + 1)
		}

		for v := n; v >= 1; v-- {
			if M[v] {
				continue
			}

			M[v] = true
			R[start] = v

			if v == 1 && Search(start+1) {
				return true
			}

			if v > 1 && start+v < len(R) && R[start+v] == 0 {
				R[start+v] = v
				if Search(start + 1) {
					return true
				}

				R[start+v] = 0 // backtrack
			}

			M[v] = false // backtrack...
			R[start] = 0 // backtrack
		}

		return false
	}

	Search(0)

	return R
}

// 1980m Find Unique Binary String
func findDifferentBinaryString(nums []string) string {
	numSize := len(nums[0])
	M := map[string]struct{}{}
	for _, s := range nums {
		M[s] = struct{}{}
	}

	var Search func([]byte) string
	Search = func(bfr []byte) string {
		if len(bfr) == numSize {
			if _, ok := M[string(bfr)]; !ok {
				return string(bfr)
			}
			return ""
		}

		bfr = append(bfr, '1')
		uniqueStr := Search(bfr)
		if uniqueStr != "" {
			return uniqueStr
		}

		bfr = bfr[:len(bfr)-1]
		bfr = append(bfr, '0')
		return Search(bfr)
	}

	return Search([]byte{})
}

// 2044m Count Number of Maximum Bitwise-OR Subsets
func countMaxOrSubsets(nums []int) int {
	xVal := 0
	for _, n := range nums {
		xVal |= n
	}

	var W func(start, v int) int
	W = func(start, v int) int {
		if start == len(nums) {
			if v == xVal {
				return 1
			}
			return 0
		}

		return W(start+1, v) + W(start+1, v|nums[start])
	}

	return W(0, 0)
}

// 2698m Find the Punishment Number of an Integer
func punishmentNumber(n int) int {
	var D []int

	var CanPartition func(start, cSum, targetSum int) bool
	CanPartition = func(start, cSum, targetSum int) bool {
		if start == len(D) {
			return cSum == targetSum
		}

		partitionSum := 0
		for p := start; p < len(D); p++ {
			partitionSum *= 10
			partitionSum += D[p]

			if CanPartition(p+1, cSum+partitionSum, targetSum) {
				return true
			}
		}

		return false
	}

	pSum := 0
	for x := 1; x <= n; x++ {
		D = []int{}
		for sqr := x * x; sqr > 0; sqr /= 10 {
			D = append(D, sqr%10)
		}
		slices.Reverse(D)

		if CanPartition(0, 0, x) {
			pSum += x * x
		}
	}
	return pSum
}

// 3302m Find the Lexicographically Smallest Valid Sequence
func validSequence(word1, word2 string) []int {
	R := []int{}

	Mem := map[[3]int]struct{}{}

	var v []int
	var Search func(start, pstart, flag int) bool
	Search = func(start, pstart, flag int) bool {
		if pstart == len(word2) {
			R = append(R, v...)
			return true
		}
		if start == len(word1) {
			return false
		}

		if _, ok := Mem[[3]int{start, pstart, flag}]; ok {
			return false
		}

		if word1[start] == word2[pstart] {
			v = append(v, start)
			if Search(start+1, pstart+1, flag) {
				return true
			}
			v = v[:len(v)-1]
		} else {
			if flag == 0 {
				v = append(v, start)
				if Search(start+1, pstart+1, 1) {
					return true
				}
				v = v[:len(v)-1]
			}

			for start < len(word1) && word1[start] != word2[pstart] {
				start++
			}
			if start < len(word1) {
				v = append(v, start)
				if Search(start+1, pstart+1, flag) {
					return true
				}
				v = v[:len(v)-1]
			}
		}

		Mem[[3]int{start, pstart, flag}] = struct{}{}
		return false
	}

	v = []int{}
	Search(0, 0, 0)

	return R
}

// 3309m Maximum Possible Number by Binary Concatenation
func maxGoodNumber(nums []int) int {
	xVal := 0

	r := []int{}

	var W func(int)
	W = func(start int) {
		if start == len(nums) {
			s := "0b"
			for _, n := range r {
				s += fmt.Sprintf("%b", n)
			}
			v, _ := strconv.ParseInt(s, 0, 0)
			xVal = max(int(v), xVal)

			log.Print(" -> ", r)
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			r = append(r, nums[start])
			W(start + 1)
			r = r[:len(r)-1]
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	W(0)

	return xVal
}
