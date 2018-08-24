package gopher

import (
	"log"
	"slices"
	"strings"
)

// 212h Word Search II
func findWords(board [][]byte, words []string) []string {
	type Trie struct {
		IsWord bool
		Child  [26]*Trie
	}

	Add := func(n *Trie, w string) {
		for i := 0; i < len(w); i++ {
			c := n.Child[w[i]-'a']
			if c == nil {
				c = &Trie{}
				n.Child[w[i]-'a'] = c
			}
			n = c
		}
		n.IsWord = true
	}

	Search := func(n *Trie, w string) (isPrefix, isWord bool) {
		for i := 0; i < len(w); i++ {
			c := n.Child[w[i]-'a']
			if c == nil {
				return false, false
			}
			n = c
		}
		return true, n.IsWord
	}

	var Remove func(*Trie, string) bool
	Remove = func(n *Trie, w string) bool {
		if len(w) == 0 {
			n.IsWord = false
			for i := range n.Child {
				if n.Child[i] != nil {
					return false
				}
			}
			return true // empty Node
		}

		c := n.Child[w[0]-'a']
		if Remove(c, w[1:]) {
			n.Child[w[0]-'a'] = nil
			for i := range n.Child {
				if n.Child[i] != nil {
					return false
				}
			}
		}
		return true
	}

	trie := &Trie{}
	for _, w := range words {
		Add(trie, w)
	}

	R := []string{}
	Rows, Cols := len(board), len(board[0])

	var Vis map[[2]int]bool
	Dirs := []int{0, 1, 0, -1, 0}

	var DFS func(r, c int, w string)
	DFS = func(r, c int, w string) {
		p, n := Search(trie, w)
		if !p {
			return // not a Prefix
		}
		if n {
			R = append(R, w)
			Remove(trie, w) // to optimize BackTracking
		}

		for i := range 4 {
			x, y := r+Dirs[i], c+Dirs[i+1]
			if x >= 0 && x < Rows && y >= 0 && y < Cols && !Vis[[2]int{x, y}] {
				Vis[[2]int{x, y}] = true
				DFS(x, y, w+string(board[x][y]))
				Vis[[2]int{x, y}] = false // BackTracking ...
			}
		}
	}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			Vis = map[[2]int]bool{}
			Vis[[2]int{r, c}] = true
			DFS(r, c, string(board[r][c]))
		}
	}

	return R
}

// 440h K-th Smallest in Lexicographical Order
func findKthNumber(n int, k int) int {
	v := 1
	k--

	for k > 0 {
		count := 0

		start, end := v, v+1
		for start <= n {
			count += min(n+1, end) - start
			start *= 10
			end *= 10
		}

		if count <= k {
			k -= count
			v++
		} else {
			k--
			v *= 10
		}
	}

	return v
}

// 1233m Remove Sub-Folders from the Filesystem
func removeSubfolders(folder []string) []string {
	slices.Sort(folder)
	log.Print(" -> ", folder)

	M := map[string]struct{}{}
	prv := ""
	for i := range folder {
		if prv == "" || !strings.HasPrefix(folder[i], prv+"/") {
			M[folder[i]] = struct{}{}
			prv = folder[i]
		}
	}

	R := []string{}
	for k := range M {
		R = append(R, k)
	}
	return R
}

// 2416h Sum of Prefix Score of Strings
func sumPrefixScores(words []string) []int {
	type Trie struct {
		Child [26]*Trie
		Score int
	}

	t := &Trie{}
	Insert := func(w string) {
		n := t
		for i := 0; i < len(w); i++ {
			c := n.Child[w[i]-'a']
			if c == nil {
				c = &Trie{}
				n.Child[w[i]-'a'] = c
			}
			n = c
			n.Score++
		}
	}
	Search := func(w string) int {
		score := 0
		n := t
		for i := 0; i < len(w); i++ {
			c := n.Child[w[i]-'a']
			if c == nil {
				return 0
			}
			n = c
			score += n.Score
		}
		return score
	}

	for _, w := range words {
		Insert(w)
	}

	R := []int{}
	for _, w := range words {
		R = append(R, Search(w))
	}
	return R
}

// 3043m Find the Length of the Longest Common Prefix
func longestCommonPrefix(arr1, arr2 []int) int {
	T := map[int]int{} // Trie

	for _, n := range arr1 {
		l := 0
		for x := n; x > 0; x /= 10 {
			l++
		}
		for n > 0 {
			T[n] = l
			n /= 10
			l--
		}
	}

	log.Print(" -> ", T)

	xVal := 0
	for _, n := range arr2 {
		for n > 0 && T[n] == 0 {
			n /= 10
		}
		xVal = max(xVal, T[n])
	}
	return xVal
}
