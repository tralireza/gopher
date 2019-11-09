package gopher

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
)

// 10h Regular Expression Matching
func isMatch(s, p string) bool {
	if p == "" {
		return s == ""
	}

	fmatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || fmatch && isMatch(s[1:], p)
	}
	return fmatch && isMatch(s[1:], p[1:])
}

// 44h Wildcard Matching
func isWildcardMatch(s, p string) bool {
	M := map[[2]int]bool{}

	var Match func(i, j int) bool
	Match = func(i, j int) bool {
		if j >= len(p) {
			return i >= len(s)
		}

		if found, ok := M[[2]int{i, j}]; ok {
			return found
		}

		found := false
		if p[j] == '*' {
			found = Match(i, j+1) || i < len(s) && Match(i+1, j)
		} else if i < len(s) {
			found = (s[i] == p[j] || p[j] == '?') && Match(i+1, j+1)
		}

		M[[2]int{i, j}] = found
		return found
	}

	return Match(0, 0)
}

// 63m Unique Paths II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	Rows, Cols := len(obstacleGrid), len(obstacleGrid[0])

	D := make([][]int, Rows)
	for r := range D {
		D[r] = make([]int, Cols)
	}

	if obstacleGrid[0][0] != 1 {
		D[0][0] = 1
	}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if obstacleGrid[r][c] != 1 {
				if r > 0 {
					D[r][c] += D[r-1][c]
				}
				if c > 0 {
					D[r][c] += D[r][c-1]
				}
			}
		}
	}

	return D[Rows-1][Cols-1]
}

// 72m Edit Distance
func minDistance(word1 string, word2 string) int {
	Dist := make([][]int, len(word1)+1)
	for r := range Dist {
		Dist[r] = make([]int, len(word2)+1)
	}

	for r := 1; r <= len(word1); r++ {
		Dist[r][0] = r
	}
	for c := 1; c <= len(word2); c++ {
		Dist[0][c] = c
	}

	for r := 1; r <= len(word1); r++ {
		for c := 1; c <= len(word2); c++ {
			Dist[r][c] = Dist[r-1][c-1]

			if word1[r-1] != word2[c-1] {
				Dist[r][c] = 1 + min(Dist[r-1][c], Dist[r][c-1], Dist[r][c])
			}
		}
	}

	return Dist[len(word1)][len(word2)]
}

// 87h Scramble String
func isScramble(s1, s2 string) bool {
	M := map[[2]string]bool{}
	defer log.Print(" -> Map :: ", M)

	var Check func(s1, s2 string) bool
	Check = func(s1, s2 string) bool {
		if len(s1) == 1 && s1 == s2 {
			return true
		}
		if v, ok := M[[2]string{s1, s2}]; ok {
			return v
		}

		for i := 1; i < len(s1); i++ {
			if Check(s1[i:], s2[i:]) && Check(s1[:i], s2[:i]) || Check(s1[:i], s2[len(s2)-i:]) && Check(s1[i:], s2[:len(s2)-i]) {
				M[[2]string{s1, s2}] = false
				return true
			}
		}

		M[[2]string{s1, s2}] = false
		return false
	}

	return Check(s1, s2)
}

// 91m Decode Ways
func numDecodings(s string) int {
	D := make([]int, len(s)+1)

	if s[0] == '0' {
		return 0
	}

	D[0], D[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		v1 := s[i-1] - '0'
		v2 := (s[i-2]-'0')*10 + v1

		if v1 >= 1 && v1 <= 9 {
			D[i] += D[i-1]
		}
		if v2 >= 10 && v2 <= 26 {
			D[i] += D[i-2]
		}
	}

	return D[len(s)]
}

// 96m Unique Binary Search Trees
func numTrees(n int) int {
	D := []int{1, 1}

	for x := 2; x <= n; x++ {
		v := 0
		for i := 1; i <= x; i++ {
			v += D[i-1] * D[x-i]
		}
		D = append(D, v)
	}

	return D[n]
}

// 97m Interleaving String
func isInterleave(s1, s2, s3 string) bool {
	Mem := map[[2]string]bool{}
	defer log.Print(Mem)

	var Check func(s1, s2, s3 string) bool
	Check = func(s1, s2, s3 string) bool {
		if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
			return true
		}

		if v, ok := Mem[[2]string{s1, s2}]; ok {
			return v
		}
		if v, ok := Mem[[2]string{s2, s1}]; ok {
			return v
		}

		log.Printf("%q %q ? %q", s1, s2, s3)

		S := []string{s2, s1}
		for i, s := range []string{s1, s2} {
			t := S[i]
			for i := 1; i <= len(s); i++ {
				if s[:i] == s3[:i] {
					if Check(s[i:], t, s3[i:]) {
						Mem[[2]string{s[i:], t}] = true
						return true
					}
				}
			}
		}

		Mem[[2]string{s1, s2}] = false
		return false
	}

	return Check(s1, s2, s3)
}

// 120m Triangle
func minimumTotal(triangle [][]int) int {
	t := make([][]int, len(triangle))
	for r := range t {
		t[r] = make([]int, r+1)
	}

	t[0][0] = triangle[0][0]
	for r := 1; r < len(t); r++ {
		t[r][0] += t[r-1][0] + triangle[r][0]
		t[r][r] += t[r-1][r-1] + triangle[r][r-1]

		for c := 1; c < r; c++ {
			t[r][c] = triangle[r][c] + min(t[r-1][c-1], t[r-1][c])
		}
	}

	log.Print(t)

	return slices.Min(t[len(t)-1])
}

