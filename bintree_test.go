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
	type L = ListNode
	type T = TreeNode

	tree := &T{1, &T{4, nil, &T{2, &T{Val: 1}, nil}}, &T{4, &T{2, &T{Val: 6}, &T{8, &T{Val: 1}, &T{Val: 3}}}, nil}}

	log.Print("true ?= ", isSubPath(&L{4, &L{2, &L{Val: 8}}}, tree))
	log.Print("true ?= ", isSubPath(&L{1, &L{4, &L{2, &L{Val: 6}}}}, tree))
	log.Print("false ?= ", isSubPath(&L{1, &L{4, &L{2, &L{6, &L{Val: 8}}}}}, tree))
}
