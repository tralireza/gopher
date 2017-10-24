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
