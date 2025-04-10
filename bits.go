// Bitwise
package gopher

import (
	"fmt"
	"log"
	"math/bits"
)

// 401 Binary Watch
func readBinaryWatch(turnedOn int) []string {
	R := []string{}

HOUR:
	for H := 0; H < 12; H++ {

	MINUTE:
		for MM := 0; MM < 60; MM++ {
			var mask int

			leds := 0

			mask = 1
			for p := range 4 {
				leds += (H & mask) >> p
				mask *= 2

				if leds > turnedOn {
					continue HOUR
				}
			}

			mask = 1
			for p := range 6 {
				leds += (MM & mask) >> p
				mask *= 2

				if leds > turnedOn {
					continue MINUTE
				}
			}

			if leds == turnedOn {
				R = append(R, fmt.Sprintf("%d:%02d", H, MM))
			}
		}
	}

	return R
}

// 476 Number Complement
func findComplement(num int) int {
	bits := 0
	for x := num; x > 0; x >>= 1 {
		bits++
	}
	return (1<<bits - 1) ^ num
}

// 461 Hamming Distance
func hammingDistance(x, y int) int {
	dist := 0

	for x > 0 || y > 0 {
		dist += x&1 ^ y&1
		x >>= 1
		y >>= 1
	}

	return dist
}

// 693 Binary Number with Alternating Bits
func hasAlternatingBits(n int) bool {
	p := ^(n & 1)
	for n > 0 {
		if p^(n&1) == 0 {
			return false
		}
		p = n & 1
		n >>= 1
	}
	return true
}

// 868 Binary Gap
func binaryGap(n int) int {
	dist, cur := 0, -32
	for n > 0 {
		cur++
		if cur > dist {
			dist = cur
		}

		if n&1 == 1 {
			cur = 0
		}
		n >>= 1
	}

	return dist
}

// 2429m Minimize XOR
func minimizeXor(num1, num2 int) int {
	CountSetBits := func(n int) int {
		bitsSet := 0
		for n > 0 {
			bitsSet += n & 1
			n >>= 1
		}
		return bitsSet
	}

	x := num1
	bitsSet := CountSetBits(x)

	log.Print(" -> ", bitsSet, bits.OnesCount(uint(x)))

	p := 0 // Bit Position

	for bitsSet < CountSetBits(num2) {
		if x&(1<<p) == 0 {
			x |= 1 << p // Set at <p>
			bitsSet++
		}
		p++
	}

	for bitsSet > CountSetBits(num2) {
		if x&(1<<p) != 0 {
			x &= ^(1 << p) // Unset at <p>
			bitsSet--
		}
		p++
	}

	return x
}

// 2657m Find the Prefix Common Array of Two Arrays
func findThePrefixCommonArray(A []int, B []int) []int {
	R := []int{}

	amask, bmask := uint(0), uint(0)
	for i := range A {
		amask += 1 << A[i]
		bmask += 1 << B[i]

		R = append(R, bits.OnesCount(amask&bmask))
	}

	return R
}

// 2683m Neighboring Bitwise XOR
func doesValidArrayExist(derived []int) bool {
	v := 0
	for _, d := range derived[:len(derived)-1] {
		v ^= d
	}

	return derived[len(derived)-1]^v == 0
}

// 3315m Construct the Minimum Bitwise Array II
func minBitwiseArray(nums []int) []int {
	R := []int{}

	for _, p := range nums {
		r := -1
		for i := 31; i >= 0; i-- {
			if p&(1<<i) != 0 {
				n := p & ^(1 << i)

				log.Print(" -> ", p, i, n)

				if n|(n+1) == p {
					r = n
					break
				}
			}
		}
		R = append(R, r)
	}

	return R
}
