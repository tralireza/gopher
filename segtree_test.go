package gopher

import (
	"fmt"
	"log"
	"testing"
)

// 307m Range Sum Query - Mutable
type FenwickSum307 struct {
	Tree []int
}

func NewFenwickSum307(arr []int) FenwickSum307 {
	o := FenwickSum307{
		make([]int, len(arr)),
	}

	for i := 0; i < len(o.Tree); i++ {
		o.Update(i, arr[i])
	}
	return o
}

func (o *FenwickSum307) Update(i, delta int) {
	for ; i < len(o.Tree); i |= i + 1 {
		o.Tree[i] += delta
	}
}

func (o *FenwickSum307) Sum(i int) int {
	v := 0
	for i >= 0 {
		v += o.Tree[i]
		i &= i + 1
		i--
	}
	return v
}

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

	fws := NewFenwickSum307([]int{1, 3, 5})
	log.Print(" -> FenwickTree Sum (0..2) :: ", fws.Sum(2))
	log.Print(" -> FenwickTree Sum (1..2) :: ", fws.Sum(2)-fws.Sum(0))
	fws.Update(1, -1) // N[1] :: 3 --[delta: -1]-> 2
	log.Print(" -> FenwickTree Sum (1..2) :: ", fws.Sum(2)-fws.Sum(0))

	o := Constructor307([]int{1, 3, 5})
	Draw(o.Tree)

	log.Print("9 ?= ", o.SumRange(0, 2))
	o.Update(1, 2)
	log.Print("8 ?= ", o.SumRange(0, 2))
	log.Print("7 ?= ", o.SumRange(1, 2))
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
