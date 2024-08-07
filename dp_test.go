package gopher

import (
	"log"
	"testing"
)

// 122m Best Time to Buy and Sell Stock II
func Test122(t *testing.T) {
	log.Print("7 ?= ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Print("4 ?= ", maxProfit([]int{1, 2, 3, 4, 5}))
	log.Print("0 ?= ", maxProfit([]int{7, 6, 4, 3, 1}))
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
