package gopher

import (
	"log"
	"reflect"
	"slices"
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

// 174h Dungeon Game
func Test174(t *testing.T) {
	log.Print("7 ?= ", calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	log.Print("1 ?= ", calculateMinimumHP([][]int{{0}}))
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

// 689h Maximum Sum of 3 Non-Overlapping Subarrays
func Test689(t *testing.T) {
	Tabulation := func(nums []int, k int) []int {
		kSums := []int{}

		rSum := 0
		for i := range nums {
			rSum += nums[i]
			if i+1 >= k {
				kSums = append(kSums, rSum)
				rSum -= nums[i-k+1]
			}
		}

		log.Print(" -> kSums :: ", kSums)

		D := make([][4]int, len(kSums)) // DP
		T := make([][4]int, len(kSums)) // Trace

		for i := 0; i < len(kSums); i++ {
			for r := 1; r <= 3; r++ {
				if i > 0 {
					T[i][r] = T[i-1][r]
					if i-k >= 0 {
						D[i][r] = max(kSums[i]+D[i-k][r-1], D[i-1][r])
						if D[i][r] > D[i-1][r] {
							T[i][r] = i
						}
					} else {
						D[i][r] = max(kSums[i], D[i-1][r])
						if kSums[i] > D[i-1][r] {
							T[i][r] = i
						}
					}
				} else {
					D[i][r] = kSums[i]
					T[i][r] = i
				}
			}
		}

		log.Print(" -> DP :: ", D)
		log.Print(" -> Trace :: ", T)

		R := []int{}
		i := len(kSums) - 1
		for r := 3; r > 0; r-- {
			R = append(R, T[i][r])
			i = T[i][r] - k
		}
		slices.Reverse(R)

		log.Print(" -> ", R)

		return R
	}

	for _, f := range []func([]int, int) []int{maxSumOfThreeSubarrays, Tabulation} {
		if !reflect.DeepEqual([]int{0, 3, 5}, f([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2)) {
			t.Error("!")
		}
		if !reflect.DeepEqual([]int{0, 2, 4}, f([]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2)) {
			t.Error("!")
		}
		if !reflect.DeepEqual([]int{1, 4, 7}, f([]int{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}, 3)) {
			t.Error("!")
		}
		log.Print("--")
	}
}

// 983m Minimum Cost For Tickets
func Test983(t *testing.T) {
	Iterative := func(days []int, costs []int) int {
		D := make([]int, days[len(days)-1]+1)

		TDays := map[int]bool{}
		for _, day := range days {
			TDays[day] = true
		}

		for day := 1; day <= days[len(days)-1]; day++ {
			if TDays[day] {
				D[day] = min(D[day-1]+costs[0], D[max(day-7, 0)]+costs[1], D[max(day-30, 0)]+costs[2])
			} else {
				D[day] = D[day-1]
			}
		}

		return D[days[len(days)-1]]
	}

	for _, f := range []func([]int, []int) int{mincostTickets, Iterative} {
		log.Print("11 ?= ", f([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
		log.Print("17 ?= ", f([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
		log.Print("11 ?= ", f([]int{1, 4, 6, 7, 8, 365}, []int{2, 7, 15}))

		log.Print("1780 ?= ", f([]int{2, 7, 8, 10, 12, 13, 14, 17, 25, 28, 29, 34, 35, 37, 43, 44, 45, 53, 54, 58, 60, 61, 62, 63, 64, 65, 66, 71, 74, 82, 86, 88, 95, 97, 98, 102, 105, 106, 115, 117, 119, 120, 125, 129, 135, 136, 142, 143, 152, 153, 155, 158, 165, 166, 168, 181, 187, 189, 191, 192, 193, 194, 196, 197, 198, 201, 208, 209, 211, 212, 215, 224, 226, 236, 242, 243, 244, 245, 248, 252, 260, 261, 263, 266, 269, 272, 273, 274, 280, 284, 286, 287, 292, 297, 300, 303, 304, 312, 317, 323, 326, 328, 329, 332, 333, 337, 341, 348, 349, 351, 352, 355, 361, 364}, []int{16, 82, 359}))
		log.Print("--")
	}
}

func Benchmark983(b *testing.B) {
	MinCostTickets1 := func(days []int, costs []int) int {
		Mem := map[int]int{}

		TDays := make([]bool, 365+1)
		for _, day := range days {
			TDays[day] = true
		}

		var Search func(int) int
		Search = func(day int) int {
			if day > 365 {
				return 0
			}

			if !TDays[day] {
				return Search(day + 1)
			}

			if v, ok := Mem[day]; ok {
				return v
			}

			mCost := min(costs[0]+Search(day+1), costs[1]+Search(day+7), costs[2]+Search(day+30))
			Mem[day] = mCost
			return mCost
		}
		return Search(days[0])
	}

	MinCostTickets2 := func(days []int, costs []int) int {
		Mem := map[int]int{}

		TDays := make([]int, 365+1)
		for _, day := range days {
			TDays[day] = day
		}
		nday := 365 + 1
		for day := 365; day >= 1; day-- {
			if TDays[day] == 0 {
				TDays[day] = nday
			}
			nday = TDays[day]
		}

		var Search func(int) int
		Search = func(day int) int {
			if day > 365 {
				return 0
			}

			day = TDays[day]
			if day > 365 {
				return 0
			}

			if v, ok := Mem[day]; ok {
				return v
			}

			mCost := min(costs[0]+Search(day+1), costs[1]+Search(day+7), costs[2]+Search(day+30))
			Mem[day] = mCost
			return mCost
		}
		return Search(days[0])
	}

	Noop := func(days []int, costs []int) int { return 0 }

	days := []int{2, 7, 8, 10, 12, 13, 14, 17, 25, 28, 29, 34, 35, 37, 43, 44, 45, 53, 54, 58, 60, 61, 62, 63, 64, 65, 66, 71, 74, 82, 86, 88, 95, 97, 98, 102, 105, 106, 115, 117, 119, 120, 125, 129, 135, 136, 142, 143, 152, 153, 155, 158, 165, 166, 168, 181, 187, 189, 191, 192, 193, 194, 196, 197, 198, 201, 208, 209, 211, 212, 215, 224, 226, 236, 242, 243, 244, 245, 248, 252, 260, 261, 263, 266, 269, 272, 273, 274, 280, 284, 286, 287, 292, 297, 300, 303, 304, 312, 317, 323, 326, 328, 329, 332, 333, 337, 341, 348, 349, 351, 352, 355, 361, 364}
	costs := []int{16, 82, 359}

	for _, fn := range []func([]int, []int) int{MinCostTickets1, MinCostTickets2, Noop} {
		b.Run("", func(b *testing.B) {
			for range b.N {
				fn(days, costs)
			}
		})
	}
}

// 1014m Best Sightseeing Pair
func Test1014(t *testing.T) {
	Kadane := func(values []int) int {
		xScore := 0

		lScore := values[0] + 0
		for i := 1; i < len(values); i++ {
			xScore = max(lScore+values[i]-i, xScore)
			lScore = max(values[i]+i, lScore)
		}

		return xScore
	}

	for _, f := range []func([]int) int{maxScoreSightseeingPair, Kadane} {
		log.Print("11 ?= ", f([]int{8, 1, 5, 2, 6}))
		log.Print("2 ?= ", f([]int{1, 2}))
		log.Print("--")
	}
}

// 1025 Divisor Game
func Test1025(t *testing.T) {
	log.Print("true ?= ", divisorGame(2))
	log.Print("false ?= ", divisorGame(3))
}

// 1092h Shortest Common Supersequence
func Test1092(t *testing.T) {
	// (!TLE)
	var Recursive func(str1, str2 string) string
	Recursive = func(str1, str2 string) string {
		if str1 == "" && str2 == "" {
			return ""
		}
		if str1 == "" {
			return str2
		}
		if str2 == "" {
			return str1
		}

		switch str1[0] == str2[0] {
		case true:
			return string(str1[0]) + Recursive(str1[1:], str2[1:])
		default:
			scs1 := string(str1[0]) + Recursive(str1[1:], str2)
			scs2 := string(str2[0]) + Recursive(str1, str2[1:])

			if len(scs1) <= len(scs2) {
				return scs1
			}
			return scs2
		}
	}

	// (!MLE)
	RecursiveMemo := func(str1, str2 string) string {
		M := map[[2]string]string{}

		var Search func(str1, str2 string) string
		Search = func(str1, str2 string) string {
			if str1 == "" && str2 == "" {
				return ""
			}
			if str1 == "" {
				return str2
			}
			if str2 == "" {
				return str1
			}

			if scs, ok := M[[2]string{str1, str2}]; ok {
				return scs
			}

			if str1[0] == str2[0] {
				M[[2]string{str1, str2}] = string(str1[0]) + Search(str1[1:], str2[1:])
				return M[[2]string{str1, str2}]
			}

			scs1 := string(str1[0]) + Search(str1[1:], str2)
			scs2 := string(str2[0]) + Search(str1, str2[1:])

			if len(scs1) <= len(scs2) {
				M[[2]string{str1, str2}] = scs1
			} else {
				M[[2]string{str1, str2}] = scs2
			}

			return M[[2]string{str1, str2}]
		}

		return Search(str1, str2)
	}

	for _, f := range []func(string, string) string{shortestCommonSupersequence, Recursive, RecursiveMemo} {
		tStart := time.Now()
		log.Print("cabac ?= ", f("abac", "cab"))
		log.Print("aaaaaaaa ?= ", f("aaaaaaaa", "aaaaaaaa"))

		log.Print(" ?= ", f("abcdefghijkl", "efghijklmnopqr"))
		log.Print("-- ", time.Since(tStart))
	}
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

// 1524m Number of Sub-arrays With Odd Sum
func Test1524(t *testing.T) {
	PSum := func(arr []int) int {
		const M = 1000_000_000 + 7

		count := 0
		evens, odds := 1, 0

		psum := 0
		for _, n := range arr {
			psum += n
			switch psum & 1 {
			case 1:
				count += evens
				odds++
			case 0:
				count += odds
				evens++
			}
		}

		return count % M
	}

	for _, f := range []func([]int) int{numOfSubarrays, PSum} {
		log.Print("4 ?= ", f([]int{1, 3, 5}))
		log.Print("0 ?= ", f([]int{2, 4, 6}))
		log.Print("16 ?= ", f([]int{1, 2, 3, 4, 5, 6, 7}))
		log.Print("--")
	}
}

// 1639h Number of Ways to Form a Target String Given a Dictionary
func Test1639(t *testing.T) {
	Tabulation := func(words []string, target string) int {
		F := make([][26]int, len(words[0])) // i-th Letter Frequency in Words[]
		for p := 0; p < len(words[0]); p++ {
			for w := range words {
				F[p][words[w][p]-'a']++
			}
		}

		D := make([][]int, len(words[0])+1)
		for w := range D {
			D[w] = make([]int, len(target)+1)
		}

		for w := 0; w <= len(words[0]); w++ {
			D[w][0] = 1 // only one way to form an empty "target" string
		}

		const MOD = 1e9 + 7
		for w := 1; w <= len(words[0]); w++ {
			for t := 1; t <= len(target); t++ {
				D[w][t] = D[w-1][t]
				D[w][t] += (F[w-1][target[t-1]-'a'] * D[w-1][t-1]) % MOD
				D[w][t] %= MOD
			}
		}

		log.Print(D)

		return D[len(words[0])][len(target)]
	}

	for _, f := range []func([]string, string) int{numWays, Tabulation} {
		log.Print("6 ?= ", f([]string{"acca", "bbbb", "caca"}, "aba"))
		log.Print("4 ?= ", f([]string{"abba", "baab"}, "bab"))
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

// 1749m Maximum Absolute Sum of Any Subarray
func Test1749(t *testing.T) {
	Optimized := func(nums []int) int {
		pSum := 0
		xSum, nSum := 0, 0

		for _, n := range nums {
			pSum += n
			xSum, nSum = max(pSum, xSum), min(pSum, nSum)
		}

		return xSum - nSum
	}

	for _, f := range []func([]int) int{maxAbsoluteSum, Optimized} {
		log.Print("5 ?= ", f([]int{1, -3, 2, 3, -4}))
		log.Print("8 ?= ", f([]int{2, -5, 1, -4, 3, -2}))
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

// 2466m Count Ways to Build Good Strings
func Test2466(t *testing.T) {
	Recursive := func(low, high, zero, one int) int {
		const MOD = 1e9 + 7
		Mem := map[int]int{}

		var Search func(int) int
		Search = func(l int) int {
			if l == 0 {
				return 1
			}

			if v, ok := Mem[l]; ok {
				return v
			}

			tWays := 0
			if l >= zero {
				tWays += Search(l - zero)
			}
			if l >= one {
				tWays += Search(l - one)
			}
			tWays %= MOD

			Mem[l] = tWays
			return tWays
		}

		tWays := 0
		for l := low; l <= high; l++ {
			tWays += Search(l)
			tWays %= MOD
		}
		return tWays
	}

	for _, f := range []func(int, int, int, int) int{countGoodStrings, Recursive} {
		log.Print("8 ?= ", f(3, 3, 1, 1))
		log.Print("5 ?= ", f(2, 3, 1, 2))
		log.Print("--")
	}
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

// 2836h Maximize Value of Function in a Ball Passing Game
func Test2836(t *testing.T) {
	log.Print("6 ?= ", getMaxFunctionValue([]int{2, 0, 1}, 4))
	log.Print("10 ?= ", getMaxFunctionValue([]int{1, 1, 1, 2, 3}, 3))
}
