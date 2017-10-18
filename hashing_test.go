package gopher

import (
	"log"
	"testing"
)

// 438m Find All Anagrams in a String
func Test438(t *testing.T) {
	log.Print("[0 6] ?= ", findAnagrams("cbaebabacd", "abc"))
	log.Print("[0 1 2] ?= ", findAnagrams("abab", "ab"))
}
