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

// 838m Push Dominoes
func pushDominoes(dominoes string) string {
	log.Print(":> ", strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(dominoes, "R.L", "R|L"), ".L", "LL"), "R.", "RR"))

	N := len(dominoes)
	F := make([]int, N)
	var force int

	force = 0
	for i := 0; i < N; i++ {
		switch dominoes[i] {
		case 'R':
			force = N
		case 'L':
			force = 0
		default:
			force = max(force-1, 0)
		}

		F[i] += force
	}

	force = 0
	for i := N - 1; i >= 0; i-- {
		switch dominoes[i] {
		case 'L':
			force = N
		case 'R':
			force = 0
		default:
			force = max(force-1, 0)
		}

		F[i] -= force
	}

	log.Print("-> ", F)

	sb := strings.Builder{}
	for _, f := range F {
		if f > 0 {
			sb.WriteRune('R')
		} else if f < 0 {
			sb.WriteRune('L')
		} else {
			sb.WriteRune('.')
		}
	}

	log.Print(":: ", sb.String())

	return sb.String()
}

// 908 Smallest Range I
func smallestRangeI(nums []int, k int) int {
	return max(0, slices.Max(nums)-slices.Min(nums)-2*k)
}

// 970m Powerful Integers
func powerfulIntegers(x int, y int, bound int) []int {
	Px, Py := []int{1}, []int{1}
	if x > 1 {
		for power := x; power < bound; power *= x {
			Px = append(Px, power)
		}
	}
	if y > 1 {
		for power := y; power < bound; power *= y {
			Py = append(Py, power)
		}
	}

	log.Printf("-> Px: %v | Py: %v ", Px, Py)

	Set := map[int]struct{}{}
	for _, px := range Px {
		for _, py := range Py {
			if px+py <= bound {
				Set[px+py] = struct{}{}
			}
		}
	}

	R := []int{}
	for power := range Set {
		R = append(R, power)
	}
	slices.Sort(R)
	return R
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

// 1922m Count Good Numbers
func countGoodNumbers(n int64) int {
	const M = 1e9 + 7

	MPower := func(b, e int64) int64 {
		p := int64(1)
		for e > 0 {
			if e&1 == 1 {
				p = (p * b) % M
			}
			b = (b * b) % M
			e >>= 1
		}
		return p
	}

	return int(MPower(5, (n+1)/2) * MPower(4, n/2) % M)
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

// 2843 Count Symmetric Integers
func countSymmetricIntegers(low int, high int) int {
	count := 0

	for n := low; n <= high; n++ {
		s := strconv.Itoa(n)
		if len(s)&1 == 0 {
			l, r := 0, 0

			w := len(s)
			for i := 0; i < w/2; i++ {
				l += int(s[i] - '0')
				r += int(s[i+w/2] - '0')
			}

			if l == r {
				count++
			}
		}
	}

	// 1 <= N_i <= 10^4
	r := 0
	for n := low; n <= high; n++ {
		if n < 100 && n%11 == 0 {
			r++
		} else if 1000 <= n && n < 10000 {
			if n%10+(n%100)/10 == (n%1000)/100+n/1000 {
				r++
			}
		}
	}
	log.Print("-> ", r)

	return count
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

// 3272h Find the Count of Good Integers
func countGoodIntegers(n int, k int) int64 {
	start := 1
	for range (n - 1) / 2 {
		start *= 10
	}

	Set := map[string]struct{}{}

	for v := start; v < 10*start; v++ {
		left := strconv.Itoa(v)

		right := []byte(left)
		slices.Reverse(right)
		if n&1 == 1 {
			right = right[1:]
		}

		palindrome := left + string(right)

		v, _ := strconv.ParseInt(palindrome, 10, 64)
		if v%int64(k) == 0 {
			digits := []byte(palindrome)
			slices.Sort(digits)
			Set[string(digits)] = struct{}{}
		}
	}
	log.Print("-> ", Set)

	Fact := [10 + 1]int{}
	Fact[0], Fact[1] = 1, 1
	for n := 2; n <= 10; n++ {
		Fact[n] = Fact[n-1] * n
	}
	log.Print("-> ", math.MaxInt32, Fact, math.MaxInt64)

	count := int64(0)
	for digits := range Set {
		counter := [10]int{}
		for i := 0; i < len(digits); i++ {
			counter[digits[i]-'0']++
		}

		C_nk := int64(n-counter[0]) * int64(Fact[n-1])
		for _, count := range counter {
			C_nk /= int64(Fact[count])
		}
		count += C_nk
	}

	return count
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