// 122m Best Time to Buy and Sell Stock II
func maxProfit(prices []int) int {
	profit := 0
	for i, price := range prices[:len(prices)-1] {
		if prices[i+1] > price {
			profit += prices[i+1] - price
		}
	}
	return profit
}

// 139m Word Break
func wordBreak(s string, wordDict []string) bool {
	M := map[string]bool{}
	for _, w := range wordDict {
		M[w] = true
	}

	D := make([]bool, len(s)+1)

	D[0] = true
	for l := range len(s) {
		for w := range M {
			lw := len(w)
			if l+1-lw >= 0 && D[l+1-lw] && w == s[l+1-lw:l+1] {
				D[l+1] = true

				log.Printf("%s(%d) ~ %s|%s(%d) -> true", w, lw, s[:l+1-lw], s[l+1-lw:l+1], l+1)
			}
		}
	}

	return D[len(s)]
}

// 174h Dungeon Game
func calculateMinimumHP(dungeon [][]int) int {
	Rows, Cols := len(dungeon), len(dungeon[0])

	var Search func(r, c int) int
	Search = func(r, c int) int {
		if r >= Rows || c >= Cols {
			return math.MaxInt
		}

		if r == Rows-1 && c == Cols-1 {
			if dungeon[r][c] < 0 {
				return 1 - dungeon[r][c]
			}
			return 1
		}

		right, down := Search(r, c+1), Search(r+1, c)
		health := min(right, down) - dungeon[r][c]

		if health <= 0 {
			return 1
		}
		return health
	}

	return Search(0, 0)
}

// 198m House Robber
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	D := make([]int, len(nums))

	D[0], D[1] = nums[0], nums[1]
	for i := range len(D) - 2 {
		D[i+2] = max(D[i]+nums[i+2], D[i+1])
	}
	return D[len(D)-1]
}

// 221m Maximal Square
func maximalSquare(matrix [][]byte) int {
	Rows, Cols := len(matrix), len(matrix[0])
	D := make([][]int, Rows+1)
	for r := range D {
		D[r] = make([]int, Cols+1)
	}

	x := 0
	for r := range Rows {
		for c := range Cols {
			if matrix[r][c] == '1' {
				D[r+1][c+1] = 1 + min(D[r][c], D[r+1][c], D[r][c+1])
				x = max(D[r+1][c+1], x)
			}
		}
	}

	return x * x
}

// 233h Number of Digit One
func countDigitOne(n int) int {
	ones := 0
	for r := 1; r <= n; r *= 10 {
		dvr := r * 10
		ones += (n/dvr)*r + min(max(n%dvr-r+1, 0), r)
	}

	return ones
}

// 241m Different Ways to Add Parentheses
func diffWaysToCompute(expression string) []int {
	Mem := map[[2]int][]int{}
	defer log.Print(" -> ", Mem)

	var Calc func(start, end int) []int
	Calc = func(start, end int) []int {
		if start >= end {
			return []int{}
		}
		if end-start == 1 {
			return []int{int(expression[start] - '0')}
		}
		if end-start == 2 {
			return []int{10*int(expression[start]-'0') + int(expression[start+1]-'0')}
		}

		if mR, ok := Mem[[2]int{start, end}]; ok {
			log.Printf(" :: Mem[%d:%d] => %v", start, end, mR)
			return mR
		}

		R := []int{}
		for i := start; i < end; i++ {
			if expression[i] >= '0' && expression[i] <= '9' {
				continue
			}

			lR, rR := Calc(start, i), Calc(i+1, end)
			for _, l := range lR {
				for _, r := range rR {
					switch expression[i] {
					case '*':
						R = append(R, l*r)
					case '+':
						R = append(R, l+r)
					case '-':
						R = append(R, l-r)
					}
				}
			}
		}

		Mem[[2]int{start, end}] = R
		return R
	}

	return Calc(0, len(expression))
}

// 264m Ugly Numbers II
type PQ264 struct{ sort.IntSlice }

