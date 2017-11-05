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
