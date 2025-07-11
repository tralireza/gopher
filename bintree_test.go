package gopher

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

// 103 Binary Tree Zigzag Level Order Traversal
func Test103(t *testing.T) {
	type T = TreeNode

	log.Print("[[3] [20 9] [15 7]] ?= ", zigzagLevelOrder(&T{3, &T{Val: 9}, &T{20, &T{Val: 15}, &T{Val: 7}}}))
	log.Print("[[1]] ?= ", zigzagLevelOrder(&T{Val: 1}))
	log.Print("[] ?= ", zigzagLevelOrder(nil))

	log.Print("[[1] [3 2] [4 5]] ?= ", zigzagLevelOrder(&T{1, &T{2, &T{Val: 4}, &T{Val: 5}}, &T{Val: 3}}))
}

// 501 Find Mode in Binary Search Tree
func Test501(t *testing.T) {
	type T = TreeNode

	log.Print("[2] ?= ", findMode(&T{1, nil, &T{2, &T{Val: 2}, nil}}))
	log.Print("[0] ?= ", findMode(&T{Val: 0}))
}

// 515m Find Largest Value in Each Tree Row
func Test515(t *testing.T) {
	DFS := func(root *TreeNode) []int {
		R := []int{}

		var Run func(*TreeNode, int)
		Run = func(n *TreeNode, level int) {
			if n == nil {
				return
			}

			if level == len(R) {
				R = append(R, n.Val)
			} else {
				if n.Val > R[level] {
					R[level] = n.Val
				}
			}

			Run(n.Left, level+1)
			Run(n.Right, level+1)
		}

		Run(root, 0)
		return R
	}

	type T = TreeNode

	for _, fn := range []func(*TreeNode) []int{largestValues, DFS} {
		t.Log("[1 3 9] ?= ", fn(&T{1, &T{3, &T{Val: 5}, &T{Val: 3}}, &T{Val: 2, Right: &T{Val: 9}}}))
		t.Log("[1 3] ?= ", fn(&T{1, &T{Val: 2}, &T{Val: 3}}))
		t.Log("--")
	}
}

// 563 Binary Tree Tilt
func Test563(t *testing.T) {
	type T = TreeNode

	log.Print("1 ?= ", findTilt(&T{1, &T{Val: 2}, &T{Val: 3}}))
	log.Print("15 ?= ", findTilt(&T{4, &T{2, &T{Val: 3}, &T{Val: 5}}, &T{Val: 9, Right: &T{Val: 7}}}))
	log.Print("9 ?= ", findTilt(&T{21, &T{7, &T{1, &T{Val: 1}, &T{Val: 1}}, &T{Val: 1}}, &T{14, &T{Val: 2}, &T{Val: 2}}}))
}

// 559 Maximum Depth of N-ary Tree
func Test559(t *testing.T) {
	type T = NTreeNode

	log.Print("3 ?= ", maxDepth(&T{1, []*T{&T{3, []*T{&T{Val: 5}, &T{Val: 6}}}, &T{Val: 2}, &T{Val: 4}}}))
	log.Print("5 ?= ", maxDepth(&T{1,
		[]*T{
			&T{Val: 2},
			&T{3,
				[]*T{
					&T{Val: 6},
					&T{7, []*T{&T{11, []*T{&T{Val: 14}}}}}}},
			&T{4, []*T{&T{8, []*T{&T{Val: 12}}}}},
			&T{5, nil}}}))
}

func Test865(t *testing.T) {
	Draw := func(n *TreeNode) {
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				l, r := '-', '-'
				if n.Left != nil {
					Q = append(Q, n.Left)
					l = '*'
				}
				if n.Right != nil {
					Q = append(Q, n.Right)
					r = '*'
				}
				fmt.Printf("{%c %d %c}", l, n.Val, r)
			}
			fmt.Print("\n")
		}
	}

	Recursive := func(root *TreeNode) *TreeNode {
		var DFS func(int, *TreeNode) (int, *TreeNode)
		DFS = func(d int, n *TreeNode) (int, *TreeNode) {
			if n == nil {
				return d, nil
			}
			if n.Left == nil && n.Right == nil {
				return d, n
			}

			lD, l := DFS(d+1, n.Left)
			rD, r := DFS(d+1, n.Right)

			if l != nil && r != nil {
				if lD > rD {
					return lD, l
				} else if lD < rD {
					return rD, r
				}
				return lD | rD, n
			}
			if l != nil {
				return lD, l
			}
			return rD, r
		}

		_, t := DFS(0, root)
		return t
	}

	type T = TreeNode

	single, leaf, subtree := &T{Val: 1}, &T{Val: 2}, &T{2, &T{Val: 7}, &T{Val: 4}}

	for _, c := range []struct {
		rst  *TreeNode
		tree *TreeNode
	}{
		{subtree, &T{3, &T{5, &T{Val: 6}, subtree}, &T{1, &T{Val: 0}, &T{Val: 8}}}},
		{single, single},
		{leaf, &T{0, &T{1, nil, leaf}, &T{Val: 3}}},
	} {
		Draw(c.tree)
		fmt.Print("--\n")
		for _, f := range []func(*TreeNode) *TreeNode{subtreeWithAllDeepest, Recursive} {
			if !reflect.DeepEqual(c.rst, f(c.tree)) {
				t.FailNow()
			}
			Draw(f(c.tree))
		}
		fmt.Print("==\n")
	}
}

