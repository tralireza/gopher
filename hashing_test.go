package gopher

import (
	"log"
	"testing"
)

// 76h Minimum Window Substring
func Test76(t *testing.T) {
	log.Print("BANC ?= ", minWindow("ADOBECODEBANC", "ABC"))
	log.Print("a ?= ", minWindow("aa", "a"))
	log.Print(" ?= ", minWindow("a", "aa"))
}

// 438m Find All Anagrams in a String
func Test438(t *testing.T) {
	log.Print("[0 6] ?= ", findAnagrams("cbaebabacd", "abc"))
	log.Print("[0 1 2] ?= ", findAnagrams("abab", "ab"))
}

// 1460 Make Two Arrays Equal by Reversing Subarrays
func Test1460(t *testing.T) {
	// 1 <= Ai <= 10000

	log.Print("true ?= ", canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
	log.Print("true ?= ", canBeEqual([]int{7}, []int{7}))
	log.Print("false ?= ", canBeEqual([]int{3, 7, 9}, []int{3, 7, 11}))
}
