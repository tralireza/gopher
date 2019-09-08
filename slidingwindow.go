package gopher

import "log"

// 1358m Number of Substrings Containing All Three Characters
func numberOfSubstrings(s string) int {
	count := 0

	M := map[byte]int{}
	l := 0

	for r := 0; r < len(s); r++ {
		M[s[r]]++

		for len(M) == 3 {
			count += len(s) - r

			M[s[l]]--
			if M[s[l]] == 0 {
				delete(M, s[l])
			}

			l++
		}
	}

	return count
}

// 2401m Longest Nice Subarray
func longestNiceSubarray(nums []int) int {
	xLen := 0
	bits := 0

	l, r := 0, 0
	for r < len(nums) {
		for bits&nums[r] != 0 {
			bits ^= nums[l]
			l++
		}

		bits |= nums[r]

		log.Printf("%032b <- %032b", bits, nums[r])

		xLen = max(r-l+1, xLen)
		r++
	}

	return xLen
}

// 2537m Count the Number of Good Subarrays
func countGood(nums []int, k int) int64 {
	Map := map[int]int{}

	w := 0
	left, count := 0, int64(0)
	for right, n := range nums {
		Map[n]++
		f := Map[n]
		w += f*(f-1)/2 - (f-1)*(f-2)/2 // C_n/k :: n!/(n-2)!2! ~ n*(n-1)/2

		for w >= k {
			count += int64(len(nums) - right)

			Map[nums[left]]--
			f := Map[nums[left]]
			w -= (f+1)*f/2 - f*(f-1)/2

			left++
		}
	}

	return count
}

// 2799m Count Complete Subarrays in an Array
func countCompleteSubarrays(nums []int) int {
	// 1 <= N <= 1000, 1 <= N_i <= 2000
	Hashing := func(nums []int) int {
		S := map[int]struct{}{}
		for _, n := range nums {
			S[n] = struct{}{}
		}

		log.Print("-> Set: ", S)

		count, k := 0, len(S)
		l := 0
		M, wSize := map[int]int{}, 0
		for r := range nums {
			M[nums[r]]++
			if M[nums[r]] == 1 {
				wSize++
			}

			for wSize == k {
				count += len(nums) - r
				M[nums[l]]--
				if M[nums[l]] == 0 {
					wSize--
				}
				l++
			}
		}

		return count
	}
	log.Print(":: ", Hashing(nums))

	M, k := make([]int, 2000+1), 0
	for _, n := range nums {
		M[n]++
		if M[n] == 1 {
			k++
		}
	}
	clear(M)

	count, wSize := 0, 0
	l := 0
	for r := range nums {
		M[nums[r]]++
		if M[nums[r]] == 1 {
			wSize++
		}

		for wSize == k {
			count += len(nums) - r

			M[nums[l]]--
			if M[nums[l]] == 0 {
				wSize--
			}
			l++
		}
	}

	return count
}

// 3208m Alternating Groups II
func numberOfAlternatingGroups(colors []int, k int) int {
	groups, wSize := 0, 1

	prvColor := colors[0]
	for i := 1; i < len(colors)+k-1; i++ {
		curColor := colors[i%len(colors)]

		switch curColor {
		case prvColor:
			wSize = 1
		default:
			wSize++
			if wSize >= k {
				groups++
			}
			prvColor = curColor
		}
	}

	return groups
}
