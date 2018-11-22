package gopher

import (
	"log"
	"math"
	"testing"
)

// 1013 Partition Array Into Three Parts With Equal Sum
func Test1013(t *testing.T) {
	TwoPointers := func(arr []int) bool {
		aSum := 0
		for _, n := range arr {
			aSum += n
		}

		if aSum%3 != 0 {
			return false
		}
		t := aSum / 3

		l, r := 0, len(arr)-1
		lSum, rSum := arr[l], arr[r]
		for l+1 < r {
			if lSum != t {
				l++
				lSum += arr[l]
				continue
			}
			if rSum != t {
				r--
				rSum += arr[r]
				continue
			}

			mSum := 0
			for m := l + 1; m < r; m++ {
				mSum += arr[m]
			}
			if mSum == t {
				return true
			}
		}
		return false
	}

	for _, f := range []func([]int) bool{canThreePartsEqualSum, TwoPointers} {
		log.Print("true ?= ", f([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
		log.Print("false ?= ", f([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
		log.Print("true ?= ", f([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
		log.Print("--")
	}
}

// 1310m XOR Queries of a Subarray
func Test1310(t *testing.T) {
	Optimized := func(arr []int, queries [][]int) []int {
		for i := 1; i < len(arr); i++ {
			arr[i] ^= arr[i-1]
		}

		R := []int{}
		for _, query := range queries {
			i, j := query[0], query[1]
			if i > 0 {
				R = append(R, arr[j]^arr[i-1])
			} else {
				R = append(R, arr[j])
			}
		}
		return R
	}

	for _, fn := range []func([]int, [][]int) []int{xorQueries, Optimized} {
		log.Print("[2 7 14 8] ?= ", fn([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
		log.Print("[8 0 4 4] ?= ", fn([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
		log.Print("--")
	}
}

// 1422 Maximum Score After Splitting a String
func Test1422(t *testing.T) {
	OnePass := func(s string) int {
		xScore := math.MinInt

		zeros, ones := 0, 0
		for i := 0; i < len(s)-1; i++ {
			switch s[i] {
			case '1':
				ones++
			case '0':
				zeros++
			}

			xScore = max(xScore, zeros-ones)
		}

		if s[len(s)-1] == '1' {
			ones++
		}

		return xScore + ones
	}

	for _, f := range []func(string) int{maxScore, OnePass} {
		log.Print("5 ?= ", f("011101"))
		log.Print("5 ?= ", f("00111"))
		log.Print("1 ?= ", f("00"))
		log.Print("2 ?= ", f("01"))
		log.Print("--")
	}
}

// 1991 Find the Middle Index in Array
func Test1991(t *testing.T) {
	log.Print("3 ?= ", findMiddleIndex([]int{2, 3, -1, 8, 4}))
	log.Print("2 ?= ", findMiddleIndex([]int{1, -1, 4}))
	log.Print("-1 ?= ", findMiddleIndex([]int{2, 5}))
}

// 2134m Minimum Swaps to Group All 1's Together II
func Test2134(t *testing.T) {
	// Ai e {0, 1}

	log.Print("1 ?= ", minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	log.Print("2 ?= ", minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
	log.Print("0 ?= ", minSwaps([]int{1, 1, 0, 0, 1}))
}

// 2381m Shifting Letters II
func Test2381(t *testing.T) {
	log.Print("ace ?= ", shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}))
	log.Print("catz ?= ", shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))
}

// 2559m Count Vowel Strings in Ranges
func Test2559(t *testing.T) {
	log.Print("[2 3 0] ?= ", vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
	log.Print("[3 2 1] ?= ", vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}))
}

// 2574 Left and Right Sum Difference
func Test2574(t *testing.T) {
	log.Print("[15 1 11 22] ?= ", leftRightDifference([]int{10, 4, 8, 3}))
	log.Print("[0] ?= ", leftRightDifference([]int{1}))
}

// 2670 Find the Distinct Difference Array
func Test2680(t *testing.T) {
	log.Print("[-3 -1 1 3 5] ?= ", distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	log.Print("[-2 -1 0 2 3] ?= ", distinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

// 3152m Special Array II
func Test3152(t *testing.T) {
	SlidingWindow := func(nums []int, queries [][]int) []bool {
		xReach := make([]int, len(nums))

		end := 0
		for start := range nums {
			end = max(start, end)
			for end < len(nums)-1 && nums[end]&1 != nums[end+1]&1 {
				end++
			}
			xReach[start] = end
		}

		R := []bool{}
		for _, v := range queries {
			start, end := v[0], v[1]
			R = append(R, end <= xReach[start])
		}
		return R
	}

	for _, fn := range []func([]int, [][]int) []bool{isArraySpecial, SlidingWindow} {
		log.Print("[false] ?= ", fn([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))
		log.Print("[false true] ?= ", fn([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}))
		log.Print("--")
	}
}

// 3179m Find the N-th Value After K Seconds
func Test3179(t *testing.T) {
	log.Print("56 ?= ", valueAfterKSeconds(4, 5))
	log.Print("35 ?= ", valueAfterKSeconds(5, 3))
	log.Print("84793457 ?= ", valueAfterKSeconds(5, 1000))
}
