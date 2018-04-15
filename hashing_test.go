package gopher

import (
	"log"
	"math/rand"
	"testing"
)

// 30h Substring With Concatenation of All Words
func Test30(t *testing.T) {
	log.Print("[0 9] ?= ", findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	log.Print("[] ?= ", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	log.Print("[6 9 12] ?= ", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	log.Print("[1 2 3 4 5 6 7 8 9 10] ?= ", findSubstring("aaaaaaaaaaaaaa", []string{"aa", "aa"}))
	log.Print("[97] ?= ", findSubstring("abbaccaaabcabbbccbabbccabbacabcacbbaabbbbbaaabaccaacbccabcbababbbabccabacbbcabbaacaccccbaabcabaabaaaabcaabcacabaa", []string{"cac", "aaa", "aba", "aab", "abc"}))
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
