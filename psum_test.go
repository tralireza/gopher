package gopher

import (
	"log"
	"testing"
)

// 1991 Find the Middle Index in Array
func Test1991(t *testing.T) {
	log.Print("3 ?= ", findMiddleIndex([]int{2, 3, -1, 8, 4}))
	log.Print("2 ?= ", findMiddleIndex([]int{1, -1, 4}))
	log.Print("-1 ?= ", findMiddleIndex([]int{2, 5}))
}

// 2134m Minimum Swaps to Group All 1's Together II
func Test2134(t *testing.T) {
	// Ai e {0, 1}

	log.Print("1 ?= ", minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	log.Print("2 ?= ", minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
	log.Print("0 ?= ", minSwaps([]int{1, 1, 0, 0, 1}))
}

// 2574 Left and Right Sum Difference
func Test2574(t *testing.T) {
	log.Print("[15 1 11 22] ?= ", leftRightDifference([]int{10, 4, 8, 3}))
	log.Print("[0] ?= ", leftRightDifference([]int{1}))
}
