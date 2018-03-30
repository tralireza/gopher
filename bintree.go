package gopher

import "slices"

// 103 Binary Tree Zigzag Level Order Traversal
func zigzagLevelOrder(root *TreeNode) [][]int {
	Z := [][]int{}

	Q := []*TreeNode{}
	if root != nil {
		Q = append(Q, root)
	}

	Reverse := false
	for len(Q) > 0 {
		L := []int{} // level

		var v *TreeNode
		for range len(Q) { // BFS
			v, Q = Q[0], Q[1:]
			L = append(L, v.Val)
			for _, u := range []*TreeNode{v.Left, v.Right} {
				if u != nil {
					Q = append(Q, u)
				}
			}
		}

		if Reverse {
			slices.Reverse(L)
		}
		Reverse = !Reverse
		Z = append(Z, L)
	}

	return Z
}

// 1367 Linked List in Binary Tree
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	var Check func(*ListNode, *TreeNode) bool
	Check = func(l *ListNode, t *TreeNode) bool {
		if t == nil || l == nil {
			return l == nil
		}

		v := false
		if l.Val == t.Val {
			v = Check(l.Next, t.Left) || Check(l.Next, t.Right)
		}
		return v
	}

	if Check(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
