package gopher

import (
	"log"
	"math"
	"slices"
	"time"
)

// 11m Container With Most Water
func maxArea(height []int) int {
	A := 0

	l, r := 0, len(height)-1
	for l < r {
		L, R := height[l], height[r]
		area := min(L, R) * (r - l)
		if area > A {
			A = area
		}

		if L < R {
			l++
		} else {
			r--
		}
	}

	return A
}

// 15m 3Sum
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	R := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			v := nums[i] + nums[l] + nums[r]
			if v > 0 {
				r--
			} else if v < 0 {
				l++
			} else {
				R = append(R, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			}
		}
	}

	return R
}

// 36m Valid Sudoku
func isValidSudoku(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		M := make([]bool, 9+1) // rows
		for c := 0; c < 9; c++ {
			v := board[r][c]
			if v != '.' {
				if M[v-'0'] {
					return false
				}
				M[v-'0'] = true
			}
		}
	}

	for c := 0; c < 9; c++ {
		M := make([]bool, 9+1) // columns
		for r := 0; r < 9; r++ {
			v := board[r][c]
			if v != '.' {
				if M[v-'0'] {
					return false
				}
				M[v-'0'] = true
			}
		}
	}

	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			M := make([]bool, 9+1) // sub-boxes
			for x := range 3 {
				for y := range 3 {
					v := board[r+x][c+y]
					if v != '.' {
						if M[v-'0'] {
							return false
						}
						M[v-'0'] = true
					}
				}
			}
		}
	}

	return true
}

// 53m Maximum Subarray
func maxSubArray(nums []int) int {
	// Kadane's algorithm

	kX := nums[0]
	curX := nums[0]

	for _, n := range nums[1:] {
		curX = max(curX, 0) + n
		kX = max(kX, curX)
	}

	return kX
}

// 134m Gas Station
func canCompleteCircuit(gas []int, cost []int) int {
	p, tank, tankTotal := 0, 0, 0

	for i := range gas {
		tank += gas[i] - cost[i]
		tankTotal += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			p = i + 1
		}
	}

	if p == len(cost) || tankTotal < 0 {
		return -1
	}
	return p
}

// 135m Candy
func candy(ratings []int) int {
	C := make([]int, len(ratings))

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			C[i] = C[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			C[i] = max(C[i], C[i+1]+1)
		}
	}

	t := 0
	for _, c := range C {
		t += c
	}
	return t + len(ratings)
}

// 149h Max Points on a Line
func maxPointsOnLine(points [][]int) int {
	P := points
	if len(P) <= 2 {
		return len(P)
	}

	xP := 0

	for i := 0; i < len(P)-1; i++ {
		x, y := P[i][0], P[i][1]

		Mem := map[float64]int{}
		for j := i + 1; j < len(P); j++ {
			dx, dy := P[j][0]-x, P[j][1]-y

			if dx == 0 {
				Mem[math.MaxFloat64]++
			} else {
				Mem[float64(dy)/float64(dx)]++
			}
		}
		log.Print(P[i], " -> ", Mem)

		for _, v := range Mem {
			if v > xP {
				xP = v
			}
		}
	}

	return xP + 1
}

// 167m Two Sum II - Input Array Is Sorted
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		v := numbers[l] + numbers[r]
		if v == target {
			return []int{l + 1, r + 1}
		}

		if v < target {
			l++
		} else {
			r--
		}
	}

	return []int{0, 0}
}

// 670m Maximum Swap
func maximumSwap(num int) int {
	D := []int{}
	for n := num; n > 0; n /= 10 {
		D = append(D, n%10)
	}
	slices.Reverse(D)
	log.Print(" -> ", D)

	rMax := make([]int, len(D)) // RightMax
	rMax[len(D)-1] = len(D) - 1
	for i := len(D) - 2; i >= 0; i-- {
		rMax[i] = rMax[i+1]
		if D[i] > D[rMax[i+1]] {
			rMax[i] = i
		}
	}

	log.Print(" -> rMax (D_i) :: ", rMax)

	for i := 0; i < len(D)-1; i++ {
		if D[i] < D[rMax[i+1]] {
			D[i], D[rMax[i+1]] = D[rMax[i+1]], D[i]
			x := 0
			for _, d := range D {
				x = 10*x + d
			}
			return x
		}
	}

	return num
}

// 781m Rabbits in Forest
func numRabbits(answers []int) int {
	M := make([]int, 1000)
	for _, answer := range answers {
		M[answer]++
	}

	count := 0
	for n, frq := range M {
		count += (frq + n) / (n + 1) * (n + 1)
	}

	return count
}

// 918m Maximum Sum Circular Subarray
func maxSubarraySumCircular(nums []int) int {
	// Kadane's

	tSum := nums[0] // total

	kM, kX := nums[0], nums[0] // Kadane's Minimum | Maximum
	curM, curX := nums[0], nums[0]

	for _, n := range nums[1:] {
		curM, curX = min(curM, 0)+n, max(curX, 0)+n
		kM, kX = min(kM, curM), max(kX, curX)

		tSum += n
	}

	if kM == tSum {
		return kX
	}
	return max(kX, tSum-kM)
}

