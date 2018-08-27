package gopher

import (
	"log"
	"strconv"
	"strings"
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
	// Trie (with array for Children)
	arrayTrie := func(folder []string) []string {
		type Trie struct {
			Child  [26 + 1]*Trie // [a..z/]
			IsNode bool
		}

		Insert := func(t *Trie, w string) {
			for i := 0; i < len(w); i++ {
				c, index := t.Child[26], 26
				if 'a' <= w[i] && w[i] <= 'z' {
					c = t.Child[w[i]-'a']
					index = int(w[i] - 'a')
				}
				if c == nil {
					c = &Trie{}
					t.Child[index] = c
				}
				t = c
			}
			t.IsNode = true
		}

		HasPrefix := func(t *Trie, w string) bool {
			for i := 0; i < len(w); i++ {
				if w[i] != '/' {
					t = t.Child[w[i]-'a']
				} else {
					t = t.Child[26]
				}

				if t == nil {
					return false
				}
				if t.IsNode && i < len(w)-1 && w[i+1] == '/' {
					return true
				}
			}
			return false
		}

		var Dictionary func(t *Trie) []string
		Dictionary = func(t *Trie) []string {
			W := []string{}
			for i, l := range []byte("abcdefghijklmnopqrstuvwxyz/") {
				if t.Child[i] != nil {
					for _, w := range Dictionary(t.Child[i]) {
						W = append(W, string(l)+w)
					}
					if t.Child[i].IsNode {
						W = append(W, string(l))
					}
				}
			}
			return W
		}

		trie := &Trie{}
		for _, f := range folder {
			Insert(trie, f)
		}
		log.Printf(" -> Dict :: %q", Dictionary(trie))

		R := []string{}
		for _, f := range folder {
			if !HasPrefix(trie, f) {
				R = append(R, f)
			}
		}
		return R
	}

	// Trie (with maps for Children)
	mapTrie := func(folder []string) []string {
		type Trie struct {
			Child  map[string]*Trie
			IsNode bool
		}

		Insert := func(t *Trie, w string) {
			for _, f := range strings.Split(w[1:], "/") {
				c := t.Child[f]
				if c == nil {
					c = &Trie{Child: map[string]*Trie{}}
					t.Child[f] = c
				}
				t = c
			}
			t.IsNode = true
		}

		HasPrefix := func(t *Trie, w string) bool {
			fs := strings.Split(w[1:], "/")
			for i, f := range fs {
				t = t.Child[f]

				if t == nil {
					return false
				}
				if t.IsNode && i < len(fs)-1 {
					return true
				}
			}
			return false
		}

		trie := &Trie{Child: map[string]*Trie{}}
		for _, f := range folder {
			Insert(trie, f)
		}

		R := []string{}
		for _, f := range folder {
			if !HasPrefix(trie, f) {
				R = append(R, f)
			}
		}
		return R
	}

	for _, fn := range []func([]string) []string{removeSubfolders, arrayTrie, mapTrie} {
		log.Printf(`["/a" "/c/d" "/c/f"] ?= %q`, fn([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
		log.Printf(`["/a"] ?= %q`, fn([]string{"/a", "/a/b/c", "/a/b/d"}))
		log.Printf(`["/a/b/c" "/a/b/ca" "/a/b/d"] ?= %q`, fn([]string{"/a/b/c", "/a/b/ca", "/a/b/d"}))
		log.Print("--")
	}
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
