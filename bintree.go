package gopher

import (
	"container/heap"
	"log"
	"slices"
)

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

// 951m Flip Equivalent Binary Trees
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}
	if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
		return true
	}
	return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
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

// 2583m Kth Largest Sum in a Binary Tree
type PQ2583 []int64

func (h PQ2583) Len() int           { return len(h) }
func (h PQ2583) Less(i, j int) bool { return h[i] > h[j] }
func (h PQ2583) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ2583) Push(_ any)        {}
func (h *PQ2583) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	S := PQ2583{}

	Q := []*TreeNode{}
	n := root

	Q = append(Q, n)
	for len(Q) > 0 {
		lSum := int64(0)
		for range len(Q) {
			n, Q = Q[0], Q[1:]
			lSum += int64(n.Val)

			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					Q = append(Q, c)
				}
			}
		}
		S = append(S, lSum)
	}

	if S.Len() < k {
		return -1
	}

	heap.Init(&S)
	for range k - 1 {
		heap.Pop(&S)
	}
	return heap.Pop(&S).(int64)
}

// 2641m Cousins in Binary Tree II
func replaceValueInTree(root *TreeNode) *TreeNode {
	lSums := map[int]int{}

	Q := []*TreeNode{root}
	var n *TreeNode

	l := 0
	for len(Q) > 0 {
		lSum := 0
		for range len(Q) {
			n, Q = Q[0], Q[1:]
			lSum += n.Val
			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					Q = append(Q, c)
				}
			}
		}
		lSums[l] = lSum
		l++
	}

	log.Print(" -> lSums :: ", lSums)

	var W func(n *TreeNode, sVal, l int)
	W = func(n *TreeNode, sVal, l int) {
		if n != nil {
			n.Val = lSums[l] - sVal

			v := 0
			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					v += c.Val
				}
			}
			W(n.Left, v, l+1)
			W(n.Right, v, l+1)
		}
	}

	W(root, root.Val, 0)
	return root
}
