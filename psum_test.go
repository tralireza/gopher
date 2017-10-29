package gopher

import (
	"log"
	"testing"
)

// 2134m Minimum Swaps to Group All 1's Together II
func Test2134(t *testing.T) {
	// Ai e {0, 1}

	log.Print("1 ?= ", minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	log.Print("2 ?= ", minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
	log.Print("0 ?= ", minSwaps([]int{1, 1, 0, 0, 1}))
}
