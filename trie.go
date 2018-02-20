package gopher

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
			DFS(r, c, string(board[r][c]))
		}
	}

	return R
}
