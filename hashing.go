package gopher

import (
	"bytes"
	"log"
	"maps"
	"slices"
	"strings"
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

// 214h Shortest Palindrome
func shortestPalindrome(s string) string {
	rs := []byte(s)
	slices.Reverse(rs)

	fs := []byte(s)
	log.Printf(" -> %q ~ %q", fs, rs)

	for i := range len(s) {
		if bytes.Equal(fs[:len(s)-i], rs[i:]) {
			return string(rs[:i]) + s
		}
	}
	return ""
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

// 575 Distribute Candies
func distributeCandies(candyType []int) int {
	M := map[int]int{}
	for _, n := range candyType {
		M[n]++
	}

	return min(len(M), len(candyType)/2)
}

// 599 Minimum Index Sum of Two Lists
func findRestaurant(list1, list2 []string) []string {
	for isum := 0; isum < len(list1)+len(list2)-1; isum++ {
		R := []string{}

		for p := 0; p <= isum && p < len(list1); p++ {
			for q := 0; p+q <= isum && q < len(list2); q++ {
				if list1[p] == list2[q] {
					R = append(R, list1[p])
				}
			}
		}

		if len(R) > 0 {
			return R
		}
	}

	return []string{}
}

// 884 Uncommon Words from Two Sentences
func uncommonFromSentences(s1 string, s2 string) []string {
	M := map[string]int{}
	for _, w := range strings.Split(s1+" "+s2, " ") {
		M[w]++
	}

	R := []string{}
	for w, f := range M {
		if f == 1 {
			R = append(R, w)
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

	mMask := map[int]int{0: -1} // mask -> first index in string
	mask, xSub := 0, 0

	for i := range len(s) {
		mask ^= V[s[i]]
		if _, ok := mMask[mask]; !ok && mask > 0 {
			mMask[mask] = i
		}
		xSub = max(i-mMask[mask], xSub)
	}

	return xSub
}

// 1400m Construct K Palindrome Strings
func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	F := [26]int{}
	for i := 0; i < len(s); i++ {
		F[s[i]-'a']++
	}

	oCount := 0
	for _, f := range F {
		oCount += f & 1
	}

	return oCount <= k
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

// 1590m Make Sum Divisible by P
func minSubarray(nums []int, p int) int {
	tSum := 0
	for _, n := range nums {
		tSum = (tSum + n) % p
	}

	t := tSum % p
	if t == 0 {
		return 0
	}

	Mem := map[int]int{0: -1}

	nVal := len(nums)
	curSum := 0
	for i, n := range nums {
		curSum = (curSum + n) % p

		r := (curSum - t + p) % p
		if j, ok := Mem[r]; ok {
			nVal = min(nVal, i-j)
		}

		Mem[curSum] = i
	}

	if nVal == len(nums) {
		return -1
	}
	return nVal
}

// 1930m Unique Length-3 Palindromic Subsequences
func countPalindromicSubsequence(s string) int {
	Left, Right := [26]int{}, [26]int{}
	for l := range Left {
		Left[l] = -1
	}

	for l := 0; l < len(s); l++ {
		if Left[s[l]-'a'] == -1 {
			Left[s[l]-'a'] = l
		}
	}
	for r := len(s) - 1; r >= 0; r-- {
		if Right[s[r]-'a'] == 0 {
			Right[s[r]-'a'] = r
		}
	}

	log.Print(" -> ", Left)
	log.Print(" -> ", Right)

	tCount := 0
	for v := 0; v < 26; v++ {
		if Left[v] != -1 {
			M := map[byte]bool{}
			for i := Left[v] + 1; i < Right[v]; i++ {
				M[s[i]] = true
			}
			tCount += len(M)
		}
	}

	return tCount
}

// 2206 Divide Array Into Equal Pairs
func divideArray(nums []int) bool {
	F := make([]int, 500+1)
	for _, n := range nums {
		F[n]++
	}

	for _, f := range F {
		if f&1 == 1 {
			return false
		}
	}
	return true
}

// 2342m Max Sum of a Pair With Equal Sum of Digits
func maximumSum(nums []int) int {
	xSum := -1

	M := make([]int, 9*9+1)
	for _, n := range nums {
		dSum := 0
		for x := n; x > 0; x /= 10 {
			dSum += x % 10
		}

		if M[dSum] > 0 {
			xSum = max(M[dSum]+n, xSum)
		}
		M[dSum] = max(n, M[dSum])
	}

	return xSum
}

// 2425m Bitwise XOR of All Parings
func xorAllNums(nums1 []int, nums2 []int) int {
	F := map[int]int{}

	for _, n := range nums1 {
		F[n] += len(nums2)
	}
	for _, n := range nums2 {
		F[n] += len(nums1)
	}

	x := 0
	for n, f := range F {
		if f&1 == 1 {
			x ^= n
		}
	}
	return x
}

// 2491m Divide Players Into Teams of Equal Skill
func dividePlayers(skill []int) int64 {
	F := map[int]int{}
	for _, s := range skill {
		F[s]++
	}

	log.Print(" -> ", F)

	t := slices.Min(skill) + slices.Max(skill)
	chemistry := int64(0)
	for s, f := range F {
		if F[t-s] != f {
			return -1
		}
		chemistry += int64(s * (t - s) * f)
	}
	return chemistry / 2
}

// 2570 Merge Two 2D Arrays by Summing Values
func mergeArrays(nums1, nums2 [][]int) [][]int {
	M := make([]int, 1000+1)

	for _, Nums := range [][][]int{nums1, nums2} {
		for _, v := range Nums {
			M[v[0]] += v[1]
		}
	}

	R := [][]int{}
	for id, val := range M {
		if val > 0 {
			R = append(R, []int{id, val})
		}
	}

	return R
}

// 2661m First Completely Painted Row or Column
func firstCompleteIndex(arr []int, mat [][]int) int {
	M := map[int][2]int{}

	Rows, Cols := len(mat), len(mat[0])
	for r := range Rows {
		for c := range Cols {
			M[mat[r][c]] = [2]int{r, c}
		}
	}

	log.Print(" -> ", M)

	rCount, cCount := make([]int, Rows), make([]int, Cols)
	for i, n := range arr {
		r, c := M[n][0], M[n][1]
		rCount[r]++
		cCount[c]++

		if rCount[r] == Cols || cCount[c] == Rows {
			return i
		}
	}

	return -1
}

// 2965 Find Missing and Repeated Values
func findMissingAndRepeatedValues(grid [][]int) []int {
	M := make([]int, len(grid)*len(grid)+1)
	for x := range grid {
		for y := range grid[x] {
			M[grid[x][y]]++
		}
	}

	R := []int{0, 0}
	for n := 1; n < len(M); n++ {
		switch M[n] {
		case 2:
			R[0] = n
		case 0:
			R[1] = n
		}
	}

	return R
}

// 2981m Find Longest Special Substring That Counts Thrice I
func maximumLength(s string) int {
	Count := map[string]int{}

	for start := 0; start < len(s); start++ {
		B := []byte{}

		for end := start; end < len(s); end++ {
			if len(B) == 0 || B[len(B)-1] == s[end] {
				B = append(B, s[end])
				Count[string(B)]++
			} else {
				break
			}
		}
	}

	lMax := 0
	for s, count := range Count {
		if count >= 3 && len(s) > lMax {
			lMax = len(s)
		}
	}

	if lMax == 0 {
		return -1
	}
	return lMax
}

// 3305m Count of Substrings Containing Every Vowel and K Consonants I
func countOfSubstrings(word string, k int) int {
	Mask := [6]int{}
	Vows := []byte("aeiou")

	Update := func(letter byte, diff int) {
		i := 0
		for i < 5 && letter != Vows[i] {
			i++
		}
		Mask[i] += diff
	}
	Good := func() bool {
		for i := range 5 {
			if Mask[i] == 0 {
				return false
			}
		}
		return Mask[5] == k
	}

	t := 0
	for r := range len(word) {
		Update(word[r], 1)

		M := Mask
		for l := 0; l <= r-5-k+1; l++ {
			if Good() {
				t++
			}
			Update(word[l], -1)
		}

		Mask = M
	}

	return t
}

// 3306m Count of Substrings Containing Every Vowel and K Consonants II
func countOfSubstringsII(word string, k int) int64 {
	AtLeast := func(k int) int64 {
		t := int64(0)
		M := map[byte]int{}

		l, consts := 0, 0
		for r := range len(word) {
			switch word[r] {
			case 'a', 'e', 'i', 'o', 'u':
				M[word[r]]++
			default:
				consts++
			}

			for len(M) == 5 && consts >= k {
				t += int64(len(word) - r)

				switch word[l] {
				case 'a', 'e', 'i', 'o', 'u':
					M[word[l]]--
					if M[word[l]] == 0 {
						delete(M, word[l])
					}
				default:
					consts--
				}

				l++
			}
		}
		return t
	}

	return AtLeast(k) - AtLeast(k+1)
}
