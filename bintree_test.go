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
}
