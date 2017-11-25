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
