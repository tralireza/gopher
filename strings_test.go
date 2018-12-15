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

// 557 Reverse Words in a String III
func Test557(t *testing.T) {
	log.Printf(`"s'teL ekat edoCteeL tsetnoc" ?= %q`, reverseWords("Let's take LeetCode contest"))
	log.Printf(`"rM gniD" ?= %q`, reverseWords("Mr Ding"))
}

// 696 Count Binary Substrings
func Test696(t *testing.T) {
	log.Print("6 ?= ", countBinarySubstrings("00110011"))
	log.Print("4 ?= ", countBinarySubstrings("10101"))
}

// 824 Goat Latin
func Test824(t *testing.T) {
	// 1 <= S.Length

	if toGoatLatin("I speak Goat Latin") != "Imaa peaksmaaa oatGmaaaa atinLmaaaaa" {
		t.Fail()
	}
	if toGoatLatin("The quick brown fox jumped over the lazy dog") != "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa" {
		t.Fail()
	}
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
