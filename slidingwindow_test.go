package gopher

import (
	"log"
	"testing"
)

// 1358m Number of Substrings Containing All Three Characters
func Test1358(t *testing.T) {
	for _, c := range []struct {
		rst int
		s   string
	}{
		{10, "abcabc"},
		{3, "aaacb"},
		{1, "abc"},
	} {
		rst, s := c.rst, c.s
		log.Printf("%d ?= %d", rst, numberOfSubstrings(s))
		if rst != numberOfSubstrings(s) {
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
