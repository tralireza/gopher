package gopher

import (
	"log"
	"testing"
	"time"
)

// 10h Regular Expression Matching
func Test10(t *testing.T) {
	// 1 <= T|P Length <= 20

	DP := func(s, p string) bool {
		M := map[[2]int]bool{}

		var Match func(i, j int) bool
		Match = func(i, j int) bool {
			if j >= len(p) {
				return i >= len(s)
			}

			if found, ok := M[[2]int{i, j}]; ok {
				return found
			}

			fmatch := i < len(s) && (p[j] == s[i] || p[j] == '.')

			found := false
			if j+2 <= len(p) && p[j+1] == '*' {
				found = Match(i, j+2) || fmatch && Match(i+1, j)
			} else {
				found = fmatch && Match(i+1, j+1)
			}

			M[[2]int{i, j}] = found
			return found
		}

		return Match(0, 0)
	}

	for _, fn := range []func(string, string) bool{isMatch, DP} {
		ts := time.Now()
		log.Print("false ?= ", fn("aa", "a"))
		log.Print("true ?= ", fn("aa", "a*"))
		log.Print("true ?= ", fn("ab", ".*"))
		log.Print("true ?= ", fn("aab", "c*a*b"))
		log.Print(" -> ", time.Since(ts))
		log.Print("--")
	}
}

// 44h Wildcard Matching
func Test44(t *testing.T) {
	// 0 <= T|P Length <= 2000

	for _, c := range []struct {
		r bool
		P [2]string
	}{
		{false, [2]string{"aa", "a"}},
		{true, [2]string{"aa", "*"}},
		{false, [2]string{"cb", "?a"}},
	} {
		t.Run("", func(t *testing.T) {
			if c.r != isWildcardMatch(c.P[0], c.P[1]) {
				t.Fail()
			}
		})
	}
}

// 63m Unique Paths II
func Test63(t *testing.T) {
	log.Print("2 ?= ", uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	log.Print("1 ?= ", uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))

	log.Print("0 ?= ", uniquePathsWithObstacles([][]int{{1}}))
	log.Print("0 ?= ", uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}}))
}

// 72m Edit Distance
func Test72(t *testing.T) {
	log.Print("3 ?= ", minDistance("horse", "ros"))
	log.Print("5 ?= ", minDistance("intention", "execution"))
}

// 87h Scramble String
func Test87(t *testing.T) {
	Iterative := func(s1, s2 string) bool {
		N := len(s1) // || len(s2)!

		D := make([][][]bool, N+1)
		for x := range D {
			D[x] = make([][]bool, N)
			for y := range D[x] {
				D[x][y] = make([]bool, N)
			}
		}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				D[1][i][j] = s1[i] == s2[j]
			}
		}

		for l := 2; l <= N; l++ {
			for i := 0; i <= N-l; i++ {
				for j := 0; j <= N-l; j++ {
					for x := 1; x < l; x++ { // Split Index
						lSplit := D[x][i]
						rSplit := D[l-x][i+x]

						D[l][i][j] = lSplit[j] && rSplit[j+x] || lSplit[j+l-x] && rSplit[j]
					}
				}
			}
		}

		return D[N][0][0]
	}

	for _, fn := range []func(string, string) bool{isScramble, Iterative} {
		for _, c := range []struct {
			r      bool
			s1, s2 string
		}{
			{true, "great", "rgeat"},
			{false, "abcde", "caebd"},
			{true, "a", "a"},
		} {
			t.Run("", func(t *testing.T) {
				if c.r != fn(c.s1, c.s2) {
					t.Fail()
				}
			})
		}
	}
}

// 91m Decode Ways
func Test91(t *testing.T) {
	log.Print("2 ?= ", numDecodings("12"))
	log.Print("3 ?= ", numDecodings("226"))
	log.Print("0 ?= ", numDecodings("06"))
}

