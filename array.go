// Array
package gopher

import (
	"log"
	"slices"
)

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
	R := make([]int, len(nums))
	for i, n := range nums {
		R[i] = nums[n]
	}

	return R
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
