package gopher

import (
	"log"
	"math"
	"math/bits"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
)

type Assert struct {
	t *testing.T
}

func (o *Assert) Equal(want, got any, msgs ...any) {
	if want != got {
		o.t.Error(msgs...)
	}
}

// 30h Substring With Concatenation of All Words
func Test30(t *testing.T) {
	log.Print("[0 9] ?= ", findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	log.Print("[] ?= ", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	log.Print("[6 9 12] ?= ", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	log.Print("[1 2 3 4 5 6 7 8 9 10] ?= ", findSubstring("aaaaaaaaaaaaaa", []string{"aa", "aa"}))
	log.Print("[97] ?= ", findSubstring("abbaccaaabcabbbccbabbccabbacabcacbbaabbbbbaaabaccaacbccabcbababbbabccabacbbcabbaacaccccbaabcabaabaaaabcaabcacabaa", []string{"cac", "aaa", "aba", "aab", "abc"}))
}

// 76h Minimum Window Substring
func Test76(t *testing.T) {
	log.Print("BANC ?= ", minWindow("ADOBECODEBANC", "ABC"))
	log.Print("a ?= ", minWindow("aa", "a"))
	log.Print(" ?= ", minWindow("a", "aa"))
}

// 214h Shortest Palindrome
func Test214(t *testing.T) {
	log.Print("aaacecaaa ?= ", shortestPalindrome("aacecaaa"))
	log.Print("dcbabcd ?= ", shortestPalindrome("abcd"))
}

// 380m Insert Delete GetRandom O(1)
type RandomizedSet struct {
	Mem  map[int]int
	Vals []int
}

func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		Mem:  map[int]int{},
		Vals: []int{},
	}
}

func (o *RandomizedSet) Insert(v int) bool {
	if _, ok := o.Mem[v]; ok {
		return false
	}
	o.Mem[v] = len(o.Vals)
	o.Vals = append(o.Vals, v)
	return true
}
func (o *RandomizedSet) Remove(v int) bool {
	if i, ok := o.Mem[v]; ok {
		// o.Vals: swap last value with this one, remove last, also update index in map

		lastVal := o.Vals[len(o.Vals)-1]
		o.Vals[i] = lastVal
		o.Mem[lastVal] = i

		o.Vals = o.Vals[:len(o.Vals)-1] // trim it
		delete(o.Mem, v)
		return true
	}
	return false
}
func (o *RandomizedSet) GetRandom() int { return o.Vals[rand.Intn(len(o.Vals))] }

func Test380(t *testing.T) {
	mSet := NewRandomizedSet()
	mSet.Insert(1)
	log.Print(mSet, " :: ", mSet.GetRandom())
	mSet.Remove(1)
	log.Print(mSet)

	for range 16 {
		v := rand.Intn(16)
		log.Printf("---I-> %d :: %v", v, mSet.Insert(v))
	}
	log.Print(mSet)
	for range 8 {
		v := rand.Intn(16)
		log.Printf("---D-> %d :: %v", v, mSet.Remove(v))
	}
	log.Print(mSet)
	for range 8 {
		log.Printf("<-G--- %d", mSet.GetRandom())
	}
}

// 438m Find All Anagrams in a String
func Test438(t *testing.T) {
	log.Print("[0 6] ?= ", findAnagrams("cbaebabacd", "abc"))
	log.Print("[0 1 2] ?= ", findAnagrams("abab", "ab"))
}

// 575 Distribute Candies
func Test575(t *testing.T) {
	log.Print("3 ?= ", distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	log.Print("2 ?= ", distributeCandies([]int{1, 1, 2, 3}))
	log.Print("1 ?= ", distributeCandies([]int{6, 6, 6, 6}))
}

// 599 Minimum Index Sum of Two Lists
func Test599(t *testing.T) {
	// 1 <= N1, N2 <= 1000

	WithHash := func(list1, list2 []string) []string {
		M := map[int][]string{}

		MList := map[string]int{}
		for p, str := range list1 {
			MList[str] = p
		}

		for q, str := range list2 {
			if p, ok := MList[str]; ok {
				M[p+q] = append(M[p+q], str)
			}
		}

		nSum := math.MaxInt
		for p := range M {
			nSum = min(p, nSum)
		}

		if nSum == math.MaxInt {
			return []string{}
		}
		return M[nSum]
	}

	for _, f := range []func([]string, []string) []string{findRestaurant, WithHash} {
		log.Printf(`["Shogun"] ?= %q`, f(
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
		log.Printf(`["Shogun"] ?= %q`, f([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
		log.Printf(`["happy" "sad"] ?= %q`, f([]string{"happy", "sad", "good"}, []string{"sad", "happy", "good"}))
		log.Print("--")
	}
}

// 884 Uncommon Words from Two Sentences
func Test884(t *testing.T) {
	log.Print(" ?= ", uncommonFromSentences("this apple is sweet", "this apple is sour"))
	log.Print(" ?= ", uncommonFromSentences("apple apple", "banana"))
}

// 874m Walking Robot Simulation
func Test874(t *testing.T) {
	log.Print("25 ?= ", robotSim([]int{4, -1, 3}, [][]int{}))
	log.Print("65 ?= ", robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
	log.Print("36 ?= ", robotSim([]int{6, -1, -1, 6}, [][]int{}))
}

// 1372m Find the Longest Substring Containing Vowels in Even Counts
func Test1372(t *testing.T) {
	log.Print("13 ?= ", findTheLongestSubstring("eleetminicoworoep"))
	log.Print("5 ?= ", findTheLongestSubstring("leetcodeisgreat"))
	log.Print("6 ?= ", findTheLongestSubstring("bcbcbc"))
}

// 1400m Construct K Palindrome Strings
func Test1400(t *testing.T) {
	Bits := func(s string, k int) bool {
		if len(s) < k {
			return false
		}
		if len(s) == k {
			return true
		}

		oCount := uint(0)
		for i := 0; i < len(s); i++ {
			oCount ^= 1 << (s[i] - 'a')
		}
		return bits.OnesCount(oCount) < k
	}

	for _, f := range []func(string, int) bool{canConstruct, Bits} {
		log.Print("true ?= ", f("annabelle", 2))
		log.Print("false ?= ", f("leetcode", 3))
		log.Print("true ?= ", f("true", 4))
		log.Print("--")
	}
}

// 1460 Make Two Arrays Equal by Reversing Subarrays
func Test1460(t *testing.T) {
	// 1 <= Ai <= 10000

	Optimized := func(target, arr []int) bool {
		hA := make([]int, 10_000+1)
		for _, n := range arr {
			hA[n]++
		}

		for _, n := range target {
			if hA[n] == 0 {
				return false
			}
			hA[n]--
		}
		return true
	}

	for _, f := range []func([]int, []int) bool{canBeEqual, Optimized} {
		log.Print("true ?= ", f([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
		log.Print("true ?= ", f([]int{7}, []int{7}))
		log.Print("false ?= ", f([]int{3, 7, 9}, []int{3, 7, 11}))
		log.Print("--")
	}
}

// 1590m Make Sum Divisible by P
func NewHashMap1590(capacity int) HashMap1590 {
	return HashMap1590{
		make([][]MapEntry1590, capacity),
	}
}

// 1930m Unique Length-3 Palindromic Subsequences
func Test1930(t *testing.T) {
	log.Print("3 ?= ", countPalindromicSubsequence("aabca"))
	log.Print("0 ?= ", countPalindromicSubsequence("abc"))
	log.Print("4 ?= ", countPalindromicSubsequence("bbcbaba"))
}

type HashMap1590 struct {
	Store [][]MapEntry1590
}

type MapEntry1590 struct {
	Key, Value int
}

func (o *HashMap1590) Hash(k int) int {
	const MP8 = 0x7fffffff // 2^31-1 (8th Mersenne Prime)
	m := len(o.Store)
	return (k*MP8%m + m) % m
}
func (o *HashMap1590) Set(k, v int) {
	o.Store[o.Hash(k)] = append(o.Store[o.Hash(k)], MapEntry1590{k, v})
}
func (o *HashMap1590) Get(k int) (int, bool) {
	for _, e := range o.Store[o.Hash(k)] {
		if e.Key == k {
			return e.Value, true
		}
	}
	return 0, false
}

func Test1590(t *testing.T) {
	h := NewHashMap1590(9)
	for i := range 5 {
		h.Set(i, i)
	}
	h.Set(14345923, 14345923)
	log.Print(h.Get(0))
	log.Print(h.Get(14345923))
	log.Print(h.Get(5))
	log.Print(h.Store)
	log.Print("--")

	log.Print("1 ?= ", minSubarray([]int{3, 1, 4, 2}, 6))
	log.Print("2 ?= ", minSubarray([]int{6, 3, 5, 2}, 9))
	log.Print("0 ?= ", minSubarray([]int{1, 2, 3}, 3))
}

func Test2206(t *testing.T) {
	for _, c := range []struct {
		rst  bool
		nums []int
	}{
		{true, []int{3, 2, 3, 2, 2, 2}},
		{false, []int{1, 2, 3, 4}},
	} {
		rst, nums := c.rst, c.nums
		if rst != divideArray(nums) {
			t.FailNow()
		}
		log.Printf(":: %v <- %v", rst, nums)
	}
}

// 2342m Max Sum of a Pair With Equal Sum of Digits
func Test2342(t *testing.T) {
	a := Assert{t}

	a.Equal(54, maximumSum([]int{18, 43, 36, 13, 7}))
	a.Equal(-1, maximumSum([]int{10, 12, 19, 14}))
	a.Equal(872, maximumSum([]int{279, 169, 463, 252, 94, 455, 423, 315, 288, 64, 494, 337, 409, 283, 283, 477, 248, 8, 89, 166, 188, 186, 128}))
}

// 2425m Bitwise XOR of All Parings
func Test2425(t *testing.T) {
	log.Print("13 ?= ", xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0}))
	log.Print("0 ?= ", xorAllNums([]int{1, 2}, []int{3, 4}))
}

// 2491m Divide Players Into Teams of Equal Skill
func Test2491(t *testing.T) {
	log.Print("22 ?= ", dividePlayers([]int{3, 2, 5, 1, 3, 4}))
	log.Print("12 ?= ", dividePlayers([]int{3, 4}))
	log.Print("-1 ?= ", dividePlayers([]int{1, 1, 2, 3}))
}

// 2570 Merge Two 2D Arrays by Summing Values
func Test2570(t *testing.T) {
	// 1 <= N_i[id, val] <= 1000

	TwoPointers := func(nums1, nums2 [][]int) [][]int {
		R := [][]int{}

		l, r := 0, 0

		for l < len(nums1) && r < len(nums2) {
			if nums1[l][0] == nums2[r][0] {
				R = append(R, []int{nums1[l][0], nums1[l][1] + nums2[r][1]})
				l++
				r++
			} else if nums1[l][0] < nums2[r][0] {
				R = append(R, nums1[l])
				l++
			} else {
				R = append(R, nums2[r])
				r++
			}
		}

		for l < len(nums1) {
			R = append(R, nums1[l])
			l++
		}
		for r < len(nums2) {
			R = append(R, nums2[r])
			r++
		}

		return R
	}

	for _, f := range []func([][]int, [][]int) [][]int{mergeArrays, TwoPointers} {
		log.Print(" ?= ", f([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}))
		log.Print(" ?= ", f([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}))
		log.Print("--")
	}
}

// 2661m First Completely Painted Row or Column
func Test2661(t *testing.T) {
	ReverseMapping := func(arr []int, mat [][]int) int {
		M := map[int]int{}
		for i, n := range arr {
			M[n] = i
		}

		bVal := len(arr)

		for r := range len(mat) {
			curVal := -1
			for c := range len(mat[r]) {
				curVal = max(M[mat[r][c]], curVal)
			}
			bVal = min(curVal, bVal)
		}

		for c := range len(mat[0]) {
			curVal := -1
			for r := range len(mat) {
				curVal = max(M[mat[r][c]], curVal)
			}
			bVal = min(curVal, bVal)
		}

		return bVal
	}

	for _, f := range []func([]int, [][]int) int{firstCompleteIndex, ReverseMapping} {
		log.Print("2 ?= ", f([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))
		log.Print("3 ?= ", f([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}))
		log.Print("--")
	}
}

// 2965 Find Missing and Repeated Values
func Test2965(t *testing.T) {
	// E (1+...+n^2) = n^2(n^2+1)/2, E (1^2+...+(n^2)^2) = n^2(n^2+1)(2n^2+1)/6

	Math := func(grid [][]int) []int {
		n := len(grid)

		sum, sqrSum := 0, 0
		for r := range grid {
			for c := range grid[r] {
				sum += grid[r][c]
				sqrSum += grid[r][c] * grid[r][c]
			}
		}

		n *= n
		diffSum, diffSqr := sum-n*(n+1)/2, sqrSum-n*(n+1)*(2*n+1)/6

		return []int{(diffSqr/diffSum + diffSum) / 2, (diffSqr/diffSum - diffSum) / 2}
	}

	for _, f := range []func([][]int) []int{findMissingAndRepeatedValues, Math} {
		log.Print("[2 4] ?= ", f([][]int{{1, 3}, {2, 2}}))
		log.Print("[9 5] ?= ", f([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}))
		log.Print("--")
	}
}

// 2981m Find Longest Special Substring That Counts Thrice I
func Test2981(t *testing.T) {
	// O(n)
	OnePass := func(s string) int {
		Count := map[[2]int]int{}

		start := 0
		for start < len(s) {
			l := 0

			end := start
			for end < len(s) && s[start] == s[end] {
				l++
				end++
			}
			Count[[2]int{int(s[start]), l}]++

			for l > 1 {
				l--
				Count[[2]int{int(s[start]), l}] += Count[[2]int{int(s[start]), l + 1}] + 1
			}

			start = end
		}

		log.Print(" -> ", Count)

		lMax := 0
		for e, count := range Count {
			if count >= 3 {
				lMax = max(e[1], lMax)
			}
		}

		if lMax == 0 {
			return -1
		}
		return lMax
	}

	// O(n^2)
	Optimized := func(s string) int {
		Count := map[[2]int]int{}

		for start := 0; start < len(s); start++ {
			l := 0
			for end := start; end < len(s); end++ {
				if s[start] == s[end] {
					l++
					Count[[2]int{int(s[start]), l}]++
				} else {
					break
				}
			}
		}

		lMax := 0
		for e, count := range Count {
			if count >= 3 {
				lMax = max(e[1], lMax)
			}
		}

		if lMax == 0 {
			return -1
		}
		return lMax
	}

	for _, fn := range []func(string) int{maximumLength, Optimized, OnePass} {
		log.Print("2 ?= ", fn("aaaa"))
		log.Print("-1 ?= ", fn("abcdef"))
		log.Print("1 ?= ", fn("abcaba"))
		log.Print("4 ?= ", fn("cddedeedccedcedecdedcdeededdddcdddddcdeecdcddeecdc"))
		log.Print("--")
	}
}

// 3305m Count of Substrings Containing Every Vowel and K Consonants I
func Test3305(t *testing.T) {
	// 5 <= word <= 250

	log.Print("0 ?= ", countOfSubstrings("aeioqq", 1))
	log.Print("1 ?= ", countOfSubstrings("aeiou", 0))
	log.Print("3 ?= ", countOfSubstrings("ieaouqqieaouqq", 1))
}

// 3306m Count of Substrings Containing Every Vowel and K Consonants II
func Test3306(t *testing.T) {
	// 5 <= word <= 2*10^5

	SlidingWindow := func(word string, k int) int64 {
		Next := make([]int, len(word))
		nextCons := len(word)
		for i := len(word) - 1; i >= 0; i-- {
			Next[i] = nextCons
			switch word[i] {
			case 'a', 'e', 'i', 'o', 'u':
			default:
				nextCons = i
			}
		}

		t := int64(0)

		M := map[byte]int{}
		l := 0

		for r := 0; r < len(word); r++ {
			switch word[r] {
			case 'a', 'e', 'i', 'o', 'u':
				M[word[r]]++
			default:
				k--
			}

			for k < 0 {
				switch word[l] {
				case 'a', 'e', 'i', 'o', 'u':
					M[word[l]]--
					if M[word[l]] == 0 {
						delete(M, word[l])
					}
				default:
					k++
				}
				l++
			}

			for l < len(word) && k == 0 && len(M) == 5 {
				t += int64(Next[r] - r)
				switch word[l] {
				case 'a', 'e', 'i', 'o', 'u':
					M[word[l]]--
					if M[word[l]] == 0 {
						delete(M, word[l])
					}
				default:
					k++
				}
				l++
			}

		}

		return t
	}

	for _, c := range []struct {
		rst  int64
		word string
		k    int
	}{
		{0, "aeioqq", 1},
		{1, "aeiou", 0},
		{3, "ieaouqqieaouqq", 1},
	} {
		rst, word, k := c.rst, c.word, c.k
		for _, f := range []func(string, int) int64{countOfSubstringsII, SlidingWindow} {
			log.Printf("%27s :: %d ?= %d", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), k, f(word, k))
			if rst != f(word, k) {
				t.FailNow()
			}
		}
	}
}

func Test3375(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
		k    int
	}{
		{2, []int{5, 2, 5, 4, 5}, 2},
		{-1, []int{2, 1, 2}, 2},
		{4, []int{9, 7, 5, 3}, 1},
	} {
		if c.rst != minOperations_EqualK(c.nums, c.k) {
			t.FailNow()
		}
		log.Printf(":: %d   <- %v | %d", c.rst, c.nums, c.k)
	}
}
