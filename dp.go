package gopher

import "log"

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
