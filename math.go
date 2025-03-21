package gopher

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

// 319m Bulb Switcher
func bulbSwitch(n int) int {
	q := 0 // Square Root
	for q*q <= n {
		q++
	}
	return q - 1
}

// 326 Power of Three
func isPowerOfThree(n int) bool {
	p := 1
	for p < n {
		p *= 3
	}
	return p == n
}

// 342 Power of Four
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

// 598 Range Addition II
func maxCount(m int, n int, ops [][]int) int {
	x, y := m, n
	for _, o := range ops {
		x = min(o[0], x)
		y = min(o[1], y)
	}

	return x * y
}

// 908 Smallest Range I
func smallestRangeI(nums []int, k int) int {
	return max(0, slices.Max(nums)-slices.Min(nums)-2*k)
}

// 989 Add to Array-Form of Integer
func addToArrayForm(num []int, k int) []int {
	R := []int{}

	carry, p := k, len(num)
	for carry > 0 || p > 0 {
		if p > 0 {
			p--
			carry += num[p]
		}
		R = append(R, carry%10)
		carry /= 10
	}

	slices.Reverse(R)
	return R
}

// 1154 Day of the Year
func dayOfYear(date string) int {
	Days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	D := [3]int{}
	for i, sval := range strings.Split(date, "-") {
		D[i], _ = strconv.Atoi(sval)
	}

	y, m, d := D[0], D[1], D[2]
	if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
		Days[1] += 1
	}

	dy := d
	for i := 0; i < m-1; i++ {
		dy += Days[i]
	}

	return dy
}

// 1780m Check if Number is a Sum of Powers of Three
func checkPowersOfThree(n int) bool {
	P := []int{}
	power := 1
	for power <= 10000_000 {
		P = append(P, power)
		power *= 3
	}
	slices.Reverse(P)

	log.Print("-> ", P)

	for _, power := range P {
		if n >= power {
			n -= power
			if n == 0 {
				return true
			}
		}
	}
	return false
}

// 1998h GCD Sort of an Array
func gcdSort(nums []int) bool {
	xVal := slices.Max(nums)

	factors := make([][]int, xVal+1)
	for p := 2; p <= xVal; p++ {
		if len(factors[p]) > 0 { // p :: not a Prime
			continue
		}
		for m := p; m <= xVal; m += p {
			factors[m] = append(factors[m], p)
		}
	}

	log.Print(" -> factors :: ", factors)

	G := map[int][]int{} // Virtual Graph of all N and their factors
	for _, n := range nums {
		for _, f := range factors[n] {
			G[n], G[f] = append(G[n], f), append(G[f], n)
		}
	}

	log.Print(" -> Graph :: ", G)

	Ranks := make([]int, xVal+1)
	DJS := make([]int, xVal+1)
	for n := range DJS {
		DJS[n] = n
	}

	var FindSet func(int) int
	FindSet = func(x int) int {
		if DJS[x] != x {
			DJS[x] = FindSet(DJS[x])
		}
		return DJS[x]
	}

	Union := func(x, y int) {
		x, y = FindSet(x), FindSet(y)
		if x == y {
			return
		}
		if Ranks[y] > Ranks[x] {
			DJS[x] = y
		} else {
			if Ranks[x] == Ranks[y] {
				Ranks[x]++
			}
			DJS[y] = x
		}
	}

	for v := range G {
		for _, u := range G[v] {
			Union(v, u)
		}
	}

	log.Print(" -> DJS :: ", DJS)

	sorted := make([]int, len(nums))
	copy(sorted, nums)
	slices.Sort(sorted)

	for i := range nums {
		if nums[i] != sorted[i] {
			if FindSet(nums[i]) != FindSet(sorted[i]) {
				return false
			}
		}
	}

	return true
}

// 2523m Closest Prime Numbers in Range
func closestPrimes(left int, right int) []int {
	Sieve := make([]int, right+1)
	for n := range Sieve {
		Sieve[n] = n
	}
	Sieve[1] = 0

	for p := 2; p <= right; p++ {
		if Sieve[p] == p {
			for m := p * p; m <= right; m += p {
				Sieve[m] = p
			}
		}
	}

	log.Print("-> ", Sieve)

	Prime := []int{}
	for p := range Sieve {
		if Sieve[p] == p && p >= left {
			Prime = append(Prime, p)
		}
	}

	log.Print("-> ", Prime)

	if len(Prime) < 2 {
		return []int{-1, -1}
	}

	R := []int{Prime[0], Prime[1]}
	for i := range Prime[:len(Prime)-1] {
		if Prime[i+1]-Prime[i] < R[1]-R[0] {
			R[0], R[1] = Prime[i], Prime[i+1]
		}
	}
	return R
}

// 2579m Count Total Number of Colored Cells
func coloredCells(n int) int64 {
	cells := int64(1)
	for i := 2; i <= n; i++ {
		cells += int64(4 * (n - 1))
	}

	log.Print("-> ", 1+2*n*(n-1))

	return cells
}

// 3151 Special Array I
func isArraySpecialI(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if (nums[i-1]^nums[i])&1 == 0 {
			return false
		}
	}

	return true
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
