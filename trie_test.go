package gopher

import (
	"log"
	"strconv"
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

// 1233m Remove Sub-Folders from the Filesystem
func Test1233(t *testing.T) {
	log.Printf(`["/a" "/c/d" "/c/f"] ?= %q`, removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
	log.Printf(`["/a"] ?= %q`, removeSubfolders([]string{"/a", "/a/b/c", "/a/b/d"}))
	log.Printf(`["/a/b/c" "/a/b/ca" "/a/b/d"] ?= %q`, removeSubfolders([]string{"/a/b/c", "/a/b/ca", "/a/b/d"}))
}

// 2416h Sum of Prefix Score of Strings
func Test2416(t *testing.T) {
	log.Print("[5 4 3 2] ?= ", sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
	log.Print("[4] ?= ", sumPrefixScores([]string{"abcd"}))
}

// 3043m Find the Length of the Longest Common Prefix
func Test3043(t *testing.T) {
	Optimized := func(arr1, arr2 []int) int {
		type Trie struct {
			Child [10]*Trie
		}

		trie := &Trie{}
		Insert := func(w string) {
			n := trie
			for i := 0; i < len(w); i++ {
				c := n.Child[w[i]-'0']
				if c == nil {
					c = &Trie{}
					n.Child[w[i]-'0'] = c
				}
				n = c
			}
		}
		Search := func(w string) bool {
			n := trie
			for i := 0; i < len(w); i++ {
				c := n.Child[w[i]-'0']
				if c == nil {
					return false
				}
				n = c
			}
			return true
		}

		for _, n := range arr1 {
			w := strconv.Itoa(n)
			Insert(w)
		}

		xVal := 0
		for _, n := range arr2 {
			w := strconv.Itoa(n)
			for l := len(w) - 1; l > xVal; l-- {
				if Search(w[:l]) {
					xVal = l
				}
			}
		}
		return xVal
	}

	for _, fn := range []func(a, b []int) int{longestCommonPrefix, Optimized} {
		log.Print("3 ?= ", fn([]int{1, 10, 100}, []int{1000}))
		log.Print("0 ?= ", fn([]int{1, 2, 3}, []int{4, 4, 4}))
		log.Print("--")
	}
}