func (h *PQ264) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *PQ264) Pop() any {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func nthUglyNumber(n int) int {
	type PQ = PQ264

	h := PQ{}
	heap.Push(&h, 1)

	Mem := map[int]struct{}{}
	Mem[1] = struct{}{}

	var u int
	for range n {
		u = heap.Pop(&h).(int)
		delete(Mem, u)

		for _, f := range []int{2, 3, 5} {
			if _, ok := Mem[f*u]; !ok {
				heap.Push(&h, f*u)
				Mem[f*u] = struct{}{}
			}
		}
	}

	return u
}

// 300m Longest Increasing Subsequence
func lengthOfLIS(nums []int) int {
	D := make([]int, len(nums))
	for i := range D {
		D[i] = 1
	}

	for r := 1; r < len(nums); r++ {
		for l := 0; l < r; l++ {
			if nums[l] < nums[r] {
				D[r] = max(D[r], D[l]+1)
			}
		}
	}

	return slices.Max(D)
}

// 312h Burst Balloons
func maxCoins(nums []int) int {
	M := make([][]int, 300+1)
	for r := range M {
		M[r] = make([]int, 300+1)
	}

	var Search func(i, j int) int
	Search = func(i, j int) int {
		if i > j {
			return 0
		}

		if i == j {
			coins := nums[i]
			if i > 0 {
				coins *= nums[i-1]
			}
			if j < len(nums)-1 {
				coins *= nums[j+1]
			}
			return coins
		}

		if M[i][j] > 0 {
			return M[i][j]
		}

		xCoins := 0
		for k := i; k <= j; k++ {
			coins := nums[k]
			if i > 0 {
				coins *= nums[i-1]
			}
			if j < len(nums)-1 {
				coins *= nums[j+1]
			}

			xCoins = max(Search(i, k-1)+coins+Search(k+1, j), xCoins)
		}

		M[i][j] = xCoins
		return xCoins
	}

	return Search(0, len(nums)-1)
}

// 329h Longest Increasing Path in a Matrix
func longestIncreasingPath(matrix [][]int) int {
	D := make([][]int, len(matrix))
	for r := range D {
		D[r] = make([]int, len(matrix[r]))
	}

	var DFS func(r, c int)
	DFS = func(r, c int) {
		log.Printf("-> (%d %d)", r, c)

		if D[r][c] > 0 {
			return
		}

		v, steps := matrix[r][c], 1
		D[r][c] = 1

		Dirs := []int{-1, 0, 1, 0, -1}
		for d := range 4 {
			r, c := r+Dirs[d], c+Dirs[d+1]
			if 0 <= r && r < len(matrix) && 0 <= c && c < len(matrix[r]) && matrix[r][c] > v {
				log.Printf("-> . (%d %d)", r, c)

				if D[r][c] == 0 {
					DFS(r, c)
				}

				steps = max(D[r][c]+1, steps)
			}
		}

		D[r][c] = steps
	}

	for r := range matrix {
		for c := range matrix[r] {
			if D[r][c] == 0 {
				DFS(r, c)
			}
		}
	}

	log.Print("-> ", D)

	lPath := 0
	for _, row := range D {
		lPath = max(slices.Max(row), lPath)
	}
	return lPath
}

// 368m Largest Divisible Subset
func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	log.Print("-> ", nums)

	D := make([]int, len(nums))

	lMax, idx := 1, 0 // Maximum Subset Length, Last Index
	for i, N_i := range nums {
		D[i] = 1

		for j, N_j := range nums[:i] {
			if N_i%N_j == 0 {
				D[i] = max(D[j]+1, D[i])

				if D[i] > lMax {
					lMax, idx = D[i], i
				}
			}
		}
	}

	log.Print("-> ", lMax, idx, D)

	R := []int{}

	n := nums[idx]
	for idx >= 0 {
		if D[idx] == lMax && n%nums[idx] == 0 {
			R = append(R, nums[idx])
			n = nums[idx]
			lMax--
		}
		idx--
	}
	slices.Reverse(R)

	return R
}

// 377m Combination Sum IV
func combinationSum4(nums []int, target int) int {
	Sums := []int{1}
	for t := 1; t <= target; t++ {
		tsum := 0
		for _, n := range nums {
			if t-n >= 0 {
				tsum += Sums[t-n]
			}
		}

		Sums = append(Sums, tsum)
	}

	log.Print("-> ", Sums)
	return Sums[target]
}

// 494m Target Sum
func findTargetSumWays(nums []int, target int) int {
	M := map[[2]int]int{}

	var Search func(i, t int) int
	Search = func(i, t int) int {
		if i >= len(nums) {
			if t == 0 {
				return 1
			}
			return 0
		}

		if v, ok := M[[2]int{i, t}]; ok {
			return v
		}

		M[[2]int{i, t}] = Search(i+1, t+nums[i]) + Search(i+1, t-nums[i])
		return M[[2]int{i, t}]
	}

	return Search(0, target)
}

// 516m Longest Palindromic Subsequence
func longestPalindromeSubseq(s string) int {
	LPS := make([][]int, len(s))
	for r := range LPS {
		LPS[r] = make([]int, len(s))
	}

	for r := len(s) - 1; r >= 0; r-- {
		LPS[r][r] = 1

		for c := r + 1; c < len(s); c++ {
			if s[r] == s[c] {
				LPS[r][c] = LPS[r+1][c-1] + 2
			} else {
				LPS[r][c] = max(LPS[r+1][c], LPS[r][c-1])
			}
		}
	}

	return LPS[0][len(s)-1]
}

// 646m Maximum Length of Pair Chain
func findLongestChain(pairs [][]int) int {
	slices.SortFunc(pairs, func(a, b []int) int { return a[1] - b[1] })
	log.Print(pairs)

	D := make([]int, len(pairs))
	for i := range D {
		D[i] = 1
	}

	for r := 1; r < len(D); r++ {
		for l := 0; l < r; l++ {
			if pairs[r][0] > pairs[l][1] {
				D[r] = max(D[l]+1, D[r])
			}
		}
	}
	log.Print(D)
	return slices.Max(D)
}

// 664h Strange Printer
func strangePrinter(s string) int {
	Mem := map[[2]int]int{}
	defer log.Print(Mem)

	var MinTurn func(start, end int) int
	MinTurn = func(start, end int) int {
		if start > end {
			return 0
		}

		if v, ok := Mem[[2]int{start, end}]; ok {
			return v
		}

		v := 1 + MinTurn(start+1, end)
		for k := start + 1; k <= end; k++ {
			if s[start] == s[k] {
				v = min(MinTurn(start, k-1)+MinTurn(k+1, end), v)
			}
		}

		Mem[[2]int{start, end}] = v
		return v
	}

	return MinTurn(0, len(s)-1)
}

