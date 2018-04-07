package gopher

import (
	"log"
	"testing"
)

// 63m Unique Paths II
func Test63(t *testing.T) {
	log.Print("2 ?= ", uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	log.Print("1 ?= ", uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))

	log.Print("0 ?= ", uniquePathsWithObstacles([][]int{{1}}))
	log.Print("0 ?= ", uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}}))
}

// 97m Interleaving String
func Test97(t *testing.T) {
	Tabulation := func(s1, s2, s3 string) bool {
		if len(s1)+len(s2) != len(s3) {
			return false
		}

		D := make([][]bool, len(s1)+1)
		for r := range D {
			D[r] = make([]bool, len(s2)+1)
		}

		D[0][0] = true
		for r := 1; r <= len(s1); r++ {
			D[r][0] = D[r-1][0] && s1[r-1] == s3[r-1]
		}
		for c := 1; c <= len(s2); c++ {
			D[0][c] = D[0][c-1] && s2[c-1] == s3[c-1]
		}

		for r := 1; r <= len(s1); r++ {
			for c := 1; c <= len(s2); c++ {
				D[r][c] = D[r-1][c] && s1[r-1] == s3[r+c-1] || D[r][c-1] && s2[c-1] == s3[r+c-1]
			}
		}

		return D[len(s1)][len(s2)]
	}

	for _, fn := range []func(string, string, string) bool{isInterleave, Tabulation} {
		log.Print("true ?= ", fn("aabcc", "dbbca", "aadbbcbcac"))
		log.Print("false ?= ", fn("aabcc", "dbbca", "aadbbbaccc"))
		log.Print("true ?= ", fn("", "", ""))
		log.Print("--")
	}
}

