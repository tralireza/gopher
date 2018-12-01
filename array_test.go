package gopher

import (
	"log"
	"testing"
)

// 485 Max Consecutive Ones
func Test485(t *testing.T) {
	log.Print("3 ?= ", findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	log.Print("2 ?= ", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
