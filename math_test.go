package gopher

import (
	"log"
	"testing"
)

// 319m Bulb Switcher
func Test319(t *testing.T) {
	log.Print("1 ?= ", bulbSwitch(3))
	log.Print("0 ?= ", bulbSwitch(0))
	log.Print("1 ?= ", bulbSwitch(1))

	log.Print("3 ?= ", bulbSwitch(9))
	log.Print("5 ?= ", bulbSwitch(27))
}

// 3312h Sorted GCD Pair Queries
func Test3312(t *testing.T) {
	// 1 <= N_i <= 5*10^4, N.length <= 10^5

	log.Print("[1 2 2] ?= ", gcdValues([]int{2, 3, 4}, []int64{0, 2, 2}))
	log.Print("[4 2 1 1] ?= ", gcdValues([]int{4, 4, 2, 1}, []int64{5, 3, 1, 0}))
	log.Print("[2 2] ?= ", gcdValues([]int{2, 2}, []int64{0, 0}))
}
