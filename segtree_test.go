package gopher

import (
	"fmt"
	"log"
	"testing"
)

// 307m Range Sum Query - Mutable
func Test307(t *testing.T) {
	Draw := func(n *SNode307) {
		Q := []*SNode307{n}
		for len(Q) > 0 {
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				l, r := '/', '/'
				if n.Left != nil {
					l = '*'
					Q = append(Q, n.Left)
				}
				if n.Right != nil {
					r = '*'
					Q = append(Q, n.Right)
				}
				fmt.Printf("{%c [%d:%d] %d %c}", l, n.Start, n.End, n.Sum, r)
			}
			fmt.Print("\n")
		}
	}

	o := Constructor307([]int{1, 3, 5})
	Draw(o.Tree)

	log.Print("9 ?= ", o.SumRange(0, 2))
	o.Update(1, 2)
	log.Print("8 ?= ", o.SumRange(0, 2))
	log.Print("7 ?= ", o.SumRange(1, 2))
}
