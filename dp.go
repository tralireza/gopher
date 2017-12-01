package gopher

import (
	"log"
	"slices"
)

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
		for c := 2; c < Cols; c++ {
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
