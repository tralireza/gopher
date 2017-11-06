package gopher

import (
	"log"
	"testing"
)

// 239h Sliding Window Maximum
func Test239(t *testing.T) {
	log.Print("[3 3 5 5 6 7] ?= ", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	log.Print("[1 -1] ?= ", maxSlidingWindow([]int{1, -1}, 1))
}

// 1508m Range Sum of Sorted Subarray Sums
func Test1508(t *testing.T) {
	log.Print("13 ?= ", rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
	log.Print("6 ?= ", rangeSum([]int{1, 2, 3, 4}, 4, 3, 4))
	log.Print("50 ?= ", rangeSum([]int{1, 2, 3, 4}, 4, 1, 10))
}
