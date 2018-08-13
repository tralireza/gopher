package gopher

import (
	"fmt"
	"log"
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
func findKthBit(n int, k int) byte {
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
