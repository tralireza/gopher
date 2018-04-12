package gopher

import (
	"maps"
)

// 30h Substring With Concatenation of All Words
func findSubstring(s string, words []string) []int {
	Mem := map[string]int{}
	for _, w := range words {
		Mem[w]++
	}

	lW := len(words[0])      // Word length
	wSize := lW * len(words) // Window size

	R := []int{}
	for x := 0; x < lW; x++ {
		l, r := x, x
		V := map[string]int{}

		for r <= len(s)-lW {
			w := s[r : r+lW]
			if Mem[w] == 0 {
				r += lW
				if r-l > wSize {
					if V[s[l:l+lW]] > 0 {
						V[s[l:l+lW]]--
					}
					l += lW
				}
				continue
			}

			r += lW
			V[w]++

			for r-l > wSize {
				if V[s[l:l+lW]] > 0 {
					V[s[l:l+lW]]--
				}
				l += lW
			}

			if r-l == wSize {
				if maps.Equal(Mem, V) {
					R = append(R, l)
				}
			}
		}
	}
	return R
}

// 76h Minimum Window Substring
func minWindow(s string, t string) string {
	idx := func(l byte) int {
		switch {
		case 'a' <= l && l <= 'z':
			return int(l - 'a')
		default:
			return int(l-'A') + 26
		}
	}

	hS, hT := [52]int{}, [52]int{}

	isGood := func() bool {
		for i := 0; i < 52; i++ {
			if hS[i] < hT[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(t); i++ {
		hT[idx(t[i])]++
	}

	l, minStr := 0, ""
	for r := 0; r < len(s); r++ {
		hS[idx(s[r])]++

		for isGood() {
			if minStr == "" || r-l < len(minStr) {
				minStr = s[l : r+1]
			}

			hS[idx(s[l])]--
			l++
		}
	}
	return minStr
}

// 438m Find All Anagrams in a String
func findAnagrams(s string, p string) []int {
	fP, fS := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		fP[p[i]-'a']++
	}

	R := []int{}
	for i := 0; i < len(s); i++ {
		fS[s[i]-'a']++
		if i >= len(p) {
			fS[s[i-len(p)]-'a']--
		}
		if fS == fP {
			R = append(R, i-len(p)+1)
		}
	}
	return R
}

// 874m Walking Robot Simulation
func robotSim(commands []int, obstacles [][]int) int {
	Obs := make(map[[2]int]bool, len(obstacles))
	for _, p := range obstacles {
		Obs[[2]int{p[0], p[1]}] = true
	}

	Dirs := []int{0, 1, 0, -1, 0}
	dir := 0 // Right(-1) -> +1, Left(-2) -> +3  (mod 4)

	dist := 0

	x, y := 0, 0
	for _, c := range commands {
		switch c {
		case -1:
			dir = (dir + 1) % 4
		case -2:
			dir = (dir + 3) % 4
		default:
			for range c {
				X, Y := x+Dirs[dir], y+Dirs[dir+1]
				if Obs[[2]int{X, Y}] {
					break
				}
				x, y = X, Y
				dist = max(x*x+y*y, dist)
			}
		}
	}

	return dist
}

// 1372m Find the Longest Substring Containing Vowels in Even Counts
func findTheLongestSubstring(s string) int {
	V := map[byte]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}

	xSub := 0

	mMask := map[int]int{0: -1} // mask -> first index in string
	mask := 0

	for i := range len(s) {
		mask ^= V[s[i]]
		if _, ok := mMask[mask]; !ok && mask > 0 {
			mMask[mask] = i
		}
		xSub = max(i-mMask[mask], xSub)
	}

	return xSub
}

// 1460 Make Two Arrays Equal by Reversing Subarrays
func canBeEqual(target []int, arr []int) bool {
	hT, hA := make([]int, 10_000+1), make([]int, 10_000+1)
	for i, n := range arr {
		hA[n]++
		hT[target[i]]++
	}

	for n, f := range hA {
		if hT[n] != f {
			return false
		}
	}
	return true
}
