package gopher

import (
	"fmt"
	"log"
	"testing"
)

// 17m Letter Combinations of a Phone Number
func Test17(t *testing.T) {
	log.Printf(`["ae" "af" "be" "bf" "ce" "cf"] ?= %q`, letterCombinations("23"))
	log.Printf(`[] ?= %q`, letterCombinations(""))
	log.Printf(`["a" "b" "c"] ?= %q`, letterCombinations("2"))
}

// 40m Combination Sum II
func Test40(t *testing.T) {
	log.Print(" ?= ", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	log.Print(" ?= ", combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

// 46m Permutations
func Test46(t *testing.T) {
	log.Print(" ?= ", permute([]int{1, 2, 3}))
	log.Print(" ?= ", permute([]int{0, 1}))
	log.Print(" ?= ", permute([]int{1}))

	log.Print(" ?= ", permute([]int{0, 4, 5, 9}))
}

// 52h N-Queens II
func Test52(t *testing.T) {
	// 1 <= N <= 9

	log.Print("2 ?= ", totalNQueens(4, func(int, [][]byte) {}))
	log.Print("1 ?= ", totalNQueens(1, func(int, [][]byte) {}))

	log.Print(" ?= ", totalNQueens(8, func(i int, B [][]byte) {
		Row := func(r []byte) string {
			s := ""
			for _, b := range r {
				if b == '~' {
					b = '-'
				}
				s += "|" + string(b)
			}
			return s + "|"
		}
		if i == 89 || i == 92 {
			for r := range B {
				log.Printf(" -> %d :: %s", i, Row(B[r]))
			}
		}
	}))
}

// 77m Combinations
func Test77(t *testing.T) {
	log.Print(" ?= ", combine(4, 2))
	log.Print(" ?= ", combine(1, 1))

	log.Print(" ?= ", combine(7, 3))
}

// 224h Basic Calculator
func Test224(t *testing.T) {
	log.Print("2 ?= ", calculate("1 + 1"))
	log.Print("3 ?= ", calculate(" 2-1 + 2 "))
	log.Print("23 ?= ", calculate("(1+(4+5+2)-3)+(6+8)"))

	log.Print("0 ?= ", calculate("0"))
	log.Print("-2 ?= ", calculate("-2"))
	log.Print("1 ?= ", calculate("+1"))
	log.Print("-3 ?= ", calculate("1-(1+3)"))
	log.Print("2 ?= ", calculate("13-(1+3+((3+1)+4))+1"))
}

// 386m Lexicographical Numbers
func Test386(t *testing.T) {
	Recursive := func(n int) []int {
		R := []int{}

		var W func(int)
		W = func(v int) {
			R = append(R, v)
			for r := 0; r <= 9 && 10*v+r <= n; r++ {
				W(10*v + r)
			}
		}

		for r := 1; r <= 9 && r <= n; r++ {
			W(r)
		}

		return R
	}

	for _, fn := range []func(int) []int{Recursive, lexicalOrder} {
		log.Print("[1 10 11 12 13 2 3 4 5 6 7 8 9] ?= ", fn(13))
		log.Print(" ?= ", fn(23))
		log.Print(" ?= ", fn(137)[:57], "...")
		log.Print("--")
	}
}

// 427m Construct Quad Tree
func (n *Node427) String() string {
	tl, tr, bl, br := "/", "/", "/", "/"
	if n.TopLeft != nil {
		tl = "->"
	}
	if n.TopRight != nil {
		tr = "->"
	}
	if n.BottomLeft != nil {
		bl = "->"
	}
	if n.BottomRight != nil {
		br = "->"
	}
	return fmt.Sprintf("{ Val: %t  IsLeaf: %t   TL %s  TR %s  BL %s  BR %s }", n.Val, n.IsLeaf, tl, tr, bl, br)
}

func Test427(t *testing.T) {
	// n = 2^x, 0 <= x <= 6

	log.Printf("%+v", construct([][]int{{0, 1}, {1, 0}}))
	log.Print("--")

	qtree := construct([][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}})
	log.Print(qtree)
	log.Print("   TL -> ", qtree.TopLeft)
	log.Print("   TR -> ", qtree.TopRight)
	log.Print("      TR -> ", qtree.TopRight.TopRight)
	log.Print("      BL -> ", qtree.TopRight.BottomLeft)
	log.Print("   BR -> ", qtree.BottomRight)
}

// 650m 2 Keys Keyboard
func Test650(t *testing.T) {
	// 1 <= n <= 1000

	BottomUp := func(n int) int {
		// f(i) = f(j) + i/j | i mod j == 0

		D := make([]int, 1001)
		for i := 2; i <= n; i++ {
			D[i] = 1000 // init -> Max

			for j := 1; j <= n/2; j++ {
				if i%j == 0 {
					D[i] = min(D[i], D[j]+i/j)
				}
			}
		}
		return D[n]
	}

	for _, f := range []func(int) int{minSteps, BottomUp} {
		log.Print("3 ?= ", f(3))
		log.Print("0 ?= ", f(1))
		log.Print("8 ?= ", f(15))
		log.Print("21 ?= ", f(1000))
		log.Print("--")
	}
}

// 1140m Stone Games II
func Test1140(t *testing.T) {
	log.Print("10 ?= ", stoneGameII([]int{2, 7, 9, 4, 4}))
	log.Print("104 ?= ", stoneGameII([]int{1, 2, 3, 4, 5, 100}))
}

// 2044m Count Number of Maximum Bitwise-OR Subsets
func Test2044(t *testing.T) {
	// 1 <= N_i <= 10^5, N.length <= 16

	PowerSet := func(nums []int) int {
		orVal := 0
		for _, n := range nums {
			orVal |= n
		}

		count := 0
		for mask := range 1 << len(nums) {
			v := 0
			for i := range len(nums) {
				if mask&(1<<i) != 0 {
					v |= nums[i]
				}
			}

			if v == orVal {
				count++
			}
		}

		return count
	}

	for _, fn := range []func([]int) int{countMaxOrSubsets, PowerSet} {
		log.Print("2 ?= ", fn([]int{3, 1}))
		log.Print("7 ?= ", fn([]int{2, 2, 2}))
		log.Print("6 ?= ", fn([]int{3, 2, 1, 5}))
		log.Print("--")
	}
}

// 3302m Find the Lexicographically Smallest Valid Sequence
func Test3302(t *testing.T) {
	log.Print("[0 1 2] ?= ", validSequence("vbcca", "abc"))
	log.Print("[1 2 4] ?= ", validSequence("bacdc", "abc"))
	log.Print("[] ?= ", validSequence("aaaaaa", "aaabc"))
	log.Print("[0 1] ?= ", validSequence("abc", "ab"))
}

// 3309m Maximum Possible Number by Binary Concatenation
func Test3309(t *testing.T) {
	// 1 <= N_i <= 127, N.length = 3

	log.Print("30 ?= ", maxGoodNumber([]int{1, 2, 3}))
	log.Print("1296 ?= ", maxGoodNumber([]int{2, 8, 16}))
}
