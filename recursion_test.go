package gopher

import (
	"log"
	"testing"
)

// 40m Combination Sum II
func Test40(t *testing.T) {
	log.Print(" ?= ", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	log.Print(" ?= ", combinationSum2([]int{2, 5, 2, 1, 2}, 5))
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
