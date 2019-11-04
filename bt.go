package gopher

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 106m Construct Binary Tree from Inorder and Postorder Traversal
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	r := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for r.Val != inorder[i] {
		i++
	}
	r.Left = buildTree(inorder[:i], postorder[:i])
	r.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return r
}

// 297h Serialize and Deserialize Binary Tree
type Code297 struct{}

func NewCode297() Code297 { return Code297{} }

func (o *Code297) serialize(root *TreeNode) string {
	T := []string{}

	var preOrder func(*TreeNode)
	preOrder = func(n *TreeNode) {
		if n == nil {
			T = append(T, "*")
		} else {
			T = append(T, strconv.Itoa(n.Val))
			preOrder(n.Left)
			preOrder(n.Right)
		}
	}

	preOrder(root)
	return strings.Join(T, "|")
}

func (o *Code297) deserialize(data string) *TreeNode {
	T, t := strings.Split(data, "|"), ""

	var preOrder func() *TreeNode
	preOrder = func() *TreeNode {
		t, T = T[0], T[1:]
		if t == "*" {
			return nil
		}

		v, _ := strconv.Atoi(t)
		n := &TreeNode{Val: v}
		n.Left, n.Right = preOrder(), preOrder()
		return n
	}

	return preOrder()
}

// 352h Data Stream as Disjoint Intervals
type SummaryRanges struct {
	bstVals *TreeNode
}

func NewSummaryRanges() SummaryRanges { return SummaryRanges{} }

func (o *SummaryRanges) AddNum(value int) { o.bstVals = o.bstVals.Insert352(value) }
func (o *SummaryRanges) GetIntervals() [][]int {
	Vs := [][]int{}
	if o.bstVals == nil {
		return Vs
	}

	left, right := -1, -1
	fnVisit := func(n *TreeNode) {
		if left == -1 {
			left, right = n.Val, n.Val
		} else if right+1 == n.Val {
			right = n.Val
		} else {
			Vs = append(Vs, []int{left, right})
			left, right = n.Val, n.Val
		}
	}

	o.bstVals.InOrder352(fnVisit)
	Vs = append(Vs, []int{left, right})

	return Vs
}

func (o *TreeNode) InOrder352(fnVisit func(*TreeNode)) {
	if o.Left != nil {
		o.Left.InOrder352(fnVisit)
	}

	fnVisit(o)

	if o.Right != nil {
		o.Right.InOrder352(fnVisit)
	}
}

func (o *TreeNode) Insert352(v int) *TreeNode {
	if o == nil {
		return &TreeNode{Val: v}
	}

	if v < o.Val {
		o.Left = o.Left.Insert352(v)
	} else if v > o.Val {
		o.Right = o.Right.Insert352(v)
	}
	return o
}

// 1110m Delete Nodes And Return Forest
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	// 1 <= n.Val, length(to_delete) <= 1000
	D := make([]bool, 1000+1)
	for _, n := range to_delete {
		D[n] = true
	}

	F := []*TreeNode{} // Forest

	var postOrder func(*TreeNode) *TreeNode
	postOrder = func(n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}

		n.Left = postOrder(n.Left)
		n.Right = postOrder(n.Right)

		if D[n.Val] {
			if n.Left != nil {
				F = append(F, n.Left)
			}
			if n.Right != nil {
				F = append(F, n.Right)
			}
			return nil
		}

		return n
	}

	root = postOrder(root)
	if root != nil {
		F = append(F, root)
	}

	return F
}

// 1530m Number of Good Leaf Nodes Pairs
func countPairs(root *TreeNode, distance int) int {
	// 1 <= #Nodes <= 2^10
	// 1 <= Node.Val <= 100

	// (re-)Label Nodes
	var L func(*TreeNode, int) int
	L = func(n *TreeNode, l int) int {
		if n == nil {
			return l
		}

		n.Val = l
		l++
		l = L(n.Left, l)
		return L(n.Right, l)
	}

	n := L(root, 0)

	// Tree -> Graph
	G := make([][]int, n)

	var W func(*TreeNode)
	W = func(n *TreeNode) {
		for _, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				G[n.Val] = append(G[n.Val], c.Val)
				G[c.Val] = append(G[c.Val], n.Val)
				W(c)
			}
		}
	}

	W(root)

	log.Print(G)

	r := 0

	for n := 1; n < len(G); n++ { // Root: G[0]
		if len(G[n]) == 1 { // Leaf: run DFS
			Vis := make([]bool, len(G))
			Q := []int{n}
			d := 0
			var v int
			for len(Q) > 0 && d < distance {
				d++
				for range len(Q) {
					v, Q = Q[0], Q[1:]
					Vis[v] = true
					for _, u := range G[v] {
						if !Vis[u] {
							if len(G[u]) == 1 && u != 0 { // Root: 0
								r++
							}
							Q = append(Q, u)
						}
					}
				}
			}
		}
	}

	return r / 2
}

// 2096m Step-By-Step Directions From a Binary Tree Node to Another
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var lCA func(*TreeNode) *TreeNode // [Lowest] Common-Ancestor
	lCA = func(n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}
		if n.Val == startValue || n.Val == destValue {
			return n
		}

		l, r := lCA(n.Left), lCA(n.Right)
		if l != nil && r != nil {
			return n
		}
		if l != nil {
			return l
		}
		return r
	}

	lca := lCA(root)

	var p bytes.Buffer
	var BackTrack func(*TreeNode, int) bool
	BackTrack = func(n *TreeNode, tVal int) bool {
		log.Print(n, tVal, p)
		if n == nil {
			return false
		}
		if n.Val == tVal {
			return true
		}

		p.WriteByte('L')
		if BackTrack(n.Left, tVal) {
			return true
		}
		p.Truncate(p.Len() - 1)

		p.WriteByte('R')
		if BackTrack(n.Right, tVal) {
			return true
		}
		p.Truncate(p.Len() - 1)

		return false
	}

	p = bytes.Buffer{}
	BackTrack(lca, startValue)
	pStart := p.String()
	log.Printf("+++ lCA -> Start: %q", pStart)

	p.Truncate(0)
	BackTrack(lca, destValue)
	pDest := p.String()
	log.Printf("+++ lCA -> Dest: %q", pDest)

	return strings.Repeat("U", len(pStart)) + pDest
}
