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
