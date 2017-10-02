package gopher

import (
	"log"
	"testing"
)

// 3224m Minimum Array Changes to Make Difference Equal
func Test3224(t *testing.T) {
	// 0 <= Array[i] <= k <= 10^5

	log.Print("2 ?= ", minChanges([]int{1, 0, 1, 2, 4, 3}, 4))
	log.Print("2 ?= ", minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6))
}