func Test617(t *testing.T) {
	type T = TreeNode

	var Check func(root1, root2 *T) bool
	Check = func(root1, root2 *T) bool {
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}

		return root1.Val == root2.Val &&
			Check(root1.Left, root2.Left) &&
			Check(root1.Right, root2.Right)
	}

	for _, c := range []struct {
		rst, root1, root2 *TreeNode
	}{
		{
			&T{3, &T{4, &T{Val: 5}, &T{Val: 4}}, &T{5, nil, &T{Val: 7}}},
			&T{1, &T{3, &T{Val: 5}, nil}, &T{Val: 2}},
			&T{2, &T{1, nil, &T{Val: 4}}, &T{3, nil, &T{Val: 7}}},
		},
	} {
		log.Print("* ")
		if !Check(c.rst, mergeTrees(c.root1, c.root2)) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test653(t *testing.T) {
	type T = TreeNode
	for _, c := range []struct {
		rst  bool
		root *TreeNode
		k    int
	}{
		{true, &T{5, &T{3, &T{Val: 2}, &T{Val: 4}}, &T{6, nil, &T{Val: 7}}}, 9},
		{false, &T{5, &T{3, &T{Val: 2}, &T{Val: 4}}, &T{6, nil, &T{Val: 7}}}, 28},
	} {
		log.Print("* ", c.root, c.k)
		if c.rst != findTarget(c.root, c.k) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test671(t *testing.T) {
	type T = TreeNode

	for _, c := range []struct {
		rst  int
		root *TreeNode
	}{
		{5, &T{2, &T{Val: 2}, &T{5, &T{Val: 5}, &T{Val: 7}}}},
		{-1, &T{2, &T{Val: 2}, &T{Val: 2}}},
	} {
		log.Print("* ", c.root)
		if c.rst != findSecondMinimumValue(c.root) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

// 889 Construct Binary Tree from Preorder and Postorder Traversal
func Test889(t *testing.T) {
	Recursive := func(preorder []int, postorder []int) *TreeNode {
		pre, post := 0, 0

		var Recover func() *TreeNode
		Recover = func() *TreeNode {
			n := &TreeNode{Val: preorder[pre]}
			pre++

			if n.Val != postorder[post] {
				n.Left = Recover()
			}
			if n.Val != postorder[post] {
				n.Right = Recover()
			}

			post++
			return n
		}

		return Recover()
	}

	Draw := func(n *TreeNode) {
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				l, r := '-', '-'
				if n.Left != nil {
					l = '*'
					Q = append(Q, n.Left)
				}
				if n.Right != nil {
					r = '*'
					Q = append(Q, n.Right)
				}
				fmt.Printf("{%c %d %c}", l, n.Val, r)
			}
			fmt.Print("\n")
		}
	}

	for _, f := range []func([]int, []int) *TreeNode{constructFromPrePost, Recursive} {
		Draw(f([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}))
	}
}

// 951m Flip Equivalent Binary Trees
func Test951(t *testing.T) {
	type T = TreeNode

	log.Print("true ?= ", flipEquiv(&T{1, &T{2, &T{Val: 4}, &T{5, &T{Val: 7}, &T{Val: 8}}}, &T{3, &T{Val: 6}, nil}}, &T{1, &T{3, nil, &T{Val: 6}}, &T{2, &T{Val: 4}, &T{5, &T{Val: 8}, &T{Val: 7}}}}))
	log.Print("true ?= ", flipEquiv(nil, nil))
	log.Print("false ?= ", flipEquiv(nil, &T{Val: 1}))
}

// 1028h Recover a Tree From Preorder Traversal
func Test1028(t *testing.T) {
	Draw := func(n *TreeNode) {
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				l, r := '-', '-'
				if n.Left != nil {
					l = '*'
					Q = append(Q, n.Left)
				}
				if n.Right != nil {
					r = '*'
					Q = append(Q, n.Right)
				}
				fmt.Printf("{%c %d %c}", l, n.Val, r)
			}
			fmt.Print("\n")
		}
	}

	for _, traversal := range []string{
		"1-2--3--4-5--6--7",
		"1-2--3---4-5--6---7",
		"1-401--349---90--88",
	} {
		Draw(recoverFromPreorder(traversal))
		log.Print("--")
	}
}

func Test1123(t *testing.T) {
	Draw := func(n *TreeNode) {
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				l, r := '-', '-'
				if n.Left != nil {
					Q = append(Q, n.Left)
					l = '*'
				}
				if n.Right != nil {
					Q = append(Q, n.Right)
					r = '*'
				}
				fmt.Printf("{%c %d %c}", l, n.Val, r)
			}
			fmt.Print("\n")
		}
	}

	type T = TreeNode
	for _, c := range []struct {
		tree *TreeNode
	}{
		{&T{3, &T{5, &T{Val: 6}, &T{2, &T{Val: 7}, &T{Val: 4}}}, &T{1, &T{Val: 0}, &T{Val: 8}}}},
		{&T{Val: 1}},
		{&T{0, &T{1, nil, &T{Val: 2}}, &T{Val: 3}}},
	} {
		Draw(lcaDeepestLeaves(c.tree))
		fmt.Print("\n")
	}
}

