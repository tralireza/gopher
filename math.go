package gopher

import (
	"log"
	"slices"
)

// 319m Bulb Switcher
func bulbSwitch(n int) int {
	q := 0 // Square Root
	for q*q <= n {
		q++
	}
	return q - 1
}

// 3312h Sorted GCD Pair Queries
func gcdValues(nums []int, queries []int64) []int {
	xVal := slices.Max(nums)
	freq := make([]int, xVal+1)
	for _, n := range nums {
		freq[n]++
	}

	log.Print(" -> frequency :: ", freq)

	GCD := make([]int, xVal+1) // count of Pairs: {Ni, Nj} with gcd(Ni, Nj) = GCD[g]
	for g := xVal; g >= 1; g-- {
		count := 0
		for m := g; m <= xVal; m += g { // multiples of g
			count += freq[m]
		}
		GCD[g] = count * (count - 1) / 2 // nC2 ~ n!/2!.(n-2)!

		for m := 2 * g; m <= xVal; m += g { // remove double-counted
			GCD[g] -= GCD[m]
		}
	}

	log.Print(" -> GCD[g] :: ", GCD)

	pSum := make([]int64, xVal+1)
	for g := 1; g <= xVal; g++ {
		pSum[g] = pSum[g-1] + int64(GCD[g])
	}

	log.Print(" -> Sigma GCD[g] :: ", pSum)

	R := []int{}
	for _, q := range queries {
		l, r := 0, xVal+1
		for l < r {
			m := l + (r-l)>>1
			if pSum[m] > q {
				r = m
			} else {
				l = m + 1
			}
		}
		R = append(R, r)
	}
	return R
}
