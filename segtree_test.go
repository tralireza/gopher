package gopher

import (
	"fmt"
	"log"
	"testing"
)

// 307m Range Sum Query - Mutable
/*
(1-based index)
BIT[i] :: [N_(i-1<<LSB(i)+1) ... N_i]
*/
type FenwickSum307 []int

func NewFenwickSum307(arr []int) FenwickSum307 {
	fws := FenwickSum307(make([]int, len(arr)))
	log.Printf(" ++ %T %[1]p %[1]v", fws)
	for i := 0; i < len(fws); i++ {
		fws.Update(i, arr[i])
	}
	return fws
}

func (fws *FenwickSum307) Update(i, delta int) {
	for i < len(*fws) {
		(*fws)[i] += delta
		i |= i + 1
	}
}

func (fws *FenwickSum307) Sum(i int) int {
	v := 0
	for i >= 0 {
		v += (*fws)[i]
		i &= i + 1
		i--
	}
	return v
}

func Test307(t *testing.T) {
	// FenwickTree :: Sum
	fws := NewFenwickSum307([]int{1, 2, 3, 4, 5})
	log.Printf(" -> %T %[1]p %[1]v", fws)
	log.Print(" -> FenwickTree Sum (0..4) :: ", fws.Sum(4))
	log.Print(" -> FenwickTree Sum (3..4) :: ", fws.Sum(4)-fws.Sum(2))
	fws.Update(2, -1) // N[2] :: 3 --[delta: -1]-> 2
	log.Printf(" -> %T %[1]p %[1]v", fws)
	log.Print(" -> FenwickTree Sum (1..2) :: ", fws.Sum(2)-fws.Sum(0))

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

func (o STNode715) String() string {
	nVal := ' '
	if o.nVal {
		nVal = '*'
	}
	return fmt.Sprintf("{%d..%d %c}", o.left, o.right, nVal)
}

func Test715(t *testing.T) {
	// 1 <= left < right <= 10^9

	type Node = STNode715

	var Draw func(*Node, string, bool)
	Draw = func(n *Node, indent string, lastOne bool) {
		if n == nil {
			return
		}

		if lastOne {
			log.Printf("%s!- %v", indent, n)
			indent += "   "
		} else {
			log.Printf("%s+- %v", indent, n)
			indent += "|  "
		}

		Draw(n.lNode, indent, false)
		Draw(n.rNode, indent, true)
	}

	o := NewRangeModule()

	o.AddRange(10, 20)
	o.RemoveRange(14, 16)

	Draw(o.rtSeg, "", true)

	for _, c := range []struct {
		rst         bool
		left, right int
	}{
		{true, 10, 14},
		{false, 13, 15},
		{true, 16, 17},
	} {
		if c.rst != o.QueryRange(c.left, c.right) {
			t.Error()
		}
	}

	log.Print("---")
	o.AddRange(15, 17)
	o.AddRange(2, 6)
	Draw(o.rtSeg, "", true)
}

// 731m My Calendar II
func Test731(t *testing.T) {
	// 0 <= Event(start, end) <= 10^9
	events := [][]int{{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}}
	results := []bool{true, true, true, false, true, true}

	o := Constructor731()
	for i, e := range events {
		log.Printf("%t ?= %t", results[i], o.Book(e[0], e[1]))
	}
}