// 120m Triangle
func Test120(t *testing.T) {
	log.Print("11 ?= ", minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	log.Print("-10 ?= ", minimumTotal([][]int{{-10}}))
}

// 122m Best Time to Buy and Sell Stock II
func Test122(t *testing.T) {
	log.Print("7 ?= ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Print("4 ?= ", maxProfit([]int{1, 2, 3, 4, 5}))
	log.Print("0 ?= ", maxProfit([]int{7, 6, 4, 3, 1}))
}

// 221m Maximal Square
func Test221(t *testing.T) {
	log.Print("4 ?= ", maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	log.Print("1 ?= ", maximalSquare([][]byte{{'0', '1'}, {'1', '0'}}))
	log.Print("0 ?= ", maximalSquare([][]byte{{'0'}}))
}

// 264m Ugly Numbers II
func Test264(t *testing.T) {
	// Ugly: prime factors [only]: 2, 3, 5
	WithDP := func(n int) int {
		D := make([]int, n)
		D[0] = 1

		i2, i3, i5 := 0, 0, 0
		for i := range n - 1 {
			D[i+1] = min(D[i2]*2, D[i3]*3, D[i5]*5)

			if D[i+1] == D[i2]*2 {
				i2++
			}
			if D[i+1] == D[i3]*3 {
				i3++
			}
			if D[i+1] == D[i5]*5 {
				i5++
			}
		}

		return D[n-1]
	}

	for _, f := range []func(int) int{nthUglyNumber, WithDP} {
		log.Print("12 ?= ", f(10))
		log.Print("1 ?= ", f(1))
		log.Print("--")
	}
}

// 664h Strange Printer
func Test664(t *testing.T) {
	log.Print("2 ?= ", strangePrinter("aaabbb"))
	log.Print("2 ?= ", strangePrinter("aba"))
}

// 1014m Best Sightseeing Pair
func Test1014(t *testing.T) {
	log.Print("11 ?= ", maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	log.Print("2 ?= ", maxScoreSightseeingPair([]int{1, 2}))
}

// 1105m Filling Bookcase Shelves
func Test1105(t *testing.T) {
	//Book[i]: Width, Height

	Tabulation := func(books [][]int, shelfWidth int) int {
		DP := make([]int, len(books)+1)
		DP[0], DP[1] = 0, books[0][1]

		for i := 2; i <= len(books); i++ {
			book := books[i-1]
			w, h := book[0], book[1]

			DP[i] = DP[i-1] + h // this book "Goes" to next row

			for j := i - 1; j > 0 && w+books[j-1][0] <= shelfWidth; j-- {
				w += books[j-1][0]
				h = max(books[j-1][1], h)
				DP[i] = min(h+DP[j-1], DP[i])
			}
		}

		return DP[len(books)]
	}

	for _, f := range []func([][]int, int) int{minHeightShelves, Tabulation} {
		log.Print("6 ?= ", f([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
		log.Print("4 ?= ", f([][]int{{1, 3}, {2, 4}, {3, 2}}, 6))
		log.Print("--")
	}
}

// 1155m Number of Dice Rolls With Target Sum
func Test1155(t *testing.T) {
	MemOptimized := func(n, k, target int) int {
		const M = 1000_000_007

		cur := make([]int, max(k, target)+1)
		for r := range k {
			cur[r+1] = 1
		}

		for range n - 1 {
			prv := cur
			cur = make([]int, len(prv))
			for r := 1; r <= k; r++ {
				for s := 1; s+r <= target; s++ {
					cur[s+r] += prv[s]
					cur[s+r] %= M
				}
			}
		}

		return cur[target]
	}

	for _, fn := range []func(int, int, int) int{numRollsToTarget, MemOptimized} {
		log.Print("1 ?= ", fn(1, 6, 3))
		log.Print("6 ?= ", fn(2, 6, 7))
		log.Print("222616187 ?= ", fn(30, 30, 500))
		log.Print("--")
	}
}

// 1395m Count Number of Teams
func Test1395(t *testing.T) {
	// 1 <= Rating[i] <= 10^5

	Recursion := func(rating []int) int {
		Mem := map[[2]int]int{}

		var Count func(i, size int) int
		Count = func(i, size int) int {
			if size == 3 {
				return 1
			}

			if v, ok := Mem[[2]int{i, size}]; ok {
				return v
			}

			log.Print(" -> ", i, size)

			v := 0
			for j := i + 1; j < len(rating); j++ {
				if rating[j] > rating[i] {
					v += Count(j, size+1)
				}
			}

			Mem[[2]int{i, size}] = v
			return v
		}

		teams := 0
		for start := 0; start < len(rating)-2; start++ {
			teams += Count(start, 1) // increasing
		}
		log.Print(Mem)

		clear(Mem)
		for i := range rating {
			rating[i] *= -1
		}
		for start := 0; start < len(rating)-2; start++ {
			teams += Count(start, 1) // decreasing
		}
		log.Print(Mem)

		return teams
	}

	Tabulation := func(rating []int) int {
		Incs, Decs := make([][4]int, len(rating)), make([][4]int, len(rating))

		// team of size 1 with (last) member at index i -> 1 (team)
		for i := range len(rating) {
			Incs[i][1], Decs[i][1] = 1, 1
		}

		for size := 2; size <= 3; size++ {
			for i := 0; i < len(rating); i++ {
				for j := i + 1; j < len(rating); j++ {
					if rating[j] > rating[i] {
						Incs[j][size] += Incs[i][size-1]
					}

					if rating[j] < rating[i] {
						Decs[j][size] += Decs[i][size-1]
					}
				}
			}
		}

		log.Print(Incs)
		log.Print(Decs)

		teams := 0
		for i := range Incs {
			teams += Incs[i][3]
			teams += Decs[i][3]
		}
		return teams
	}

	for _, f := range []func([]int) int{numTeams, Recursion, Tabulation} {
		log.Print("3 ?= ", f([]int{2, 5, 3, 4, 1}))
		log.Print("0 ?= ", f([]int{2, 1, 3}))
		log.Print("4 ?= ", f([]int{1, 2, 3, 4}))
		log.Print("--")
	}
}

// 1653m Minimum Deletions to Make String Balanced
func Test1653(t *testing.T) {
	WithStack := func(s string) int {
		Q := []byte{}

		dels := 0
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case 'a':
				if len(Q) > 0 && Q[0] == 'b' {
					Q = Q[:len(Q)-1]
					dels++
				}
			case 'b':
				Q = append(Q, 'b')
			}
		}
		return dels
	}

	for _, f := range []func(string) int{minimumDeletions, WithStack} {
		log.Print("2 ?= ", f("aababbab"))
		log.Print("2 ?= ", f("bbaaaaabb"))
		log.Print("--")
	}
}

// 1937m Maximum Number of Points with Cost
func Test1937(t *testing.T) {
	log.Print("9 ?= ", maxPoints([][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}))
	log.Print("11 ?= ", maxPoints([][]int{{1, 5}, {2, 3}, {4, 2}}))
}

// 2016 Maximum Difference Between Increasing Elements
func Test2016(t *testing.T) {
	log.Print("4 ?= ", maximumDifference([]int{7, 1, 5, 4}))
	log.Print("-1 ?= ", maximumDifference([]int{9, 4, 3, 2}))
	log.Print("9 ?= ", maximumDifference([]int{1, 5, 2, 10}))
}