// 673m Number of Longest Increasing Subsequence
func findNumberOfLIS(nums []int) int {
	counter := make([]int, len(nums))
	for i := range counter {
		counter[i] = 1
	}

	D := make([]int, len(nums))
	for i := range D {
		D[i] = 1
	}

	for r := 1; r < len(nums); r++ {
		for l := 0; l < r; l++ {
			if nums[l] < nums[r] {
				if D[r] < D[l]+1 {
					D[r] = D[l] + 1
					counter[r] = counter[l]
				} else if D[r] == D[l]+1 {
					counter[r] += counter[l]
				}
			}
		}
	}

	log.Print(D, " -> ", counter)

	xVal := slices.Max(D)
	count := 0
	for i, n := range D {
		if n == xVal {
			count += counter[i]
		}
	}
	return count
}

// 689h Maximum Sum of 3 Non-Overlapping Subarrays
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	kSums := []int{}

	rSum := 0
	for i := range nums {
		rSum += nums[i]
		if i+1 >= k {
			kSums = append(kSums, rSum)
			rSum -= nums[i-k+1]
		}
	}

	log.Print(" -> ", kSums)

	D := make([][4]int, len(kSums))
	for i := range D {
		for r := range 4 {
			D[i][r] = -1
		}
	}

	var KS func(i, r int) int // 0:1 Knapsack
	KS = func(i, r int) int {
		if r == 0 {
			return 0
		}
		if i >= len(kSums) {
			return math.MinInt32
		}

		if D[i][r] != -1 {
			return D[i][r]
		}
		D[i][r] = max(kSums[i]+KS(i+k, r-1), KS(i+1, r))
		return D[i][r]
	}

	KS(0, 3)

	log.Print(" -> Mem :: ", D)

	R := []int{}

	var Trace func(i, r int)
	Trace = func(i, r int) {
		if i >= len(kSums) || r == 0 {
			return
		}

		take, ntake := kSums[i]+KS(i+k, r-1), KS(i+1, r)
		if take >= ntake {
			R = append(R, i)
			Trace(i+k, r-1)
		} else {
			Trace(i+1, r)
		}
	}

	Trace(0, 3)

	return R
}

// 730h Count Different Palindromic Subsequences
func countPalindromicSubsequences(s string) int {
	Mem := map[[2]int]int{}
	defer func() { log.Printf("-> #%d %v", len(Mem), Mem) }()

	const M = 1000_000_007

	var Search func(start, end int) int
	Search = func(start, end int) int {
		if start >= end {
			return 0
		}

		if count, ok := Mem[[2]int{start, end}]; ok {
			return count
		}

		count := 0
		for _, chr := range []byte("abcd") {
			l, r := strings.IndexByte(s[start:end], chr), strings.LastIndexByte(s[start:end], chr)
			if l == -1 || r == -1 {
				continue
			}

			if l == r {
				count++
			} else {
				count += 2 + Search(start+l+1, start+r)
			}

			count %= M
		}
		Mem[[2]int{start, end}] = count
		return count
	}

	return Search(0, len(s))
}

// 790m Domino and Tromino Tiling
func numTilings(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	const M = 1000_000_007

	// 1 <= n <= 1000
	TwoStates := func(n int) int {
		F, L := make([]int, 1000+1), make([]int, 1000+1) // Full & L-Shape Cover
		F[1], F[2], L[2] = 1, 2, 1

		for w := 3; w <= n; w++ { // all widths of 2xN
			F[w] = (F[w-1] + F[w-2] + 2*L[w-1]) % M
			L[w] = (L[w-1] + F[w-2]) % M
		}

		return F[n]
	}
	log.Print(":: ", TwoStates(n))

	D := make([]int, n+1)
	D[1], D[2], D[3] = 1, 2, 5

	for i := 4; i <= n; i++ {
		D[i] = (2*D[i-1] + D[i-3]) % M
	}

	return D[n]
}

// 983m Minimum Cost For Tickets
func mincostTickets(days []int, costs []int) int {
	Mem := map[int]int{}

	TDays := make([]bool, 365+1)
	for _, day := range days {
		TDays[day] = true
	}

	var Search func(d int) int
	Search = func(d int) int {
		if d > 365 {
			return 0
		}

		if !TDays[d] {
			return Search(d + 1)
		}

		if v, ok := Mem[d]; ok {
			return v
		}

		mCost := 0
		mCost = min(costs[0]+Search(d+1), costs[1]+Search(d+7), costs[2]+Search(d+30))

		Mem[d] = mCost
		return mCost
	}

	return Search(days[0])
}

// 1014m Best Sightseeing Pair
func maxScoreSightseeingPair(values []int) int {
	// Score: i < j :: Vi+Vj - (j-i)
	n := len(values)

	D := make([]int, n)
	D[n-1] = values[n-1] - (n - 1) // Vj - j
	for j := n - 2; j > 0; j-- {
		D[j] = max(D[j+1], values[j]-j)
	}

	xVal := 0
	for i := range values[:n-1] {
		xVal = max(xVal, D[i+1]+values[i]+i)
	}
	return xVal
}

