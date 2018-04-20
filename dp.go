package gopher

import (
	"container/heap"
	"log"
	"slices"
	"sort"
)

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
