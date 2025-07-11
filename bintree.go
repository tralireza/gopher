package gopher

import (
	"container/heap"
	"container/list"
	"fmt"
	"log"
	"math"
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

// 501 Find Mode in Binary Search Tree
func findMode(root *TreeNode) []int {
	fMap := map[int]int{}

	var Walk func(*TreeNode)
	Walk = func(n *TreeNode) {
		if n != nil {
			Walk(n.Left)
			Walk(n.Right)
			fMap[n.Val]++
		}
	}

	Walk(root)

	log.Print(" -> ", fMap)

	fMax := math.MinInt
	for _, f := range fMap {
		if f > fMax {
			fMax = f
		}
	}

	R := []int{}
	for v, f := range fMap {
		if f == fMax {
			R = append(R, v)
		}
	}

	return R
}

// 515m Find Largest Value in Each Tree Row
func largestValues(root *TreeNode) []int {
	R := []int{}
	if root == nil {
		return R
	}

	Q := list.New()
	Q.PushBack(root)

	for Q.Len() > 0 {
		xVal := math.MinInt

		for range Q.Len() {
			n := Q.Remove(Q.Front()).(*TreeNode)
			if n.Val > xVal {
				xVal = n.Val
			}

			if n.Left != nil {
				Q.PushBack(n.Left)
			}
			if n.Right != nil {
				Q.PushBack(n.Right)
			}
		}

		R = append(R, xVal)
	}

	return R
}

// 559 Maximum Depth of N-ary Tree
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func maxDepth(root *NTreeNode) int {
	depth := 0

	var Walk func(*NTreeNode, int)
	Walk = func(n *NTreeNode, d int) {
		if n != nil {
			depth = max(d, depth)
			for _, cn := range n.Children {
				Walk(cn, d+1)
			}
		}
	}

	Walk(root, 1)

	return depth
}

// 563 Binary Tree Tilt
func findTilt(root *TreeNode) int {
	tilt := 0
	var PostOrder func(*TreeNode) int
	PostOrder = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l, r := PostOrder(n.Left), PostOrder(n.Right)
		if l > r {
			l, r = r, l
		}

		tilt += r - l
		return n.Val + l + r
	}

	PostOrder(root)
	return tilt
}

// 617 Merge Two Binary Trees
func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	n := &TreeNode{Val: root1.Val + root2.Val}
	n.Left = mergeTrees(root1.Left, root2.Left)
	n.Right = mergeTrees(root1.Right, root2.Right)
	return n
}

// 653 Two Sum IV - Input is a BST
func findTarget(root *TreeNode, k int) bool {
	M := map[int]struct{}{}

	var Walk func(*TreeNode) bool
	Walk = func(n *TreeNode) bool {
		if n == nil {
			return false
		}

		if _, ok := M[k-n.Val]; ok {
			return true
		}
		M[n.Val] = struct{}{}
		return Walk(n.Left) || Walk(n.Right)
	}

	return Walk(root)
}

// 671 Second Minimum Node in a Binary Tree
func findSecondMinimumValue(root *TreeNode) int {
	mVal, rVal := root.Val, math.MaxInt

	var Walk func(*TreeNode)
	Walk = func(n *TreeNode) {
		if n == nil {
			return
		}

		if n.Val != mVal && n.Val < rVal {
			rVal = n.Val
		} else {
			Walk(n.Left)
			Walk(n.Right)
		}
	}

	Walk(root)

	if rVal != math.MaxInt {
		return rVal
	}
	return -1
}

// 865m Smallest Subtree
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var Depth func(*TreeNode) int
	Depth = func(n *TreeNode) int {
		if n != nil {
			return 1 + max(Depth(n.Left), Depth(n.Right))
		}
		return 0
	}

	depth := Depth(root)

	var LCA func(int, *TreeNode) *TreeNode
	LCA = func(d int, n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}

		if n.Left == nil && n.Right == nil {
			if d == depth {
				return n
			}
			return nil
		}

		l, r := LCA(d+1, n.Left), LCA(d+1, n.Right)
		if l != nil && r != nil {
			return n
		}
		if l != nil {
			return l
		}
		return r
	}

	return LCA(1, root)
}

// 889 Construct Binary Tree from Preorder and Postorder Traversal
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	switch len(preorder) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: preorder[0]}
	}

	n := &TreeNode{Val: preorder[0]}

	x := 0
	for postorder[x] != preorder[1] {
		x++
	}
	x++

	n.Left = constructFromPrePost(preorder[1:x+1], postorder[0:x])
	n.Right = constructFromPrePost(preorder[x+1:], postorder[x:len(postorder)-1])

	return n
}