// 1025 Divisor Game
func divisorGame(n int) bool {
	D := make([]bool, n+1)

	for N := 2; N <= n; N++ {
		for x := N / 2; x >= 1; x-- {
			if N%x != 0 {
				continue
			}

			if !D[N-x] {
				D[N] = true
			}
		}
	}

	return D[n]
}

// 1092h Shortest Common Supersequence
func shortestCommonSupersequence(str1 string, str2 string) string {
	LCS := make([][]int, len(str1)+1)
	for p := range LCS {
		LCS[p] = make([]int, len(str2)+1)
	}

	for p := range len(str1) {
		for q := range len(str2) {
			if str1[p] == str2[q] {
				LCS[p+1][q+1] = LCS[p][q] + 1
			} else {
				LCS[p+1][q+1] = max(LCS[p][q+1], LCS[p+1][q])
			}
		}
	}

	// Longest Common Subsequence
	log.Print("-> ", LCS)

	SCS := []byte{}

	p, q := len(str1), len(str2)
	for p > 0 || q > 0 {
		if p > 0 && q > 0 && str1[p-1] == str2[q-1] {
			SCS = append(SCS, str1[p-1])
			p--
			q--
		} else if p > 0 && (q == 0 || LCS[p-1][q] >= LCS[p][q-1]) {
			SCS = append(SCS, str1[p-1])
			p--
		} else if q > 0 {
			SCS = append(SCS, str2[q-1])
			q--
		}
	}

	slices.Reverse(SCS)
	return string(SCS)
}

// 1105m Filling Bookcase Shelves
func minHeightShelves(books [][]int, shelfWidth int) int {
	rCalls, Mem := 0, map[[3]int]int{}

	defer func() {
		log.Print("-> ", rCalls, " :: ", Mem)
	}()

	var Check func(i, curW, curH int) int
	Check = func(i, curW, curH int) int {
		rCalls++

		if i == len(books) {
			return curH
		}

		if v, ok := Mem[[3]int{i, curW, curH}]; ok {
			return v
		}

		book := books[i]

		v := Check(i+1, shelfWidth-book[0], book[1]) + curH // Go next row
		if book[0] <= curW {                                // Stay in current row
			v = min(Check(i+1, curW-book[0], max(curH, book[1])), v)
		}

		Mem[[3]int{i, curW, curH}] = v
		return v
	}

	return Check(0, shelfWidth, 0)
}

// 1155m Number of Dice Rolls With Target Sum
func numRollsToTarget(n, k, target int) int {
	D := make([][]int, n)
	for d := range D {
		D[d] = make([]int, max(target, k)+1)
	}

	for r := range k {
		D[0][r+1] = 1
	}

	const M = 1000_000_007

	for d := 1; d < n; d++ {
		for v := 1; v <= target; v++ {
			for r := 1; r <= k && r+v <= target; r++ {
				D[d][r+v] += D[d-1][v]
				D[d][r+v] %= M
			}
		}
	}

	return D[n-1][target]
}

// 1277m Count Square Submatrices with All Ones
func countSquares(matrix [][]int) int {
	Rows, Cols := len(matrix), len(matrix[0])

	D := make([][]int, Rows+1)
	for r := range D {
		D[r] = make([]int, Cols+1)
	}

	count := 0
	for r := range Rows {
		for c := range Cols {
			if matrix[r][c] == 1 {
				D[r+1][c+1] = 1 + min(D[r+1][c], D[r][c+1], D[r][c])
				count += D[r+1][c+1]
			}
		}
	}
	return count
}

// 1395m Count Number of Teams
func numTeams(rating []int) int {
	x := 0

	for m := 1; m < len(rating)-1; m++ {
		l, r := 0, 0

		for i := 0; i < m; i++ {
			if rating[i] < rating[m] {
				l++
			}
		}

		for i := m + 1; i < len(rating); i++ {
			if rating[m] < rating[i] {
				r++
			}
		}

		x += l * r                               // Rating[l] < Rating[m] < Rating[r]
		x += (m - l) * (len(rating) - m - 1 - r) // Rating[l] > Raring[m] > Rating[r]
	}

	return x
}

// 1524m Number of Sub-arrays With Odd Sum
func numOfSubarrays(arr []int) int {
	const M = 1000_000_000 + 7

	E, O := make([]int, len(arr)), make([]int, len(arr))

	if arr[len(arr)-1]&1 == 1 {
		O[len(arr)-1] = 1
	} else {
		E[len(arr)-1] = 1
	}

	for i := len(arr) - 2; i >= 0; i-- {
		switch arr[i] & 1 {
		case 1:
			O[i] = (1 + E[i+1]) % M
			E[i] = O[i+1]
		case 0:
			O[i] = O[i+1]
			E[i] = (1 + E[i+1]) % M
		}
	}

	count := 0
	for _, n := range O {
		count += n % M
	}
	return count % M
}

// 1639h Number of Ways to Form a Target String Given a Dictionary
func numWays(words []string, target string) int {
	const MOD = 1e9 + 7

	W := len(words[0])
	F := make([][26]int, W)
	for p := 0; p < W; p++ {
		for w := range words {
			F[p][words[w][p]-'a']++
		}
	}

	Mem := map[[2]int]int64{}

	var Search func(w, t int) int64
	Search = func(w, t int) int64 {
		if t == len(target) {
			return 1
		}
		if w == W {
			return 0
		}

		if v, ok := Mem[[2]int{w, t}]; ok {
			return v
		}

		cSum := Search(w+1, t) % MOD
		cSum += int64(F[w][target[t]-'a']) * Search(w+1, t+1) % MOD
		cSum %= MOD

		Mem[[2]int{w, t}] = cSum
		return cSum
	}

	return int(Search(0, 0))
}

