package gopher

import (
	"fmt"
	"iter"
	"log"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

// 212h Word Search II
func Test212(t *testing.T) {
	log.Printf(`["oath" "eat"] ?= %q`, findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
	log.Printf(`[] ?= %q`, findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcd"}))
}

func (o Trie336) String() string {
	bfr := slices.Repeat([]byte{'-'}, 26)
	for i, c := range o.Child {
		if c != nil {
			bfr[i] = 'a' + byte(i)
		}
	}
	return fmt.Sprintf("{%s %v}", string(bfr), o.I)
}

func Keys336[Map ~map[[2]int]V, V any](m Map) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for k := range m {
			if !yield([]int{k[0], k[1]}) {
				return
			}
		}
	}
}

func Test336(t *testing.T) {
	Hashing := func(words []string) [][]int {
		H := map[string]int{}
		for j, w := range words {
			bfr := []byte(w)
			slices.Reverse(bfr)
			H[string(bfr)] = j
		}

		IsPalindrome := func(w string) bool {
			for l, r := 0, len(w)-1; l < r; l, r = l+1, r-1 {
				if w[l] != w[r] {
					return false
				}
			}
			return true
		}

		M := map[[2]int]struct{}{}
		for i, word := range words {
			for p := 0; p <= len(word); p++ {
				w := word[:p]
				if j, ok := H[w]; ok && j != i && IsPalindrome(word[p:]) {
					M[[2]int{i, j}] = struct{}{}
				}

				v := word[len(word)-p:]
				if j, ok := H[v]; ok && j != i && IsPalindrome(word[:len(word)-p]) {
					M[[2]int{j, i}] = struct{}{}
				}
			}
		}

		log.Print("-> ", H)

		return slices.Collect(Keys336(M))
	}

	for _, c := range []struct {
		rst   [][]int
		words []string
	}{
		{[][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}}, []string{"abcd", "dcba", "lls", "s", "sssll"}},
		{[][]int{{0, 1}, {1, 0}}, []string{"bat", "tab", "cat"}},
		{[][]int{{0, 1}, {1, 0}}, []string{"a", ""}},

		{[][]int{{0, 3}, {2, 3}, {3, 0}, {3, 2}}, []string{"a", "abc", "aba", ""}},
		{[][]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}}, []string{"a", "aa", "aaa"}},
	} {
		log.Printf("* %q", c.words)
		if !reflect.DeepEqual(palindromePairs(c.words), c.rst) {
			t.FailNow()
		}

		log.Print(":: ", Hashing(c.words))
	}
}

func Test440(t *testing.T) {
	for _, c := range []struct {
		rst, n, k int
	}{
		{10, 13, 2},
		{1, 1, 1},
		{104, 127, 7},

		{288990744, 719885387, 209989719}, // TLE 42/69
	} {
		log.Print("* ", c.n, c.k)
		if c.rst != findKthNumber(c.n, c.k) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
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
