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

// 342 Power of Four
func Test342(t *testing.T) {
	log.Print("true ?= ", isPowerOfFour(16))
	log.Print("false ?= ", isPowerOfFour(5))
	log.Print("true ?= ", isPowerOfFour(1))
}

// 598 Range Addition II
func Test598(t *testing.T) {
	log.Print("4 ?= ", maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
	log.Print("4 ?= ", maxCount(3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}))
	log.Print("9 ?= ", maxCount(3, 3, [][]int{}))
}

// 908 Smallest Range I
func Test908(t *testing.T) {
	log.Print("0 ?= ", smallestRangeI([]int{1}, 0))
	log.Print("6 ?= ", smallestRangeI([]int{0, 10}, 2))
	log.Print("0 ?= ", smallestRangeI([]int{1, 3, 6}, 3))
}

// 989 Add to Array-Form of Integer
func Test989(t *testing.T) {
	log.Print("[1 2 3 4] ?= ", addToArrayForm([]int{1, 2, 0, 0}, 34))
	log.Print("[4 5 5] ?= ", addToArrayForm([]int{2, 7, 4}, 181))
	log.Print("[1 0 2 1] ?= ", addToArrayForm([]int{2, 1, 5}, 806))
}

// 1154 Day of the Year
func Test1154(t *testing.T) {
	log.Print("9 ?= ", dayOfYear("2019-01-09"))
	log.Print("41 ?= ", dayOfYear("2019-02-10"))
}

// 1998h GCD Sort of an Array
func Test1998(t *testing.T) {
	// 2 <= N_i <= 10^5, N.length <= 3*10^4

	log.Println("true ?= ", gcdSort([]int{7, 21, 3}))
	log.Println("false ?= ", gcdSort([]int{5, 2, 6, 2}))
	log.Println("true ?= ", gcdSort([]int{10, 5, 9, 3, 15}))
}

// 3151 Special Array I
func Test3151(t *testing.T) {
	log.Print("true ?= ", isArraySpecialI([]int{1}))
	log.Print("true ?= ", isArraySpecialI([]int{2, 1, 4}))
	log.Print("false ?= ", isArraySpecialI([]int{4, 3, 1, 6}))
}

// 3312h Sorted GCD Pair Queries
func Test3312(t *testing.T) {
	// 1 <= N_i <= 5*10^4, N.length <= 10^5

	log.Print("[1 2 2] ?= ", gcdValues([]int{2, 3, 4}, []int64{0, 2, 2}))
	log.Print("[4 2 1 1] ?= ", gcdValues([]int{4, 4, 2, 1}, []int64{5, 3, 1, 0}))
	log.Print("[2 2] ?= ", gcdValues([]int{2, 2}, []int64{0, 0}))
}
