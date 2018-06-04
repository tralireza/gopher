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
