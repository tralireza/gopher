// Array
package gopher

import (
	"bytes"
	"log"
	"slices"
)

// 73m Set Matrix Zeroes
// -2^31 <= M_ij <= 2^31-1
// 1 <= Rows, Cols <= 200
func setZeroes(matrix [][]int) {
	Rows, Cols := len(matrix), len(matrix[0])

	rowZero, colZero := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			rowZero = true
		}
	}
	for _, row := range matrix {
		if row[0] == 0 {
			colZero = true
		}
	}

	for r := 1; r < Rows; r++ {
		for c := 1; c < Cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c], matrix[r][0] = 0, 0
			}
		}
	}

	for r := 1; r < Rows; r++ {
		for c := 1; c < Cols; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if rowZero {
		for c := range Cols {
			matrix[0][c] = 0
		}
	}
	if colZero {
		for r := range Rows {
			matrix[r][0] = 0
		}
	}

	log.Print(":: ", matrix)
}

// 485 Max Consecutive Ones
func findMaxConsecutiveOnes(nums []int) int {
	tCur, tMax := 0, 0
	for _, n := range nums {
		if tCur+n > tCur {
			tCur++
		} else {
			tCur = 0
		}

		tMax = max(tCur, tMax)
	}

	return tMax
}

// 661 Image Smoother
func imageSmoother(img [][]int) [][]int {
	imgSm := make([][]int, len(img))

	for r := 0; r < len(img); r++ {
		for c := 0; c < len(img[r]); c++ {
			v, n := 0, 0

			for x := r - 1; x <= r+1; x++ {
				for y := c - 1; y <= c+1; y++ {
					if 0 <= x && x < len(img) && 0 <= y && y < len(img[r]) {
						v += img[x][y]
						n++
					}
				}
			}

			imgSm[r] = append(imgSm[r], v/n)
		}
	}

	log.Print(":? ", imgSm)
	return imgSm
}

// 766 Toeplitz Matrix
func isToeplitzMatrix(matrix [][]int) bool {
	for r := len(matrix) - 1; r > 0; r-- {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][c] != matrix[r-1][c-1] {
				return false
			}
		}
	}

	return true
}

// 798h Smallest Rotation with Highest Score
func bestRotation(nums []int) int {
	// N_i <= i -> +1 score
	N := len(nums)

	Scores := make([]int, N)
	for i, n := range nums {
		// k: Shift interval [left..right]
		leftK, rightK := (N+i-n+1)%N, (i+1)%N
		Scores[leftK]--
		Scores[rightK]++
		if leftK > rightK {
			Scores[0]--
		}
	}

	best, cur := -N, 0
	x := 0
	for i, score := range Scores {
		cur += score
		if cur > best {
			best = cur
			x = i
		}
	}

	return x
}

// 821 Shortest Distance to a Character
func shortestToChar(s string, c byte) []int {
	D := make([]int, len(s))
	for i := range D {
		D[i] = len(s)
	}

	if s[0] == c {
		D[0] = 0
	}

	for i := 1; i < len(s); i++ {
		switch s[i] {
		case c:
			D[i] = 0
		default:
			D[i] = D[i-1] + 1
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		D[i] = min(D[i+1]+1, D[i])
	}

	return D
}

// 1394 Find Lucky Integer in an Array
func findLucky(arr []int) int {
	F := [500 + 1]int{}
	for _, n := range arr {
		F[n]++
	}

	for n := 500; n > 0; n-- {
		if n == F[n] {
			return n
		}
	}
	return -1
}

// 1437 Check If All 1's Are at Least Length K Places Away
func kLengthApart(nums []int, k int) bool {
	dist := k
	for _, n := range nums {
		switch n {
		case 1:
			if dist < k {
				return false
			}
			dist = 0
		case 0:
			dist++
		}
	}

	return true
}

// 1534 Count Good Triplets
func countGoodTriplets(arr []int, a, b, c int) int {
	// 0 <= A_i <= 1000

	Abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	// O(N^2 + k*N)
	Optimized := func() int {
		count := 0
		pSum := make([]int, 1000+1)

		// O(k*N) -> O(N*logN)
		BIT := make([]int, 1000+1 /* Shift for N_i == 0 :: 0 -> 1 */ +1 /* BIT: Root */) // BIT: Binary Indexed Tree
		Update := func(i int) {
			for i <= 1001 {
				BIT[i]++
				i += i & (-i)
			}
		}
		Query := func(i int) int {
			r := 0
			for i > 0 {
				r += BIT[i]
				i -= i & (-i)
			}
			return r
		}

		// Intervals:
		// [arr[j] - a ... arr[j] + a]  [arr[k] - c ... arr[k] + c]
		for j := 0; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if Abs(arr[j]-arr[k]) <= b {
					left := max(0, max(arr[j]-a, arr[k]-c))
					right := min(1000, min(arr[j]+a, arr[k]+c))

					// Count all arr[i] that are in: [left, right]
					if left <= right {
						if left == 0 {
							count += pSum[right]
						} else {
							count += pSum[right] - pSum[left-1]
						}

						// O(logN)
						log.Printf("%d +%d", count, Query(right+1)-Query(left))
					}
				}
			}

			for v := arr[j]; v <= 1000; v++ {
				pSum[v]++
			}

			// O(logN)
			Update(arr[j] + 1)
		}

		log.Print("-> BIT: ", BIT)

		return count
	}

	count := 0
	for i, x := range arr[:len(arr)-2] {
		for j, y := range arr[i+1 : len(arr)-1] {
			if Abs(x-y) <= a {
				for _, z := range arr[i+1+j+1:] {
					if Abs(y-z) <= b && Abs(z-x) <= c {
						count++
					}
				}
			}
		}
	}

	log.Printf(":: %d ~ %d", count, Optimized())

	return count
}

