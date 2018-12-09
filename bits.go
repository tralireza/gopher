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
