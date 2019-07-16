package gopher

import (
	"log"
	"testing"
)

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
