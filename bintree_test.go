package gopher

import (
	"log"
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

	type T = TreeNode

	for _, t := range []*T{
		&T{5, &T{4, &T{Val: 1}, &T{Val: 10}}, &T{9, nil, &T{Val: 7}}},
		&T{3, &T{Val: 1}, &T{Val: 2}},
	} {
		log.Printf("%v -> %v", Pack(t), Pack(replaceValueInTree(t)))
	}
}