// 1261m Find Elements in a Contaminated Binary Tree
func Test1261(t *testing.T) {
	type T = TreeNode
	o := Constructor1261(&T{-1, &T{-1, &T{Val: -1}, &T{Val: -1}}, &T{Val: -1}})
	for _, tc := range []struct {
		target int
		rst    bool
	}{
		{1, true},
		{3, true},
		{5, false},
	} {
		if tc.rst != o.Find(tc.target) {
			t.Fail()
		}
	}
}

// 1367 Linked List in Binary Tree
func Test1367(t *testing.T) {
	// Knuth-Morris-Pratt KMP: Failure Function
	KMP := func(l *ListNode) ([]int, []int) {
		Vals := []int{l.Val} // ie, Pattern
		uF := []int{0}

		l = l.Next
		for pIndex := 0; l != nil; l = l.Next {
			for pIndex > 0 && l.Val != Vals[pIndex] {
				pIndex = uF[pIndex-1]
			}

			if l.Val == Vals[pIndex] {
				pIndex++
			} else {
				pIndex = 0
			}

			Vals = append(Vals, l.Val)
			uF = append(uF, pIndex)
		}

		return uF, Vals
	}

	KMPSearch := func(haystack []int, needle []int) int {
		uF, k := []int{0}, 0
		for j := 1; j < len(needle); j++ {
			for k > 0 && needle[k] != needle[j] {
				k = uF[k-1]
			}
			if needle[j] == needle[k] {
				k++
			} else {
				k = 0
			}
			uF = append(uF, k)
		}

		log.Print(needle, " -> uF: ", uF)

		k = 0
		for j := range haystack {
			for k > 0 && needle[k] != haystack[j] {
				k = uF[k-1]
			}
			if needle[k] == haystack[j] {
				k++
				if k == len(needle) {
					return j - k + 1
				}
			}
		}
		return -1
	}

	Iterative := func(head *ListNode, root *TreeNode) bool {
		if root == nil {
			return false
		}

		Check := func(l *ListNode, t *TreeNode) bool {
			Q, R := []*TreeNode{t}, []*ListNode{l}

			for len(Q) > 0 && len(R) > 0 {
				t, Q, l, R = Q[len(Q)-1], Q[:len(Q)-1], R[len(R)-1], R[:len(R)-1]
				if t.Val == l.Val {
					l = l.Next
					if l != nil {
						if t.Left != nil {
							Q, R = append(Q, t.Left), append(R, l)
						}
						if t.Right != nil {
							Q, R = append(Q, t.Right), append(R, l)
						}
					} else {
						return true
					}
				}
			}

			return false
		}

		Q := []*TreeNode{root}
		var t *TreeNode

		for len(Q) > 0 {
			t, Q = Q[len(Q)-1], Q[:len(Q)-1]
			if Check(head, t) {
				return true
			}

			for _, c := range []*TreeNode{t.Left, t.Right} {
				if c != nil {
					Q = append(Q, c)
				}
			}
		}
		return false
	}

	type L = ListNode
	type T = TreeNode

	uF, Pattern := KMP(&L{1, &L{2, &L{1, &L{2, &L{Val: 3}}}}})
	log.Print("+++ KMP uF & Pattern -> ", uF, Pattern)
	log.Print("KMP Search :: 2 ?= ", KMPSearch([]int{1, 2, 1, 2, 1, 2, 3}, []int{1, 2, 1, 2, 3}))
	log.Print("--")

	tree := &T{1, &T{4, nil, &T{2, &T{Val: 1}, nil}}, &T{4, &T{2, &T{Val: 6}, &T{8, &T{Val: 1}, &T{Val: 3}}}, nil}}

	for _, fn := range []func(*ListNode, *TreeNode) bool{isSubPath, Iterative} {
		log.Print("true ?= ", fn(&L{4, &L{2, &L{Val: 8}}}, tree))
		log.Print("true ?= ", fn(&L{1, &L{4, &L{2, &L{Val: 6}}}}, tree))
		log.Print("false ?= ", fn(&L{1, &L{4, &L{2, &L{6, &L{Val: 8}}}}}, tree))
		log.Print("--")
	}
}

