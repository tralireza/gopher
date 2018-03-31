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

	tree := &T{1, &T{4, nil, &T{2, &T{Val: 1}, nil}}, &T{4, &T{2, &T{Val: 6}, &T{8, &T{Val: 1}, &T{Val: 3}}}, nil}}

	for _, fn := range []func(*ListNode, *TreeNode) bool{isSubPath, Iterative} {
		log.Print("true ?= ", fn(&L{4, &L{2, &L{Val: 8}}}, tree))
		log.Print("true ?= ", fn(&L{1, &L{4, &L{2, &L{Val: 6}}}}, tree))
		log.Print("false ?= ", fn(&L{1, &L{4, &L{2, &L{6, &L{Val: 8}}}}}, tree))
		log.Print("--")
	}
}