// 1653m Minimum Deletions to Make String Balanced
func minimumDeletions(s string) int {
	A := make([]int, len(s))

	x := 0
	for i := len(s) - 1; i >= 0; i-- {
		A[i] = x
		if s[i] == 'a' {
			x++
		}
	}

	dels := len(s)
	for Bi, i := 0, 0; i < len(s); i++ {
		if Bi+A[i] < dels {
			dels = Bi + A[i]
		}
		if s[i] == 'b' {
			Bi++
		}
	}
	return dels
}

// 1749m Maximum Absolute Sum of Any Subarray
func maxAbsoluteSum(nums []int) int {
	xVal := 0

	xSum, nSum := math.MinInt, math.MaxInt
	pSum := 0
	for _, n := range nums {
		pSum += n

		xSum, nSum = max(pSum, xSum), min(pSum, nSum)

		if pSum > 0 {
			xVal = max(max(pSum, pSum-nSum), xVal)
		} else if pSum < 0 {
			xVal = max(max(-pSum, xSum-pSum), xVal)
		}
	}

	return xVal
}

// 1931m Painting a Grid With Three Different Colors
func colorTheGrid(m int, n int) int {
	Masks := map[int][]int{}
	for mask := range []int{1, 3, 9, 27, 81, 243}[m] {
		v, colors := mask, []int{}
		for range m {
			colors = append(colors, v%3)
			v /= 3
		}

		valid := true
		for i := 1; i < len(colors); i++ {
			if colors[i-1] == colors[i] {
				valid = false
			}
		}
		if valid {
			Masks[mask] = colors
		}
	}
	log.Print("-> Masks: ", Masks)

	Adjs := map[int][]int{}
	for mask, colors := range Masks {
		for adjMask, adjColors := range Masks {
			valid := true
			for r := range colors {
				if colors[r] == adjColors[r] {
					valid = false
				}
			}
			if valid {
				Adjs[mask] = append(Adjs[mask], adjMask)
			}
		}
	}
	log.Print("-> Adjacent Cols: ", Adjs)

	const M = 1000_000_007

	dpCur := make([]int, []int{1, 3, 9, 27, 81, 243}[m])
	for mask := range Masks {
		dpCur[mask] = 1
	}
	for range n - 1 {
		dpNext := make([]int, len(dpCur))
		for mask, count := range dpCur {
			if count > 0 {
				count := 0
				for _, adjMask := range Adjs[mask] {
					count = (count + dpCur[adjMask]) % M
				}
				dpNext[mask] = count
			}
		}
		dpCur = dpNext
	}

	total := 0
	for _, count := range dpCur {
		total = (total + count) % M
	}

	return total
}

// 1937m Maximum Number of Points with Cost
func maxPoints(points [][]int) int64 {
	Rows, Cols := len(points), len(points[0])
	prv := append([]int{}, points[0]...)

	for r := 1; r < Rows; r++ {
		cur := make([]int, Cols)

		left := make([]int, Cols)
		left[0] = prv[0]
		for c := 1; c < Cols; c++ {
			left[c] = max(prv[c], left[c-1]-1)
		}

		right := make([]int, Cols)
		right[Cols-1] = prv[Cols-1]
		for c := Cols - 2; c >= 0; c-- {
			right[c] = max(prv[c], right[c+1]-1)
		}

		for c := 0; c < Cols; c++ {
			cur[c] = points[r][c] + max(left[c], right[c])
		}
		prv = cur
	}

	return int64(slices.Max(prv))
}

// 2016 Maximum Difference Between Increasing Elements
func maximumDifference(nums []int) int {
	xVal := -1

	nVal := nums[0]
	for i := range nums[:len(nums)-1] {
		nVal = min(nVal, nums[i])
		if nums[i+1] > nVal {
			xVal = max(xVal, nums[i+1]-nVal)
		}
	}

	return xVal
}

// 2140m Solving Questions With Brainpower
func mostPoints(questions [][]int) int64 {
	D := make([][2]int64, len(questions)+1)

	for i := len(questions) - 1; i >= 0; i-- {
		pts, skip := questions[i][0], questions[i][1]

		D[i][0] = max(D[i+1][0], D[i+1][1])
		D[i][1] = int64(pts)
		next := i + skip + 1
		if next < len(questions) {
			D[i][1] += max(D[next][0], D[next][1])
		}
	}

	log.Print("-> ", D)

	return slices.Max(D[0][:])
}

// 2466m Count Ways to Build Good Strings
func countGoodStrings(low, high, zero, one int) int {
	const MOD = 1e9 + 7

	D := make([]int, high+1)
	D[0] = 1

	for l := 0; l <= high; l++ {
		if l >= zero {
			D[l] += D[l-zero]
		}
		if l >= one {
			D[l] += D[l-one]
		}

		D[l] %= MOD
	}

	tWays := 0
	for l := low; l <= high; l++ {
		tWays += D[l]
		tWays %= MOD
	}

	return tWays
}

