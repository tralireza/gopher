package gopher

import (
	"bytes"
	"log"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
	log.Printf("+++ lca -> start: %s", pStart)

	p.Truncate(0)
	BackTrack(lca, destValue)
	pDest := p.String()
	log.Printf("+++ lca -> dest: %s", pDest)

	return strings.Repeat("U", len(pStart)) + pDest
}
