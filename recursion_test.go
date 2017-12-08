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
	log.Print("3 ?= ", minSteps(3))
	log.Print("0 ?= ", minSteps(1))
}
