package gopher

import (
	"log"
	"testing"
)

// 412 Fizz Buzz
func Test412(t *testing.T) {
	log.Print(" ?= ", fizzBuzz(3))
	log.Print(" ?= ", fizzBuzz(5))
	log.Print(" ?= ", fizzBuzz(15))
}

// 520 Detect Capital
func Test520(t *testing.T) {
	log.Print("true ?= ", detectCapitalUse("USA"))
	log.Print("false ?= ", detectCapitalUse("FlaG"))
}

// 1813m Sentence Similarity III
func Test1813(t *testing.T) {
	log.Print("true ?= ", areSentencesSimilar("Hello Jane", "Hello my name is Jane"))
	log.Print("false ?= ", areSentencesSimilar("of", "A lot of words"))
	log.Print("true ?= ", areSentencesSimilar("Eating right now", "Eating"))
}

// 2185 Counting Words With a Given Prefix
func Test2185(t *testing.T) {
	log.Print("2 ?= ", prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
	log.Print("0 ?= ", prefixCount([]string{"leetcode", "win", "loops", "success"}, "code"))
}

// 2405m Optimal Partition of String
func Test2405(t *testing.T) {
	log.Print("4 ?= ", partitionString("abacaba"))
	log.Print("6 ?= ", partitionString("ssssss"))
}

// 3042 Count Prefix and Suffix Pairs I
func Test3042(t *testing.T) {
	log.Print("4 ?= ", countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}))
	log.Print("2 ?= ", countPrefixSuffixPairs([]string{"pa", "papa", "ma", "mama"}))
	log.Print("0 ?= ", countPrefixSuffixPairs([]string{"abab", "ab"}))
}
