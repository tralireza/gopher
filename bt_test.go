package gopher

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

// 106m Construct Binary Tree from Inorder and Postorder Traversal
func Test106(t *testing.T) {
	log.Print(" ?= ", buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

func Test297(t *testing.T) {
	type T = TreeNode

	o := NewCode297()

	root := &T{0,
		&T{1, nil, &T{3, &T{Val: 5}, &T{Val: 6}}},
		&T{2, &T{5, &T{Val: 7}, nil}, &T{6, nil, &T{Val: 8}}}}
	treeStr := o.serialize(root)
	rNew := o.deserialize(treeStr)

	var Check func(t1, t2 *TreeNode) bool
	Check = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}
		return Check(t1.Left, t2.Left) && Check(t1.Right, t2.Right)
	}

	log.Printf(":: %q -> %t", treeStr, Check(root, rNew))

	var Draw func(*TreeNode, string, bool)
	Draw = func(n *TreeNode, indent string, lastOne bool) {
		if n == nil {
			return
		}

		fmt.Print(indent)
		if lastOne {
			if indent == "" {
				fmt.Print("*--")
			} else {
				fmt.Print("R--")
			}
			indent += "   "
		} else {
			fmt.Print("L--")
			indent += "|  "
		}
		fmt.Println(n.Val)
		Draw(n.Left, indent, false)
		Draw(n.Right, indent, true)
	}

	Draw(root, "", true)
	log.Print("---")
	Draw(rNew, "", true)
}

func Test352(t *testing.T) {
	o := NewSummaryRanges()
	for _, v := range []int{3, 7, 1, 6, 2} {
		o.AddNum(v)
	}

	if !reflect.DeepEqual([][]int{{1, 3}, {6, 7}}, o.GetIntervals()) {
		t.FailNow()
	}

	var Draw func(*TreeNode, string, bool)
	Draw = func(n *TreeNode, indent string, lastOne bool) {
		if n != nil {
			fmt.Print(indent)
			if lastOne {
				fmt.Print("R-")
				indent += "  "
			} else {
				fmt.Print("L-")
				indent += "| "
			}

			fmt.Println(n.Val)
			Draw(n.Left, indent, false)
			Draw(n.Right, indent, true)
		}
	}

	Draw(o.bstVals, "", true)
}

func Test897(t *testing.T) {
	type N = TreeNode

	var Check func(r1, r2 *N) bool
	Check = func(r1, r2 *N) bool {
		if r1 == nil && r2 == nil {
			return true
		}
		if r1 == nil || r2 == nil {
			return false
		}
		return r1.Val == r2.Val && Check(r1.Left, r2.Left) && Check(r1.Right, r2.Right)
	}

	for _, c := range []struct {
		rst, root *TreeNode
	}{
		{&N{1, nil, &N{5, nil, &N{Val: 7}}}, &N{5, &N{Val: 1}, &N{Val: 7}}},
		{
			&N{1, nil, &N{2, nil, &N{3, nil, &N{4, nil, &N{5, nil, &N{6, nil, &N{7, nil, &N{8, nil, &N{Val: 9}}}}}}}}},
			&N{5,
				&N{3,
					&N{2, &N{Val: 1}, nil},
					&N{Val: 4},
				},
				&N{6, nil,
					&N{8, &N{Val: 7}, &N{Val: 9}},
				},
			},
		},
	} {
		log.Print("*")
		if !Check(c.rst, increasingBST(c.root)) {
			t.FailNow()
		}
	}
}

func Test993(t *testing.T) {
	type T = TreeNode

	for _, c := range []struct {
		rst  bool
		root *TreeNode
		x, y int
	}{
		{false, &T{1, &T{2, &T{Val: 4}, nil}, &T{Val: 3}}, 4, 3},
		{true, &T{1, &T{2, nil, &T{Val: 4}}, &T{3, nil, &T{Val: 5}}}, 5, 4},
		{false, &T{1, &T{2, nil, &T{Val: 4}}, &T{Val: 3}}, 2, 3},
	} {
		log.Print("*")
		if c.rst != isCousins(c.root, c.x, c.y) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

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
	// 1 <= distance <= 10

	Refined := func(root *TreeNode, distance int) int {
		G := map[*TreeNode][]*TreeNode{} // Graph
		L := map[*TreeNode]struct{}{}    // Leaves

		var Walk func(*TreeNode)
		Walk = func(n *TreeNode) {
			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					G[n] = append(G[n], c)
					G[c] = append(G[c], n)
					Walk(c)
				}
			}
			if n.Left == nil && n.Right == nil {
				L[n] = struct{}{}
			}
		}

		Walk(root)

		r := 0

		// BFS for Leaves
		for l := range L {
			Q := []*TreeNode{l}
			Vis := map[*TreeNode]struct{}{}
			var v *TreeNode
			d := 0

			for len(Q) > 0 && d < distance {
				d++
				for range len(Q) {
					v, Q = Q[0], Q[1:]
					Vis[v] = struct{}{}
					for _, u := range G[v] {
						if _, ok := Vis[u]; !ok {
							if _, ok := L[u]; ok {
								r++
							}
							Q = append(Q, u)
						}
					}
				}
			}
		}

		return r / 2
	}

	Direct := func(root *TreeNode, distance int) int {
		R := 0

		var Distance func(*TreeNode) [11]int
		Distance = func(n *TreeNode) [11]int {
			if n == nil {
				return [11]int{}
			}
			if n.Left == nil && n.Right == nil {
				v := [11]int{}
				v[0] = 1 // n: Leaf -> is at distance 0 from 1 Leaf
				return v
			}

			lD := Distance(n.Left)
			rD := Distance(n.Right)

			v := [11]int{}
			for i := 0; i < 10; i++ {
				v[i+1] = rD[i] + lD[i]
			}

			if n.Left != nil && n.Right != nil {
				for l := range distance + 1 {
					for r := range distance + 1 {
						if l+r+2 <= distance {
							R += lD[l] * rD[r]
						}
					}
				}
			}
			return v
		}

		Distance(root)
		return R
	}

	type T = TreeNode

	for _, f := range []func(*TreeNode, int) int{countPairs, Refined, Direct} {
		log.Print("1 ?= ", f(&T{1, &T{2, nil, &T{Val: 4}}, &T{Val: 3}}, 3))
		log.Print("2 ?= ", f(&T{1, &T{2, &T{Val: 4}, &T{Val: 5}}, &T{3, &T{Val: 6}, &T{Val: 7}}}, 3))
		log.Print("--")
	}
}

// 2096m Step-By-Step Directions From a Binary Tree Node to Another
func Test2096(t *testing.T) {
	type T = TreeNode

	log.Print("UURL ?= ", getDirections(&T{5, &T{1, &T{Val: 3}, nil}, &T{2, &T{Val: 6}, &T{Val: 4}}}, 3, 6))
	log.Print("L ?= ", getDirections(&T{2, &T{Val: 1}, nil}, 2, 1))
}