// 1007m Minimum Domino Rotations For Equal Row
func minDominoRotations(tops, bottoms []int) int {
	Greedy := func(tops, bottoms []int) int {
		r := math.MaxInt

	LOOP:
		for _, n := range []int{tops[0], bottoms[0]} {
			t, b := 0, 0
			for i := range tops {
				if tops[i] != n && bottoms[i] != n {
					continue LOOP
				}

				if tops[i] != n {
					t++
				}
				if bottoms[i] != n {
					b++
				}
			}
			r = min(min(t, b), r)
		}

		if r == math.MaxInt {
			return -1
		}
		return r
	}
	log.Print(":: ", Greedy(tops, bottoms))

	r := math.MaxInt

LOOP:
	for n := 1; n <= 6; n++ {
		t, b := 0, 0

		for i := range tops {
			if tops[i] != n && bottoms[i] != n {
				continue LOOP
			}

			if tops[i] != n {
				t++
			}
			if bottoms[i] != n {
				b++
			}
		}

		r = min(r, min(t, b))
	}

	if r == math.MaxInt {
		return -1
	}
	return r
}

// 1605m Find Valid Matrix Given Row and Column Sums
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	M := make([][]int, len(rowSum))
	for r := range M {
		M[r] = make([]int, len(colSum))
	}

	for r := 0; r < len(rowSum); r++ {
		for c := 0; c < len(colSum); c++ {
			mVal := rowSum[r]
			if colSum[c] < mVal {
				mVal = colSum[c]
			}

			M[r][c] = mVal

			rowSum[r] -= mVal
			colSum[c] -= mVal
		}
	}

	return M
}

// 2202m Maximize the Topmost Element After K Moves
func maximumTop(nums []int, k int) int {
	if len(nums) == 1 && k&1 == 1 {
		return -1
	}

	nX := -1
	for i := range min(len(nums), k-1) {
		nX = max(nums[i], nX)
	}

	if k < len(nums) {
		nX = max(nums[k], nX)
	}

	return nX
}

// 2280m Minimum Lines to Represent a Line Chart
func minimumLines(stockPrices [][]int) int {
	P := stockPrices
	if len(P) == 1 {
		return 0
	}

	slices.SortFunc(P, func(x, y []int) int { return x[0] - y[0] })

	t := 1
	Dx, Dy := 0, 0
	for i := 1; i < len(stockPrices); i++ {
		dx, dy := P[i][0]-P[i-1][0], P[i][1]-P[i-1][1]
		if Dx*dy != Dy*dx {
			t++
		}
		Dx, Dy = dx, dy
	}
	return t
}

// 2900 Longest Unequal Adjacent Groups Subsequences I
func getLongestSubsequence(words []string, groups []int) []string {
	Recursive := func() []string {
		R := []string{}

		calls, r := 0, []string{}
		var Search func(start, g int)
		Search = func(start, g int) {
			calls++
			if start == len(groups) {
				if len(r) > len(R) {
					R = []string{}
					R = append(R, r...)
				}
				return
			}

			Search(start+1, g)
			if g != groups[start] {
				r = append(r, words[start])
				Search(start+1, groups[start])
				r = r[:len(r)-1]
			}

		}

		for g := range []int{0, 1} {
			Search(0, g)
		}

		log.Print("-> rCalls ", calls)

		return R
	}
	tBT := time.Now()
	log.Printf(":: Recursive (@ %[2]v)   %[1]q", Recursive(), time.Since(tBT))

	DP := func() []string {
		R := []string{}

		Lengths := make([]int, len(groups))
		Picks := make([]int, len(groups))

		for i := range len(groups) {
			Lengths[i], Picks[i] = 1, -1
		}

		lMax, iMax := 1, 0
		for l := 1; l < len(groups); l++ {
			for g := 0; g < l; g++ {
				if groups[g] != groups[l] {
					if Lengths[g]+1 > Lengths[l] {
						Lengths[l] = Lengths[g] + 1
						Picks[l] = g
					}
				}
			}

			if Lengths[l] > lMax {
				lMax, iMax = Lengths[l], l
			}
		}

		log.Printf("-> DP %v   %d|%d", Lengths, slices.Max(Lengths), lMax)

		for i := iMax; i != -1; i = Picks[i] {
			R = append(R, words[i])
		}
		slices.Reverse(R)

		return R
	}
	tDP := time.Now()
	log.Printf(":: DP (@ %[2]v)   %[1]q", DP(), time.Since(tDP))

	R, curGroup := []string{words[0]}, groups[0]
	for i, g := range groups[1:] {
		if curGroup != g {
			R = append(R, words[i+1])
			curGroup = g
		}
	}

	log.Printf(":: %q", R)

	return R
}

// 2918m Minimum Equal Sum of Two Arrays After Replacing Zeros
func minSum(nums1 []int, nums2 []int) int64 {
	sum1, zeros1 := int64(0), 0
	for _, n := range nums1 {
		if n == 0 {
			zeros1++
			n++
		}
		sum1 += int64(n)
	}

	sum2, zeros2 := int64(0), 0
	for _, n := range nums2 {
		if n == 0 {
			zeros2++
			n++
		}
		sum2 += int64(n)
	}

	if sum1 > sum2 && zeros2 == 0 || sum2 > sum1 && zeros1 == 0 {
		return -1
	}
	return max(sum1, sum2)
}

// 2938m Separate Black and White Balls
func minimumSteps(s string) int64 {
	steps := int64(0)

	zeros := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeros++
		} else {
			steps += int64(zeros)
		}
	}

	return steps
}