// 2707m Extra Characters in a String
func minExtraChar(s string, dictionary []string) int {
	M := map[int]int{}
	defer log.Print(" -> ", M)

	D := map[string]struct{}{}
	for _, w := range dictionary {
		D[w] = struct{}{}
	}

	var W func(int) int
	W = func(start int) int {
		if start >= len(s) {
			return 0
		}

		if v, ok := M[start]; ok {
			return v
		}

		v := len(s)
		for i := start; i < len(s); i++ {
			for w := range D {
				if strings.HasPrefix(s[i:], w) {
					v = min(v, i-start+W(i+len(w)))
				}
			}
		}

		M[start] = v
		return v
	}

	return W(0)
}

// 2836h Maximize Value of Function in a Ball Passing Game
func getMaxFunctionValue(receiver []int, k int64) int64 {
	B := 0 // bits
	for x := k; x > 0; x >>= 1 {
		B++
	}

	N := len(receiver)

	// Jumps & Scores
	far, score := make([][]int, N), make([][]int64, N)
	for n := range N {
		far[n] = make([]int, B)
		score[n] = make([]int64, B)
	}

	for p := range B {
		for i := range N {
			switch p {
			case 0:
				far[i][0] = receiver[i]
				score[i][0] = int64(receiver[i])
			default:
				far[i][p] = far[far[i][p-1]][p-1]
				score[i][p] = score[i][p-1] + score[far[i][p-1]][p-1]
			}
		}
	}

	log.Print(" -> ", far)
	log.Print(" -> ", score)

	xScore := int64(0)
	for istart := range N {
		iScore, i := int64(0), istart
		for p := range B {
			if 1<<p&k != 0 {
				iScore += score[i][p]
				i = far[i][p]
			}
		}
		xScore = max(iScore+int64(istart), xScore)
	}

	return xScore
}

// 2901m Longest Unequal Adjacent Groups Subsequence II
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	ValidDist := func(i, j int) bool {
		if len(words[j]) != len(words[i]) {
			return false
		}

		dist, x, y := 0, words[i], words[j]
		for i := 0; i < len(x)&len(y); i++ {
			if x[i] != y[i] {
				dist++
				if dist > 1 {
					return false
				}
			}
		}

		return true
	}

	D, Picks := make([]int, len(groups)), make([]int, len(groups))
	for i := range D {
		D[i], Picks[i] = 1, -1
	}

	lMax, iLast := 1, 0
	for i := 1; i < len(D); i++ {
		for j := 0; j < i; j++ {
			if groups[j] != groups[i] && ValidDist(j, i) {
				if D[j]+1 > D[i] {
					D[i], Picks[i] = D[j]+1, j
				}
			}
		}

		if D[i] > lMax {
			lMax, iLast = D[i], i
		}
	}

	R := []string{}
	for iLast != -1 {
		R, iLast = append(R, words[iLast]), Picks[iLast]
	}
	slices.Reverse(R)

	log.Printf(":: %q", R)

	return R
}

// 2999h Count the Number of Powerful Integers
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	M := map[int]int64{}
	defer log.Print("-> ", M)

	to := fmt.Sprintf("%d", finish)
	W := len(to)
	from := fmt.Sprintf("%0[2]*[1]d", start, W)

	log.Printf("-> [%s, %s]", from, to)

	var Search func(p int, lower, higher bool) int64
	Search = func(p int, lower, higher bool) int64 {
		if p == W {
			return 1
		}
		if count, ok := M[p]; ok && !lower && !higher {
			return count
		}

		ldigit, hdigit := byte('0'), byte('9')
		if lower {
			ldigit = from[p]
		}
		if higher {
			hdigit = to[p]
		}

		count := int64(0)
		if p < W-len(s) {
			for digit := ldigit; digit <= min(hdigit, byte(limit)+'0'); digit++ {
				count += Search(p+1, lower && digit == from[p], higher && digit == to[p])
			}
		} else {
			digit := s[p-(W-len(s))]
			if ldigit <= digit && digit <= min(hdigit, byte(limit)+'0') {
				count += Search(p+1, lower && digit == from[p], higher && digit == to[p])
			}
		}

		if !lower && !higher {
			M[p] = count
		}
		return count
	}

	return Search(0, true, true)
}

// 3068h Find the Maximum Sum of Node Values
func maximumValueSum_3068(nums []int, k int, edges [][]int) int64 {
	Recursive := func() int64 {
		M := map[[2]int]int64{}
		defer log.Print("-> Mem: ", M)

		var Search func(start, xOr int) int64
		Search = func(start, xOr int) int64 {
			if start == len(nums) {
				if xOr&1 == 0 {
					return 0
				}
				return math.MinInt64
			}

			if v, ok := M[[2]int{start, xOr}]; ok {
				return v
			}

			xVal := int64(nums[start]) + Search(start+1, xOr)
			xVal = max(int64(nums[start]^k)+Search(start+1, xOr^1), xVal)

			M[[2]int{start, xOr}] = xVal
			return xVal
		}

		return Search(0, 0)
	}
	log.Print(":: ", Recursive())

	DP := func() int64 {
		D := make([][2]int64, len(nums)+1)
		D[len(nums)][0], D[len(nums)][1] = 0, math.MinInt64

		for v := len(nums) - 1; v >= 0; v-- {
			for xOr := range 2 {
				D[v][xOr] = max(int64(nums[v])+D[v+1][xOr], int64(nums[v]^k)+D[v+1][xOr^1])
			}
		}

		log.Print("-> ", D)

		return D[0][0]
	}
	log.Print(":: ", DP())

	Diff := make([]int, len(nums))
	for i, n := range nums {
		Diff[i] = n ^ k - n
	}
	slices.Sort(Diff)
	slices.Reverse(Diff)

	log.Print("-> ", Diff)

	xVal := int64(0)
	for _, n := range nums {
		xVal += int64(n)
	}

	for i := 1; i < len(nums); i += 2 {
		if Diff[i-1]+Diff[i] > 0 {
			xVal += int64(Diff[i-1] + Diff[i])
		}
	}
	return xVal
}

