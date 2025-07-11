package gopher

import (
	"log"
	"testing"
)

func Test594(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{5, []int{1, 3, 2, 2, 5, 2, 3, 7}},
		{2, []int{1, 2, 3, 4}},
		{0, []int{1, 1, 1, 1}},
	} {
		log.Print("* ", c.nums)
		if c.rst != findLHS(c.nums) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

// 1358m Number of Substrings Containing All Three Characters
func Test1358(t *testing.T) {
	SlidingWindow := func(s string) int {
		M, wSize := [3]int{}, 0

		l, count := 0, 0
		for r := 0; r < len(s); r++ {
			if M[s[r]-'a'] == 0 {
				wSize++
			}
			M[s[r]-'a']++

			for wSize == 3 {
				count += len(s) - r

				M[s[l]-'a']--
				if M[s[l]-'a'] == 0 {
					wSize--
				}

				l++
			}
		}

		return count
	}

	for _, c := range []struct {
		rst int
		s   string
	}{
		{10, "abcabc"},
		{3, "aaacb"},
		{1, "abc"},
	} {
		rst, s := c.rst, c.s
		for _, f := range []func(string) int{numberOfSubstrings, SlidingWindow} {
			log.Printf("%d ?= %d", rst, f(s))
			if rst != numberOfSubstrings(s) {
				t.FailNow()
			}
		}
	}
}

func Test2401(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{3, []int{1, 3, 8, 48, 10}},
		{1, []int{3, 1, 5, 11, 13}},
	} {
		rst, nums := c.rst, c.nums
		if rst != longestNiceSubarray(nums) {
			t.FailNow()
		}
		log.Printf("** %v <- %v ", rst, nums)
	}
}

func Test2444(t *testing.T) {
	for _, c := range []struct {
		rst        int64
		nums       []int
		minK, maxK int
	}{
		{2, []int{1, 3, 5, 2, 7, 5}, 1, 5},
		{10, []int{1, 1, 1, 1}, 1, 1},
	} {
		if c.rst != countSubarrays(c.nums, c.minK, c.maxK) {
			t.FailNow()
		}
	}
}

func Test2537(t *testing.T) {
	for _, c := range []struct {
		rst  int64
		nums []int
		k    int
	}{
		{1, []int{1, 1, 1, 1, 1}, 10},
		{4, []int{3, 1, 4, 3, 2, 2, 4}, 2},
	} {
		if c.rst != countGood(c.nums, c.k) {
			t.FailNow()
		}
	}
}

func Test2799(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{4, []int{1, 3, 1, 2, 2}},
		{10, []int{10, 10, 10, 10}},
	} {
		if c.rst != countCompleteSubarrays(c.nums) {
			t.FailNow()
		}
	}
}

// 3208m Alternating Groups II
func Test3208(t *testing.T) {
	for _, c := range []struct {
		rst    int
		colors []int
		k      int
	}{
		{3, []int{0, 1, 0, 1, 0}, 3},
		{2, []int{0, 1, 0, 0, 1, 0, 1}, 6},
		{0, []int{1, 1, 0, 1}, 4},
	} {
		rst, colors, k := c.rst, c.colors, c.k
		log.Printf("%d ?= %v", rst, numberOfAlternatingGroups(colors, k))
		if rst != numberOfAlternatingGroups(colors, k) {
			t.FailNow()
		}
	}
}

func Test3439(t *testing.T) {
	for _, c := range []struct {
		rst                int
		eventTime, k       int
		startTime, endTime []int
	}{
		{2, 5, 1, []int{1, 3}, []int{2, 5}},
		{6, 10, 1, []int{0, 2, 9}, []int{1, 4, 10}},
		{0, 5, 2, []int{0, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5}},

		{45, 96, 3, []int{4, 11, 16, 53}, []int{11, 16, 27, 77}}, // 479/689
	} {
		log.Print("* ", c.eventTime, c.k, c.startTime, c.endTime)
		if c.rst != maxFreeTime(c.eventTime, c.k, c.startTime, c.endTime) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test3445(t *testing.T) {
	for _, c := range []struct {
		rst int
		s   string
		k   int
	}{
		{-1, "12233", 4},
		{1, "1122211", 3},
		{-1, "110", 3},
	} {
		if c.rst != maxDifference_II(c.s, c.k) {
			t.FailNow()
		}
	}
}
