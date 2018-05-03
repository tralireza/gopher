package gopher

import (
	"log"
	"testing"
)

// 212h Word Search II
func Test212(t *testing.T) {
	log.Printf(`["oath" "eat"] ?= %q`, findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
	log.Printf(`[] ?= %q`, findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcd"}))
}

// 440h K-th Smallest in Lexicographical Order
func Test440(t *testing.T) {
	log.Print("10 ?= ", findKthNumber(13, 2))
	log.Print("1 ?= ", findKthNumber(1, 1))
	log.Print("104 ?= ", findKthNumber(127, 7))
}

// 2416h Sum of Prefix Score of Strings
func Test2416(t *testing.T) {
	log.Print("[5 4 3 2] ?= ", sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
	log.Print("[4] ?= ", sumPrefixScores([]string{"abcd"}))
}

// 3043m Find the Length of the Longest Common Prefix
func Test3043(t *testing.T) {
	log.Print("3 ?= ", longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	log.Print("0 ?= ", longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}
