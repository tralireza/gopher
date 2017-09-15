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

	type T = TreeNode

	tree := &T{1, &T{2, &T{Val: 4}, &T{Val: 5}}, &T{3, &T{Val: 6}, &T{Val: 7}}}
	Draw(tree)
	log.Print(" -> ")

	for _, tree := range delNodes(tree, []int{3, 5}) {
		fmt.Println(":: ")
		Draw(tree)
	}
}

// 2096m Step-By-Step Directions From a Binary Tree Node to Another
func Test2096(t *testing.T) {
	type T = TreeNode

	log.Print("UURL ?= ", getDirections(&T{5, &T{1, &T{Val: 3}, nil}, &T{2, &T{Val: 6}, &T{Val: 4}}}, 3, 6))
	log.Print("L ?= ", getDirections(&T{2, &T{Val: 1}, nil}, 2, 1))
}
