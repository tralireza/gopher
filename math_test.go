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

// 326 Power of Three
func Test326(t *testing.T) {
	log.Print("true ? ", isPowerOfThree(27))
	log.Print("false ? ", isPowerOfThree(0))
	log.Print("false ? ", isPowerOfThree(-1))
	log.Print("true ? ", isPowerOfThree(1))
}

// 1998h GCD Sort of an Array
func Test1998(t *testing.T) {
	// 2 <= N_i <= 10^5, N.length <= 3*10^4

	log.Println("true ?= ", gcdSort([]int{7, 21, 3}))
	log.Println("false ?= ", gcdSort([]int{5, 2, 6, 2}))
	log.Println("true ?= ", gcdSort([]int{10, 5, 9, 3, 15}))
}

// 3312h Sorted GCD Pair Queries
func Test3312(t *testing.T) {
	// 1 <= N_i <= 5*10^4, N.length <= 10^5

	log.Print("[1 2 2] ?= ", gcdValues([]int{2, 3, 4}, []int64{0, 2, 2}))
	log.Print("[4 2 1 1] ?= ", gcdValues([]int{4, 4, 2, 1}, []int64{5, 3, 1, 0}))
	log.Print("[2 2] ?= ", gcdValues([]int{2, 2}, []int64{0, 0}))
}
