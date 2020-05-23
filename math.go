package gopher

import (
	"log"
	"math"
	"math/bits"
	"slices"
	"strconv"
	"strings"
	"time"
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

// 335h Self Crossing
func isSelfCrossing(distance []int) bool {
	D := distance
	for i := 3; i < len(D); i++ {
		if D[i] >= D[i-2] && D[i-3] >= D[i-1] {
			return true
		}
	}

	for i := 4; i < len(D); i++ {
		if D[i]+D[i-4] >= D[i-2] && D[i-3] == D[i-1] {
			return true
		}
	}

	for i := 5; i < len(D); i++ {
		if D[i]+D[i-4] >= D[i-2] && D[i-5]+D[i-1] >= D[i-3] && D[i-2] >= D[i-4] && D[i-3] >= D[i-1] {
			return true
		}
	}

	return false
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

// 587h Erect the Fence
func outerTrees(trees [][]int) [][]int {
	if len(trees) <= 3 {
		return trees
	}

	CrossPrd := func(o, a, b []int) int {
		return (a[0]-o[0])*(b[1]-o[1]) - (a[1]-o[1])*(b[0]-o[0])
	}

	slices.SortFunc(trees, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	Lower := [][]int{}
	for _, p := range trees {
		for len(Lower) > 1 && CrossPrd(Lower[len(Lower)-2], Lower[len(Lower)-1], p) < 0 {
			Lower = Lower[:len(Lower)-1]
		}
		Lower = append(Lower, p)
	}
	log.Print("-> Lower: ", Lower)

	slices.Reverse(trees)
	Upper := [][]int{}
	for _, p := range trees {
		for len(Upper) > 1 && CrossPrd(Upper[len(Upper)-2], Upper[len(Upper)-1], p) < 0 {
			Upper = Upper[:len(Upper)-1]
		}
		Upper = append(Upper, p)
	}
	log.Print("-> Upper: ", Upper)

	ConvexHull := make([][]int, len(Upper)-1+len(Lower)-1)
	copy(ConvexHull, Lower[:len(Lower)-1])
	copy(ConvexHull[len(Lower)-1:], Upper[:len(Upper)-1])

	Compact := make([][]int, len(ConvexHull))
	copy(Compact, ConvexHull)

	points, M := 0, map[[2]int]struct{}{}
	for _, p := range Compact {
		if _, ok := M[[2]int{p[0], p[1]}]; !ok {
			M[[2]int{p[0], p[1]}] = struct{}{}
			ConvexHull[points] = p
			points++
		}
	}

	log.Print(":: Convex Hull: ", ConvexHull[:points])
	return ConvexHull[:points]
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

// 762 Prime Number of Set Bits in Binary Representation
func countPrimeSetBits(left, right int) int {
	count := 0
	for x := left; x <= right; x++ {
		bits := 0
		for n := x; n > 0; n >>= 1 {
			bits += n & 1
		}

		for _, p := range []int{2, 3, 5, 7, 11, 13, 17, 19} {
			if p == bits {
				count++
			}
		}
	}

	return count
}

// 780h Reaching Point
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// Euclidean Algorithm: https://en.wikipedia.org/wiki/Euclidean_algorithm

	var Search func(x, y, tx, ty int) bool
	Search = func(x, y, tx, ty int) bool {
		if x > tx {
			return false
		}

		if x == tx {
			return ((ty - y) % x) == 0
		}

		return Search(y, x, ty%tx, tx)
	}

	if sx > tx || sy > ty {
		return false
	}

	if tx < ty {
		return Search(sx, sy, tx, ty)
	}
	return Search(sy, sx, ty, tx)
}

// 812 Largest Triangle Area
func largestTriangleArea(points [][]int) float64 {
	xArea := 0
	for i, a := range points[:len(points)-2] {
		for j, b := range points[i+1 : len(points)-1] {
			for _, c := range points[j+1:] {
				area := a[0]*b[1] + b[0]*c[1] + c[0]*a[1] - a[1]*b[0] - b[1]*c[0] - c[1]*a[0]
				if area < 0 {
					area *= -1
				}

				xArea = max(area, xArea)
			}
		}
	}

	return float64(xArea) / float64(2)
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

// 883 Projection Area of 3D Shapes
func projectionArea(grid [][]int) int {
	tArea := 0

	for r := range grid {
		xRow, xCol := 0, 0
		for c := range grid[r] {
			if grid[r][c] > 0 {
				tArea++
			}

			xRow, xCol = max(xRow, grid[r][c]), max(xCol, grid[c][r])
		}
		tArea += xRow + xCol
	}

	return tArea
}

// 892 Surface Area of 3D Shapes
func surfaceArea(grid [][]int) int {
	N := len(grid)
	Dirs := [5]int{-1, 0, 1, 0, -1}

	tArea := 0
	for r := range N {
		for c := range N {
			if grid[r][c] > 0 {
				tArea += 2 // Top & Bottom of Cube
			}

			area := grid[r][c]
			for d := range 4 { // Sides of Cube
				r, c := r+Dirs[d], c+Dirs[d+1]
				if 0 <= r && r < N && 0 <= c && c < N {
					tArea += max(area-grid[r][c], 0)
				} else {
					tArea += area
				}
			}
		}
	}

	return tArea
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

// 1175 Prime Arrangements
func numPrimeArrangements(n int) int {
	P := make([]int, n+1)
	for p := range P {
		P[p] = p
	}

	for p := 2; p < len(P); p++ {
		if P[p] == p {
			for m := p * p; m < len(P); m += p {
				P[m] = p
			}
		}
	}
	log.Print("-> Sieve ", P)

	primes := 0
	for p := 2; p < len(P); p++ {
		if P[p] == p {
			primes++
		}
	}

	const M = int(1e9) + 7
	factMod := func(n int) int {
		f := 1
		for n > 1 {
			f = f * n % M
			n--
		}
		return f
	}

	return factMod(n-primes) * factMod(primes) % M
}

// 1295 Find Numbers with Even Number of Digits
func findNumbers(nums []int) int {
	evens := 0
	for _, n := range nums {
		t := 0
		for n > 0 {
			n /= 10
			t++
		}

		evens += 1 ^ (t & 1)
	}

	return evens
}

// 1432m Max Difference You Can Get From Changing an Integer
func maxDiff(num int) int {
	darr := []byte(strconv.Itoa(num))

	vMax, vMin := num, 0
	for i := 0; i < len(darr); i++ {
		if darr[i] == '9' {
			continue
		}

		vMax = 0
		for j := 0; j < len(darr); j++ {
			vMax *= 10
			switch darr[j] {
			case darr[i]:
				vMax += 9
			default:
				vMax += int(darr[j] - '0')
			}
		}
		break
	}

	switch darr[0] {
	case '1':
		vMin = num
		for i := 1; i < len(darr); i++ {
			if darr[i] == '0' || darr[i] == '1' {
				continue
			}

			vMin = 1
			for j := 1; j < len(darr); j++ {
				vMin *= 10
				if darr[j] != darr[i] {
					vMin += int(darr[j] - '0')
				}
			}
			break
		}
	default:
		for j := 0; j < len(darr); j++ {
			vMin *= 10
			if darr[j] == darr[0] {
				vMin += 1
			} else {
				vMin += int(darr[j] - '0')
			}
		}
	}

	return vMax - vMin
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

// 1886 Determine Whether Matrix Can Be Obtained By Rotation
func findRotation(mat [][]int, target [][]int) bool {
	N := len(mat)

NEXT:
	for range 4 {
		// 90 Rotation: Transpose M + Row Reverse
		for r := 0; r < N; r++ {
			for c := r; c < N; c++ {
				mat[r][c], mat[c][r] = mat[c][r], mat[r][c]
			}
		}
		for r := range mat {
			if r == 0 {
				log.Print("-> T ", mat[r])
			} else {
				log.Print("     ", mat[r])
			}
		}

		for r := 0; r < N; r++ {
			for c := 0; c < N/2; c++ {
				mat[r][c], mat[r][N-c-1] = mat[r][N-c-1], mat[r][c]
			}
		}
		for r := range mat {
			if r == 0 {
				log.Print("-> R ", mat[r])
			} else {
				log.Print("     ", mat[r])
			}
		}

		for r := 0; r < N; r++ {
			for c := 0; c < N; c++ {
				if mat[r][c] != target[r][c] {
					continue NEXT
				}
			}
		}

		return true
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

// 2081h Sum of k-Mirror Numbers
func kMirror(k int, n int) int64 {
	kmSum := int64(0)

	kPalindrome := func(n, k int64) bool {
		p, t := int64(0), n
		for t > 0 {
			p *= k
			p += t % k
			t /= k
		}
		return p == n
	}

	for start := int64(1); ; start *= 10 {
		for evenOdd := range 2 {
			for v := start; v < start*10; v++ {
				p, t := v, v

				switch evenOdd {
				case 0:
					t /= 10
				}

				for t > 0 {
					p *= 10
					p += t % 10
					t /= 10
				}

				if kPalindrome(p, int64(k)) {
					log.Printf("-> [%d]  %d   %v", n, p, kPalindrome(p, int64(k)))

					n--
					kmSum += p
					if n == 0 {
						return kmSum
					}
				}
			}
		}
	}

	return -1
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

// 2566 Maximum Difference by Remapping a Digit
func minMaxDifference(num int) int {
	s := strconv.Itoa(num)

	for i := 0; i < len(s); i++ {
		if s[i] == '9' {
			continue
		}

		vMax, vMin := 0, 0
		for j := 0; j < len(s); j++ {
			vMax *= 10
			if s[j] != s[i] {
				vMax += int(s[j] - '0')
			} else {
				vMax += 9
			}

			vMin *= 10
			if s[j] != s[0] {
				vMin += int(s[j] - '0')
			}
		}

		return vMax - vMin
	}

	return num
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

// 2894 Divisible and Non-divisible Sums Difference
func differenceOfSums(n int, m int) int {
	nSum := 0
	for v := range n {
		if (v+1)%m != 0 {
			nSum += v + 1
		}
	}
	return 2*nSum - n*(n+1)/2
}

// 2929m Distribute Candies Among Children II
// 1 <= N, L <= 10^6
func distributeCandiesII(n int, limit int) int64 {
	ways := int64(0)
	for candy1 := 0; candy1 <= min(n, limit); candy1++ {
		if n-candy1 <= 2*limit {
			maxCandy2 := min(n-candy1, limit)
			minCandy2 := max(0, n-candy1-limit)

			ways += int64(maxCandy2 - minCandy2 + 1)

			log.Printf("-> Candy1: %d | Candy2: %d ~ %d   => Ways: (+%d) %d",
				candy1, minCandy2, maxCandy2,
				maxCandy2-minCandy2+1,
				ways)
		}
	}

	Recursive := func() {
		M := make([][3]int64, n+1)
		for r := range M {
			M[r][0], M[r][1], M[r][2] = -1, -1, -1
		}

		rCalls, mHits := 0, 0
		var Search func(child, leftCandies int) int64
		Search = func(child, leftCandies int) int64 {
			rCalls++
			if child == 3 {
				if leftCandies == 0 {
					return 1
				}
				return 0
			}

			if M[leftCandies][child] != -1 {
				mHits++
				return M[leftCandies][child]
			}

			ways := int64(0)
			for candy := 0; candy <= min(limit, leftCandies); candy++ {
				if leftCandies-candy <= (2-child)*limit {
					ways += Search(child+1, leftCandies-candy)
				}
			}

			M[leftCandies][child] = ways
			return ways
		}
		tStart := time.Now()
		log.Printf(":: DFS Search -> %d   [@ %v] Calls: %d | Hits: %d", Search(0, n), time.Since(tStart), rCalls, mHits)
	}

	DP := func() {
		tStart := time.Now()
		D := make([][4]int64, n+1)
		D[0][3] = 1

		for leftCandies := 0; leftCandies <= n; leftCandies++ {
			for child := 2; child >= 0; child-- {
				ways := int64(0)
				for candy := 0; candy <= min(leftCandies, limit); candy++ {
					ways += D[leftCandies-candy][child+1]
				}

				D[leftCandies][child] = ways
			}
		}

		log.Printf(":: DP -> %d   [@ %v]", D[n][0], time.Since(tStart))
	}

	Recursive()
	DP()

	// Multiset C(n, k) = C(n+k-1, k-1)
	Combinatorics := func() {
		tStart := time.Now()

		Choose_3_1, Choose_3_2, Choose_3_3 := int64(3), int64(3), int64(1)
		Choose_n_2 := func(n int) int64 {
			if n < 0 {
				return int64(0)
			}
			return int64(n) * int64(n-1) / 2
		}

		ways := Choose_n_2(n+2) -
			Choose_3_1*Choose_n_2(n-(limit+1)+2) +
			Choose_3_2*Choose_n_2(n-2*(limit+1)+2) -
			Choose_3_3*Choose_n_2(n-3*(limit+1)+2)

		log.Printf(":: Combinatorics -> %d   [@ %v]", ways, time.Since(tStart))
	}
	Combinatorics()

	return ways
}

// 3024 Type of Triangle
func triangleType(nums []int) string {
	slices.Sort(nums)

	a, b, c := nums[0], nums[1], nums[2]
	if a+b <= c {
		return "none"
	}

	if a == c {
		return "equilateral"
	}
	if a == b || b == c {
		return "isosceles"
	}
	return "scalene"
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

// 3307h Find the K-th Character in String Game II
func kthCharacterII(k int64, operations []int) byte {
	BitMath := func(k int64, operations []int) byte {
		offset := 0
		k--
		for p := bits.Len64(uint64(k)) - 1; p >= 0; p-- {
			if (k>>p)&1 == 1 {
				offset += operations[p]
			}
		}

		return 'a' + byte(offset%26)
	}
	log.Printf(":? Bit Math: %q", BitMath(k, operations))

	offset := 0
	for k != 1 {
		t := bits.Len64(uint64(k)) - 1
		if 1<<t == k {
			t--
		}
		k -= 1 << t
		if operations[t] == 1 {
			offset++
		}
	}

	return 'a' + byte(offset%26)
}

// 3312h Sorted GCD Pair Queries
func gcdValues(nums []int, queries []int64) []int {
	xVal := slices.Max(nums)
	freq := make([]int, xVal+1)
	for _, n := range nums {
		freq[n]++
	}

	log.Print("-> frequency :: ", freq)

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

	log.Print("-> GCD[g] :: ", GCD)

	pSum := make([]int64, xVal+1)
	for g := 1; g <= xVal; g++ {
		pSum[g] = pSum[g-1] + int64(GCD[g])
	}

	log.Print("-> Sigma GCD[g] :: ", pSum)

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

// 3405h Count the Number of Arrays with K Matching Adjacent Elements
func countGoodArrays(n, m, k int) int {
	const M = 1e9 + 7

	mPower := func(b, e int) int {
		power := 1
		for e > 0 {
			if e&1 == 1 {
				power = power * b % M
			}
			b = b * b % M
			e >>= 1
		}
		return power
	}

	Facts, iFacts := make([]int, n), make([]int, n)

	Facts[0] = 1
	for i := 1; i < n; i++ {
		Facts[i] = i * Facts[i-1] % M
	}
	log.Print("-> Facts: ", Facts)

	iFacts[n-1] = mPower(Facts[n-1], M-2)
	for i := n - 1; i > 0; i-- {
		iFacts[i-1] = i * iFacts[i] % M
	}
	log.Print("-> Inv. Facts: ", iFacts)

	nCk := func(n, k int) int {
		return Facts[n] * iFacts[k] % M * iFacts[n-k] % M
	}

	return m * nCk(n-1, k) % M * mPower(m-1, n-k-1) % M
}

// 3443m Maximum Manhattan Distance After K Changes
func maxDistance(s string, k int) int {
	xDist := 0

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	lat, long := 0, 0

	for i, dir := range s {
		switch dir {
		case 'N':
			lat++
		case 'S':
			lat--
		case 'W':
			long++
		case 'E':
			long--
		}

		xDist = max(min(abs(lat)+abs(long)+2*k, i+1), xDist)
	}

	return xDist
}