// 3335m Total Characters in String After Transformations I
func lengthAfterTransformations(s string, t int) int {
	const M = 1000_000_007

	Optimized := func(s string, t int) int {
		D := [26]int{}
		for i := 0; i < len(s); i++ {
			D[s[i]-'a']++
		}

		N := [26]int{}
		for range t {
			N[0] = D[25]
			N[1] = (D[0] + D[25]) % M
			for chr := 2; chr < 26; chr++ {
				N[chr] = D[chr-1]
			}

			D = N
		}

		total := 0
		for _, n := range D {
			total = (total + n) % M
		}

		return total
	}
	log.Print(":: ", Optimized(s, t))

	D := make([][26]int, t+1)
	for i := 0; i < len(s); i++ {
		D[0][s[i]-'a']++
	}

	for i := 1; i <= t; i++ {
		D[i][0] = D[i-1][25]                   // z -> (a)b
		D[i][1] = (D[i-1][0] + D[i-1][25]) % M // z -> a(b) & a -> b
		for chr := 2; chr < 26; chr++ {
			D[i][chr] = D[i-1][chr-1]
		}
	}

	log.Print("-> ", D)

	total := 0
	for _, n := range D[t] {
		total = (total + n) % M
	}

	return total
}

// 3337h Total Characters in String After Transformations II
func lengthAfterTransformationsII(s string, t int, nums []int) int {
	F := make([]int, 26)
	for i := 0; i < len(s); i++ {
		F[s[i]-'a']++
	}

	T := make([][26]int, 26)
	for chr := range T {
		for n := 1; n <= nums[chr]; n++ {
			T[chr][(chr+n)%26] = 1
		}
	}

	const M = 1000_000_007

	Multiply := func(a, b [][26]int) [][26]int {
		m := make([][26]int, 26)
		for i := range 26 {
			for k := range 26 {
				a_ik := a[i][k]
				if a_ik != 0 {
					for j := range 26 {
						m[i][j] = (m[i][j] + a_ik*b[k][j]%M) % M
					}
				}
			}
		}
		return m
	}

	P := make([][26]int, 26)
	for d := range 26 {
		P[d][d] = 1
	}

	for t > 0 {
		if t&1 == 1 {
			P = Multiply(P, T)
		}
		T = Multiply(T, T)
		t >>= 1
	}

	log.Print("-> T ", T)

	Ft := [26]int{}
	for i := range 26 {
		for j := range 26 {
			Ft[i] = (Ft[i] + F[i]*P[i][j]%M) % M
		}
	}

	log.Print("-> ", Ft)

	total := 0
	for _, f := range Ft {
		total = (total + f) % M
	}
	return total
}

// 3343h Count Number of Balanced Permutations
func countBalancedPermutations(num string) int {
	dSum := 0
	F := make([]int, 10)
	for i := 0; i < len(num); i++ {
		F[num[i]-'0']++
		dSum += int(num[i] - '0')
	}
	log.Print("-> ", F)

	if dSum&1 == 1 {
		return 0
	}

	const MOD = 1000_000_007

	C := make([][]int, (len(num)+1)/2+1) // Pascal's Triangle: C(n,k) = C(n-1,k) + C(n-1,k-1)
	for n := range C {
		C[n] = make([]int, (len(num)+1)/2+1)
		C[n][n], C[n][0] = 1, 1
		for k := 1; k < n; k++ {
			C[n][k] = (C[n-1][k] + C[n-1][k-1]) % MOD
		}
	}
	log.Print("-> nCk C(n,k) ", C)

	pSum := make([]int, 11)
	for i := 9; i >= 0; i-- {
		pSum[i] = pSum[i+1] + F[i]
	}

	M := make([][][]int, 10) // Memo: {pos,curSum,oddCount}
	for i := range M {
		M[i] = make([][]int, dSum/2+1)
		for j := range M[i] {
			M[i][j] = make([]int, (len(num)+1)/2+1)
			for k := range M[i][j] {
				M[i][j][k] = -1
			}
		}
	}
	defer log.Print("-> ", M)

	var Search func(p, curSum, oddCount int) int
	Search = func(p, curSum, oddCount int) int {
		if oddCount < 0 || pSum[p] < oddCount || dSum/2 < curSum {
			return 0
		}
		if p > 9 {
			if curSum == dSum/2 && oddCount == 0 {
				return 1
			}
			return 0
		}

		if M[p][curSum][oddCount] != -1 {
			return M[p][curSum][oddCount]
		}

		evenCount := pSum[p] - oddCount
		r := 0
		for i := max(0, F[p]-evenCount); i <= min(F[p], oddCount); i++ {
			r += (C[oddCount][i] * C[evenCount][F[p]-i] % MOD) * Search(p+1, curSum+i*p, oddCount-i) % MOD
			r %= MOD
		}

		M[p][curSum][oddCount] = r
		return r
	}

	return Search(0, 0, (len(num)+1)/2)
}
