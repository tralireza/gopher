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