// 1550 Three Consecutive Odds
func threeConsecutiveOdds(arr []int) bool {
	counter := 0
	for _, n := range arr {
		if n&1 == 1 {
			counter++
			if counter == 3 {
				return true
			}
		} else {
			counter = 0
		}
	}

	return false
}

// 1752 Check If Array Is Sorted and Rotated
func check(nums []int) bool {
	inversions := 0
	if nums[0] < nums[len(nums)-1] {
		inversions++
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			inversions++
		}
	}
	return inversions <= 1
}

// 1920 Build Array from Permutation
func buildArray(nums []int) []int {
	log.Print("** ", nums)

	InPlace := func(nums []int) []int {
		for i := range nums {
			nums[i] += 1000 * (nums[nums[i]] % 1000)
		}
		for i := range nums {
			nums[i] /= 1000
		}

		return nums
	}

	R := make([]int, len(nums))
	for i, n := range nums {
		R[i] = nums[n]
	}

	log.Print(":: ", InPlace(nums))
	return R
}

// 2200 Find All K-Distant Indices in an Array
func findKDistantIndices(nums []int, key int, k int) []int {
	kDists := []int{}

	l := 0
	for r, n := range nums {
		if n == key {
			for i := max(r-k, l); i <= min(r+k, len(nums)-1); i++ {
				kDists = append(kDists, i)
			}
			l = r + k + 1
		}
	}

	log.Print(":: ", kDists)
	return kDists
}

// 2033m Minimum Operations to Make a Uni-Value Grid
func minOperations_UniValue(grid [][]int, x int) int {
	nums := make([]int, 0, len(grid)*len(grid[0]))
	for r := range grid {
		nums = append(nums, grid[r]...)
	}

	slices.Sort(nums)
	median := nums[len(nums)/2]

	log.Print("-> ", nums, median)

	Abs := func(v int) int {
		if v >= 0 {
			return v
		}
		return -v
	}

	ops, r := 0, median%x
	for _, n := range nums {
		if n%x != r {
			return -1
		}

		ops += Abs(n-median) / x
	}

	return ops
}

// 2094 Finding 3-Digit Even Numbers
func findEvenNumbers(digits []int) []int {
	F := [10]int{}
	for _, d := range digits {
		F[d]++
	}

	R := []int{}
	for h := 1; h <= 9; h++ {
		if F[h] == 0 {
			continue
		}
		F[h]--
		for t := 0; t <= 9; t++ {
			if F[t] == 0 {
				continue
			}
			F[t]--
			for o := 0; o <= 8; o += 2 {
				if F[o] > 0 {
					R = append(R, 100*h+10*t+o)
				}
			}
			F[t]++
		}
		F[h]++
	}

	return R
}

// 2145m Count the Hidden Sequences
func numberOfArrays(differences []int, lower, upper int) int {
	// 1 <= N <= 10^5, -10^5 <= N_i <= 10^5
	S := make([]int64, 0, len(differences)+1)
	S = append(S, 0)
	for _, d := range differences {
		S = append(S, S[len(S)-1]+int64(d))
	}

	log.Print("-> ", S)

	x, n := slices.Max(S), slices.Min(S)
	if int(x-n) > upper-lower {
		return 0
	}
	return upper - lower - int(x-n) + 1
}

