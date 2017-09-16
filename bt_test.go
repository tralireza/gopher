package gopher

import (
	"fmt"
	"log"
	"testing"
)

// 1110m Delete Nodes And Return Forest
func Test1110(t *testing.T) {
	Draw := func(n *TreeNode) {
		Q := []*TreeNode{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				fmt.Printf("{%d}", n.Val)
				if n.Left != nil {
					Q = append(Q, n.Left)
				}
				if n.Right != nil {
					Q = append(Q, n.Right)
				}
			}
			fmt.Print("\n")
		}
	}

	Iterative := func(root *TreeNode, to_delete []int) []*TreeNode {
		D := make([]bool, 1000+1)
		for _, n := range to_delete {
			D[n] = true
		}

		F := []*TreeNode{}

		Q := []*TreeNode{root}
		var n *TreeNode
		for len(Q) > 0 {
			n, Q = Q[0], Q[1:]

			if n.Left != nil {
				Q = append(Q, n.Left)
				if D[n.Left.Val] {
					n.Left = nil
				}
			}
			if n.Right != nil {
				Q = append(Q, n.Right)
				if D[n.Right.Val] {
					n.Right = nil
				}
			}

			if D[n.Val] {
				if n.Left != nil {
					F = append(F, n.Left)
				}
				if n.Right != nil {
					F = append(F, n.Right)
				}
			}
		}

		if !D[root.Val] {
			F = append(F, root)
		}

		return F
	}

	type T = TreeNode

	for _, f := range []func(*TreeNode, []int) []*TreeNode{delNodes, Iterative} {
		tree := &T{1, &T{2, &T{Val: 4}, &T{Val: 5}}, &T{3, &T{Val: 6}, &T{Val: 7}}}
		Draw(tree)
		log.Print(" -> ")

		for _, tree := range f(tree, []int{3, 5}) {
			fmt.Println(":: ")
			Draw(tree)
		}

		log.Print("--")
	}
}

// 1530m Number of Good Leaf Nodes Pairs
func Test1530(t *testing.T) {
	type T = TreeNode

	log.Print("1 ?= ", countPairs(&T{1, &T{2, nil, &T{Val: 4}}, &T{Val: 3}}, 3))
	log.Print("2 ?= ", countPairs(&T{1, &T{2, &T{Val: 4}, &T{Val: 5}}, &T{3, &T{Val: 6}, &T{Val: 7}}}, 3))
}

// 2096m Step-By-Step Directions From a Binary Tree Node to Another
func Test2096(t *testing.T) {
	type T = TreeNode

	log.Print("UURL ?= ", getDirections(&T{5, &T{1, &T{Val: 3}, nil}, &T{2, &T{Val: 6}, &T{Val: 4}}}, 3, 6))
	log.Print("L ?= ", getDirections(&T{2, &T{Val: 1}, nil}, 2, 1))
}
