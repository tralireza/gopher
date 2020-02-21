package gopher

import (
	"log"
	"maps"
	"slices"
)

// 307m Range Sum Query - Mutable
type NumArray struct {
	Tree *SNode307
}

type SNode307 struct { // Segment Tree :: Node
	Start, End  int
	Left, Right *SNode307
	Sum         int
}

func Constructor307(nums []int) NumArray {
	var Build func(l, r int) *SNode307
	Build = func(l, r int) *SNode307 {
		if l > r {
			return nil
		}

		n := &SNode307{Start: l, End: r}
		if l == r {
			n.Sum = nums[l]
		} else {
			m := l + (r-l)>>1
			n.Left = Build(l, m)
			n.Right = Build(m+1, r)
			n.Sum = n.Left.Sum + n.Right.Sum
		}
		return n
	}

	return NumArray{Tree: Build(0, len(nums)-1)}
}

func (o *NumArray) Update(index, v int) {
	var Adjust func(*SNode307)
	Adjust = func(n *SNode307) {
		if n.Start == n.End {
			n.Sum = v
			return
		}

		m := n.Start + (n.End-n.Start)>>1
		if index <= m {
			Adjust(n.Left)
		} else {
			Adjust(n.Right)
		}
		n.Sum = n.Left.Sum + n.Right.Sum
	}

	Adjust(o.Tree)
}

func (o *NumArray) SumRange(left, right int) int {
	var Sum func(n *SNode307, l, r int) int
	Sum = func(n *SNode307, l, r int) int {
		if n.Start == l && n.End == r {
			return n.Sum
		}

		m := n.Start + (n.End-n.Start)>>1
		if l >= m+1 {
			return Sum(n.Right, l, r)
		} else if r <= m {
			return Sum(n.Left, l, r)
		}
		return Sum(n.Left, l, m) + Sum(n.Right, m+1, r)
	}

	return Sum(o.Tree, left, right)
}

// 715h Range Module
type RangeModule struct {
	rtSeg *STNode715
}

func NewRangeModule() RangeModule {
	return RangeModule{
		&STNode715{
			nVal:  false,
			left:  1,        // [left ...
			right: int(1e9), // ... right]
		},
	}
}

type STNode715 struct {
	nVal        bool
	left, right int // [left ... right)

	lNode, rNode *STNode715
}

func (o *STNode715) Update(left, right int, nVal bool) bool {
	if left <= o.left && o.right <= right {
		o.nVal = nVal
		o.lNode, o.rNode = nil, nil
		return o.nVal
	}

	if o.left >= right || left >= o.right {
		return o.nVal
	}

	if o.lNode == nil && o.rNode == nil {
		m := o.left + (o.right-o.left)>>1
		o.lNode = &STNode715{
			nVal:  o.nVal,
			left:  o.left,
			right: m,
		}
		o.rNode = &STNode715{
			nVal:  o.nVal,
			left:  m,
			right: o.right,
		}
	}

	l, r := o.lNode.Update(left, right, nVal), o.rNode.Update(left, right, nVal)
	o.nVal = l && r
	return o.nVal
}
func (o *STNode715) Query(left, right int) bool {
	if left <= o.left && o.right < right {
		return o.nVal
	}

	if right <= o.left || o.right <= left {
		return true
	}

	l, r := o.nVal, o.nVal
	if o.lNode != nil && o.rNode != nil {
		l = o.lNode.Query(left, right)
		r = o.rNode.Query(left, right)
	}
	return l && r
}

func (o *RangeModule) AddRange(left, right int) {
	o.rtSeg.Update(left, right, true)
}
func (o *RangeModule) RemoveRange(left, right int) {
	o.rtSeg.Update(left, right, false)
}
func (o *RangeModule) QueryRange(left, right int) bool {
	return o.rtSeg.Query(left, right)
}

// 731m My Calendar II
type MyCalendarTwo struct {
	Mem map[int]int
}

func Constructor731() MyCalendarTwo {
	return MyCalendarTwo{map[int]int{}}
}

func (o *MyCalendarTwo) Book(start, end int) bool {
	Mem := o.Mem

	Mem[start] += 1
	Mem[end] -= 1

	K := []int{}
	for k := range maps.Keys(o.Mem) {
		K = append(K, k)
	}
	slices.Sort(K)

	log.Print(" -> ", K)

	xVal, e := 0, 0
	for _, k := range K {
		e += Mem[k]
		xVal = max(e, xVal)
	}

	if xVal <= 2 {
		return true
	}

	Mem[start] -= 1
	Mem[end] += 1

	return false
}