// 2176 Count Equal and Divisible Pairs in an Array
func countPairs_Divisible(nums []int, k int) int {
	count := 0
	for i, a := range nums[:len(nums)-1] {
		for j, b := range nums[i+1:] {
			if a == b && i*(j+i+1)%k == 0 {
				count++
			}
		}
	}

	return count
}

// 2302h Count Subarrays With Score Less Than K
func countSubarrays_KScore(nums []int, k int64) int64 {
	count := int64(0)

	l, psum := 0, int64(0)
	for r, n := range nums {
		psum += int64(n)
		for psum*int64(r-l+1) >= k {
			psum -= int64(nums[l])
			l++
		}
		count += int64(r - l + 1)
	}

	return count
}

// 2780m Minimum Index of a Valid Split
func minimumIndex(nums []int) int {
	F := map[int]int{}
	for _, n := range nums {
		F[n]++
	}

	dominant, frq := 0, 0
	for n, f := range F {
		if f > frq {
			dominant, frq = n, f
		}
	}

	log.Print("-> ", dominant, frq)

	f := 0
	for i, n := range nums {
		if n == dominant {
			f++
		}

		if f*2 > (i+1) && (frq-f)*2 > len(nums)-1-i {
			return i
		}
	}

	return -1
}

// 2873 Maximum Value of an Ordered Triplet I
func maximumTripletValue(nums []int) int64 {
	xVal := int64(0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				xVal = max(xVal, int64(nums[i]-nums[j])*int64(nums[k]))
			}
		}
	}

	return xVal
}

// 2874m Maximum Value of an Ordered Triplet II
func maximumTripletValueII(nums []int) int64 {
	xVal := int64(0)
	lMax, diffMax := 0, 0
	for _, n := range nums {
		xVal = max(int64(diffMax)*int64(n), xVal)

		diffMax = max(lMax-n, diffMax)
		lMax = max(n, lMax)
	}

	return xVal
}

// 2942 Find Words Containing Character
func findWordsContaining(words []string, x byte) []int {
	R := []int{}
	for i := range words {
		if bytes.IndexByte([]byte(words[i]), x) != -1 {
			R = append(R, i)
		}
	}

	return R
}

// 3169m Count Days Without Meetings
func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(x, y []int) int {
		if x[0] == y[0] {
			return x[1] - y[1]
		}
		return x[0] - y[0]
	})

	log.Print("-> ", meetings)

	t := 0

	lDay := 0
	for _, meeting := range meetings {
		start, finish := meeting[0], meeting[1]
		if start > lDay {
			t += start - lDay - 1
		}

		lDay = max(lDay, finish)
	}
	t += days - lDay

	return t
}

// 3355m Zero Array Transformation I
func isZeroArray(nums []int, queries [][]int) bool {
	Delta := make([]int, len(nums)+1)

	for _, query := range queries {
		Delta[query[0]]++
		Delta[query[1]+1]--
	}

	pSum, v := []int{}, 0
	for i := range len(Delta) {
		v += Delta[i]
		pSum = append(pSum, v)
	}

	log.Print("-> ", pSum)

	for i, num := range nums {
		if num > pSum[i] {
			return false
		}
	}
	return true
}

// 3392 Count Subarrays of Length Three With a Condition
func countSubarrays_Length3(nums []int) int {
	count := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == 2*(nums[i-1]+nums[i+1]) {
			count++
		}
	}

	return count
}

// 3394m Check if Grid can be Cut into Sections
func checkValidCuts(n int, rectangles [][]int) bool {
	Check := func(offset int) bool {
		slices.SortFunc(rectangles, func(x, y []int) int { return x[offset] - y[offset] })

		gaps, end := 0, rectangles[0][offset+2]
		for _, rectangle := range rectangles[1:] {
			if end <= rectangle[offset] {
				gaps++
			}
			end = max(rectangle[offset+2], end)
		}

		return gaps >= 2
	}
	return Check(0) || Check(1)
}

// 3423 Maximum Difference Between Adjacent Elements in a Circular Array
func maxAdjacentDistance(nums []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	xDiff := abs(nums[0] - nums[len(nums)-1])
	for i, v := range nums[:len(nums)-1] {
		xDiff = max(abs(v-nums[i+1]), xDiff)
	}
	return xDiff
}
