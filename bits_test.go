package gopher

import (
	"log"
	"math/bits"
	"testing"
)

// 401 Binary Watch
func Test401(t *testing.T) {
	// 0 <= N <= 9

	log.Printf(" ?= %q", readBinaryWatch(1))
	log.Printf(" ?= %q", readBinaryWatch(8))
	log.Printf(" ?= %q", readBinaryWatch(9))
}

// 461 Hamming Distance
func Test461(t *testing.T) {
	Bits := func(x, y int) int {
		return bits.OnesCount(uint(x|y)) - bits.OnesCount(uint(x&y))
	}

	for _, f := range []func(int, int) int{Bits, hammingDistance} {
		log.Print("2 ?= ", f(1, 4))
		log.Print("1 ?= ", f(3, 1))
		log.Print("--")
	}
}

// 476 Number Complement
func Test476(t *testing.T) {
	// 1 <= n <= 2^31

	log.Print("C(7) -> ", (1<<bits.Len(7)-1)^7)
	log.Print("C(5) -> ", (1<<bits.Len(5)-1)^5)
	log.Print("--")

	log.Print("2 ?= ", findComplement(5))
	log.Print("0 ?= ", findComplement(1))
	log.Print("1 ?= ", findComplement(2))
}

// 2429m Minimize XOR
func Test2429(t *testing.T) {
	log.Print("3 ?= ", minimizeXor(3, 5))
	log.Print("3 ?= ", minimizeXor(1, 12))
}

// 2657m Find the Prefix Common Array of Two Arrays
func Test2657(t *testing.T) {
	log.Print("[0 2 3 4] ?= ", findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	log.Print("[0 1 3] ?= ", findThePrefixCommonArray([]int{2, 3, 1}, []int{3, 1, 2}))
}

// 3315m Construct the Minimum Bitwise Array II
func Test3315(t *testing.T) {
	// 2 <= N_i <= 10^9

	log.Print("[-1 1 4 3] ?= ", minBitwiseArray([]int{2, 3, 5, 7}))
	log.Print("[9 12 15] ?= ", minBitwiseArray([]int{11, 13, 31}))
}
