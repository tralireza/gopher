package gopher

import (
	"log"
	"testing"
)

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
		for start := 0; start < len(rating); start++ {
			teams += Count(start, 1)
		}
		log.Print(Mem)

		clear(Mem)
		for i := range rating {
			rating[i] *= -1
		}

		for start := 0; start < len(rating); start++ {
			teams += Count(start, 1)
		}
		log.Print(Mem)

		return teams
	}

	for _, f := range []func([]int) int{numTeams, Recursion} {
		log.Print("3 ?= ", f([]int{2, 5, 3, 4, 1}))
		log.Print("0 ?= ", f([]int{2, 1, 3}))
		log.Print("4 ?= ", f([]int{1, 2, 3, 4}))
		log.Print("--")
	}
}