// 951m Flip Equivalent Binary Trees
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	fmt.Print(".")
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

// 1028h Recover a Tree From Preorder Traversal
func recoverFromPreorder(traversal string) *TreeNode {
	p := 0

	var Load func(int) *TreeNode
	Load = func(d int) *TreeNode {
		if p >= len(traversal) {
			return nil
		}

		x := 0
		for p+x < len(traversal) && traversal[p+x] == '-' {
			x++
		}
		if x != d {
			return nil
		}

		p += x

		v := 0
		for p < len(traversal) && traversal[p] != '-' {
			v = 10*v + int(traversal[p]-'0')
			p++
		}

		n := &TreeNode{Val: v}
		n.Left = Load(d + 1)
		n.Right = Load(d + 1)

		return n
	}

	return Load(0)
}

// 1123m Lowest Common Ancestor of Deepest Leaves
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var Depth func(*TreeNode) int
	Depth = func(n *TreeNode) int {
		if n != nil {
			return 1 + max(Depth(n.Left), Depth(n.Right))
		}
		return 0
	}

	depth := Depth(root)

	var LCA func(int, *TreeNode) *TreeNode
	LCA = func(d int, n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}

		if n.Left == nil && n.Right == nil {
			if d == depth {
				return n
			}
			return nil
		}

		l, r := LCA(d+1, n.Left), LCA(d+1, n.Right)
		if l != nil && r != nil {
			return n
		}
		if l != nil {
			return l
		}
		return r
	}

	return LCA(1, root)
}

// 1261m Find Elements in a Contaminated Binary Tree
type FindElements struct {
	M map[int]struct{}
}

func Constructor1261(root *TreeNode) FindElements {
	M := map[int]struct{}{0: struct{}{}}

	var Load func(*TreeNode, int)
	Load = func(n *TreeNode, pVal int) {
		for i, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				M[pVal+i+1] = struct{}{}
				Load(c, (pVal+i+1)<<1)
			}
		}
	}

	Load(root, 0)

	log.Print(" -> Recursive DFS :: ", M)

	// Iterative BFS
	Mp := map[int]struct{}{}

	root.Val = 0
	Q := []*TreeNode{root}
	var n *TreeNode
	for len(Q) > 0 {
		n, Q = Q[0], Q[1:]
		Mp[n.Val] = struct{}{}

		for i, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				c.Val = 2*n.Val + i + 1
				Q = append(Q, c)
			}
		}
	}

	log.Print(" -> Iterative BFS :: ", Mp)

	return FindElements{M}
}

func (o *FindElements) Find(target int) bool {
	_, ok := o.M[target]
	return ok
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

// 2415m Reverse Odd Levels of Binary Tree
func reverseOddLevels(root *TreeNode) *TreeNode {
	var Swap func(l, r *TreeNode, level int)
	Swap = func(l, r *TreeNode, level int) {
		if l == nil || r == nil {
			return
		}

		if level&1 == 1 {
			l.Val, r.Val = r.Val, l.Val
		}
		Swap(l.Left, r.Right, level+1)
		Swap(l.Right, r.Left, level+1)
	}

	Swap(root.Left, root.Right, 1)
	return root
}

// 2471m Minimum Number of Operations to Sort a Binary Tree by Level
func minimumOperations(root *TreeNode) int {
	ops := 0

	M := make([]int, 100_000+1)
	Q := []*TreeNode{root}

	for len(Q) > 0 {
		lVals := []int{}
		for i := range Q {
			n := Q[i]

			M[n.Val] = i
			lVals = append(lVals, n.Val)

			if n.Left != nil {
				Q = append(Q, n.Left)
			}
			if n.Right != nil {
				Q = append(Q, n.Right)
			}
		}
		Q = Q[len(lVals):]

		sVals := append([]int{}, lVals...)
		slices.Sort(sVals)

		for i, sVal := range sVals {
			if sVal != lVals[i] {
				ops++

				M[lVals[i]] = M[sVal]
				lVals[M[sVal]] = lVals[i]
			}
		}
	}

	return ops
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

// 2872h Maximum Number of K-Divisible Components
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	cmps := 0

	G := make([][]int, n)
	for _, e := range edges {
		G[e[0]] = append(G[e[0]], e[1])
		G[e[1]] = append(G[e[1]], e[0])
	}

	var Search func(v, p int) int
	Search = func(v, p int) int {
		tSum := 0

		for _, u := range G[v] {
			if u != p {
				tSum += Search(u, v)
				tSum %= k
			}
		}

		tSum += values[v]
		tSum %= k
		if tSum == 0 {
			cmps++
		}

		return tSum
	}

	Search(0, -1)
	return cmps
}