// 2415m Reverse Odd Levels of Binary Tree
func Test2415(t *testing.T) {
	type T = TreeNode

	preOrder := func(n *TreeNode) []int {
		v := []int{}
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			n, Q = Q[0], Q[1:]
			v = append(v, n.Val)
			if n.Left != nil {
				Q = append(Q, n.Left)
			}
			if n.Right != nil {
				Q = append(Q, n.Right)
			}
		}
		return v
	}

	for _, tree := range []*T{
		&T{2, &T{3, &T{Val: 8}, &T{Val: 13}}, &T{5, &T{Val: 21}, &T{Val: 34}}},
		&T{7, &T{Val: 13}, &T{Val: 11}},
	} {
		log.Print(preOrder(tree), " -> ", preOrder(reverseOddLevels(tree)))
	}
}

// 2471m Minimum Number of Operations to Sort a Binary Tree by Level
func Test2471(t *testing.T) {
	// 1 <= Node.Val <= 10^5
	type T = TreeNode

	log.Print("3 ?= ", minimumOperations(&T{1, &T{4, &T{Val: 7}, &T{Val: 6}}, &T{3, &T{Val: 8, Left: &T{Val: 9}}, &T{Val: 5, Left: &T{Val: 10}}}}))
	log.Print("3 ?= ", minimumOperations(&T{1, &T{3, &T{Val: 7}, &T{Val: 6}}, &T{2, &T{Val: 5}, &T{Val: 4}}}))
}

// 2583m Kth Largest Sum in a Binary Tree
func Test2583(t *testing.T) {
	type T = TreeNode

	log.Print("13 ?= ", kthLargestLevelSum(&T{5, &T{8, &T{2, &T{Val: 4}, &T{Val: 6}}, &T{Val: 1}}, &T{9, &T{Val: 3}, &T{Val: 7}}}, 2))
	log.Print("3 ?= ", kthLargestLevelSum(&T{1, &T{2, &T{Val: 3}, nil}, nil}, 1))
}

// 2641m Cousins in Binary Tree II
func Test2641(t *testing.T) {
	// 1 <= Node.Val <= 10^4

	Pack := func(n *TreeNode) []int {
		R := []int{}
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				if n == nil {
					R = append(R, -1)
					continue
				}
				R = append(R, n.Val)
				Q = append(Q, n.Left)
				Q = append(Q, n.Right)
			}
		}
		return R
	}

	OnePass := func(root *TreeNode) *TreeNode {
		Q := []*TreeNode{root}
		lSum := root.Val

		var n *TreeNode
		for len(Q) > 0 {
			nSum := 0
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				n.Val = lSum - n.Val

				fSum := 0
				if n.Left != nil {
					fSum += n.Left.Val
				}
				if n.Right != nil {
					fSum += n.Right.Val
				}
				nSum += fSum

				if n.Left != nil {
					Q = append(Q, n.Left)
					n.Left.Val = fSum
				}
				if n.Right != nil {
					Q = append(Q, n.Right)
					n.Right.Val = fSum
				}
			}
			lSum = nSum
		}

		return root
	}

	type T = TreeNode

	for _, fn := range []func(*TreeNode) *TreeNode{replaceValueInTree, OnePass} {
		for _, t := range []*T{
			&T{5, &T{4, &T{Val: 1}, &T{Val: 10}}, &T{9, nil, &T{Val: 7}}},
			&T{3, &T{Val: 1}, &T{Val: 2}},
		} {
			log.Printf("%v -> %v", Pack(t), Pack(fn(t)))
		}
		log.Print("--")
	}
}

// 2872h Maximum Number of K-Divisible Components
func Test2872(t *testing.T) {
	log.Print("2 ?= ", maxKDivisibleComponents(5, [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, []int{1, 8, 1, 4, 4}, 6))
	log.Print("3 ?= ", maxKDivisibleComponents(7, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, []int{3, 0, 6, 1, 5, 2, 1}, 3))
}
