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