// 96m Unique Binary Search Trees
func Test96(t *testing.T) {
	log.Print("5 ?= ", numTrees(3))
	log.Print("1 ?= ", numTrees(1))
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

// 139m Word Break
func Test139(t *testing.T) {
	log.Print("true ?= ", wordBreak("applepenapple", []string{"apple", "pen"}))
	log.Print("false ?= ", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// 198m House Robber
func Test198(t *testing.T) {
	log.Print("4 ?= ", rob([]int{1, 2, 3, 1}))
	log.Print("12 ?= ", rob([]int{2, 7, 9, 3, 1}))
}

// 221m Maximal Square
func Test221(t *testing.T) {
	log.Print("4 ?= ", maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	log.Print("1 ?= ", maximalSquare([][]byte{{'0', '1'}, {'1', '0'}}))
	log.Print("0 ?= ", maximalSquare([][]byte{{'0'}}))
}

// 241m Different Ways to Add Parentheses
func Test241(t *testing.T) {
	Tabulation := func(expr string) []int {
		D := make([][][]int, len(expr))
		for i := range D {
			D[i] = make([][]int, len(expr)+1)
		}

		for i := 0; i < len(expr); i++ {
			if expr[i] >= '0' && expr[i] <= '9' {
				v := int(expr[i] - '0')
				j := i + 1
				if j < len(expr) && expr[j] >= '0' && expr[j] <= '9' {
					v = 10*v + int(expr[j]-'0')
					j++
				}
				D[i][j] = []int{v}
			}
		}

		for lexpr := 3; lexpr <= len(expr); lexpr++ {
			for i := 0; i <= len(expr)-lexpr; i++ {
				for j := i + 1; j-i < lexpr; j++ {
					if expr[j] >= '0' && expr[j] <= '9' {
						continue
					}

					for _, l := range D[i][j] {
						for _, r := range D[j+1][i+lexpr] {
							var v int
							switch expr[j] {
							case '*':
								v = l * r
							case '+':
								v = l + r
							case '-':
								v = l - r
							}
							D[i][i+lexpr] = append(D[i][i+lexpr], v)
						}
					}

				}
			}
		}

		return D[0][len(expr)]
	}

	for _, fn := range []func(string) []int{diffWaysToCompute, Tabulation} {
		log.Print("[0 2] ?= ", fn("2-1-1"))
		log.Print("[-34 -10 -14 -10 10] ?= ", fn("2*3-4*5"))
		log.Print("--")
	}
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

// 300m Longest Increasing Subsequence
func Test300(t *testing.T) {
	log.Print("4 ?= ", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	log.Print("4 ?= ", lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	log.Print("1 ?= ", lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

// 494m Target Sum
func Test494(t *testing.T) {
	// -1000 <= Target <= 1000, 1 <= Nums.Lengh <= 20

	Iterative := func(nums []int, target int) int {
		tSum := 0
		for _, n := range nums {
			tSum += n
		}

		if target > tSum {
			return 0
		}

		D := make([][]int, len(nums))
		for r := range D {
			D[r] = make([]int, 2*tSum+1)
		}

		D[0][tSum+nums[0]] += 1
		D[0][tSum-nums[0]] += 1

		for i := 1; i < len(nums); i++ {
			for t := -tSum; t <= tSum; t++ {
				if D[i-1][t+tSum] > 0 {
					D[i][t+tSum+nums[i]] += D[i-1][t+tSum]
					D[i][t+tSum-nums[i]] += D[i-1][t+tSum]
				}
			}
		}

		log.Print(D)

		return D[len(nums)-1][target+tSum]
	}

	for _, f := range []func([]int, int) int{findTargetSumWays, Iterative} {
		log.Print("5 ?= ", f([]int{1, 1, 1, 1, 1}, 3))
		log.Print("1 ?= ", f([]int{1}, 1))
		log.Print("256 ?= ", f([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
		log.Print("--")
	}
}

// 646m Maximum Length of Pair Chain
func Test646(t *testing.T) {
	log.Print("2 ?= ", findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	log.Print("3 ?= ", findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}

// 664h Strange Printer
func Test664(t *testing.T) {
	log.Print("2 ?= ", strangePrinter("aaabbb"))
	log.Print("2 ?= ", strangePrinter("aba"))
}

// 673m Number of Longest Increasing Subsequence
func Test673(t *testing.T) {
	log.Print("2 ?= ", findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	log.Print("5 ?= ", findNumberOfLIS([]int{2, 2, 2, 2, 2}))
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

// 1277m Count Square Submatrices with All Ones
func Test1277(t *testing.T) {
	Recursion := func(matrix [][]int) int {
		Rows, Cols := len(matrix), len(matrix[0])

		D := make([][]int, Rows)
		for r := range D {
			D[r] = make([]int, Cols)
			for c := range D[r] {
				D[r][c] = -1
			}
		}

		var W func(r, c int) int
		W = func(r, c int) int {
			if r >= Rows || c >= Cols {
				return 0
			}

			if matrix[r][c] == 0 {
				return 0
			}

			if D[r][c] != -1 {
				return D[r][c]
			}

			D[r][c] = 1 + min(W(r+1, c), W(r, c+1), W(r+1, c+1))
			return D[r][c]
		}

		count := 0
		for r := range Rows {
			for c := range Cols {
				count += W(r, c)
			}
		}
		return count
	}

	for _, fn := range []func([][]int) int{countSquares, Recursion} {
		log.Print("15 ?= ", fn([][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}))
		log.Print("7 ?= ", fn([][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}))
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

// 2707m Extra Characters in a String
func Test2707(t *testing.T) {
	Tabulation := func(s string, dictionary []string) int {
		Mem := map[string]struct{}{}
		for _, w := range dictionary {
			Mem[w] = struct{}{}
		}

		D := make([]int, len(s)+1)

		for start := len(s) - 1; start >= 0; start-- {
			D[start] = 1 + D[start+1]
			for end := start + 1; end <= len(s); end++ {
				if _, ok := Mem[s[start:end]]; ok {
					D[start] = min(D[start], D[end])
				}
			}
		}

		return D[0]
	}

	for _, fn := range []func(string, []string) int{minExtraChar, Tabulation} {
		log.Print("1 ?= ", fn("leetscode", []string{"leet", "code"}))
		log.Print("3 ?= ", fn("sayhelloworld", []string{"hello", "world"}))
		log.Print(" ?= ", fn("jqnrwkslbhhkkvveotpfaidoftmgcojcpzcvlctsqyvvobmlzo", []string{"nrwks", "t", "mcgjko", "xm", "vac", "ypqdr", "zwlghw", "gz", "xbsmr", "hhkkv", "qviu", "yvvobml", "cfk", "fxu", "pm", "nwobfce", "eu", "y", "krzbg", "xoktzxa", "doftmgc", "qpcpd", "oj", "bl", "kylslpr", "cpzcvlc", "ogscaz", "l", "nztlq", "ai", "o", "wdhlanl", "ot", "hqe"}))
		log.Print("--")
	}
}
