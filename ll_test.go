package gopher

import (
	"fmt"
	"log"
	"testing"
)

func llDraw(n *ListNode) {
	for n != nil {
		fmt.Printf("{%d ", n.Val)
		n = n.Next
		if n != nil {
			fmt.Print("*}->")
		} else {
			fmt.Print("/}")
		}
	}
}

// 23h Merge k Sorted Lists
func Test23(t *testing.T) {
	Draw := func(n *ListNode) {
		for n != nil {
			fmt.Printf("{%d ", n.Val)
			n = n.Next
			if n != nil {
				fmt.Print("*}->")
			} else {
				fmt.Print("/}")
			}
		}
	}

	type L = ListNode

	l := []*L{{1, &L{4, &L{Val: 5}}}, {1, &L{3, &L{Val: 4}}}, {2, &L{Val: 6}}}
	for _, e := range l {
		Draw(e)
		log.Print()
	}

	log.Print(" ---| k-Merge |--> ")
	Draw(mergeKLists(l))
	log.Print()
}

// 25h Reverse Nodes in k-Group
func Test25(t *testing.T) {
	type L = ListNode

	l := &L{1, &L{2, &L{3, &L{4, &L{Val: 5}}}}}
	llDraw(l)
	log.Print()
	log.Print(" -> ")

	llDraw(reverseKGroup(l, 2))
	log.Print()
}

// 61m Rotate Right
func Test61(t *testing.T) {
	type L = ListNode

	lists := []*L{
		{1, &L{2, &L{3, &L{4, &L{Val: 5}}}}},
		{0, &L{1, &L{Val: 2}}},
	}
	ks := []int{2, 4}

	for i, k := range ks {
		l := lists[i]
		llDraw(l)
		fmt.Print("   =>   ")
		llDraw(rotateRight(l, k))
		fmt.Print("\n")
	}
}

// 82m Remove Duplicates from Sorted List II
func Test82(t *testing.T) {
	type L = ListNode

	for _, l := range []*L{
		{1, &L{2, &L{3, &L{3, &L{4, &L{4, &L{Val: 5}}}}}}},
		{1, &L{1, &L{1, &L{2, &L{Val: 3}}}}},
	} {
		llDraw(l)
		fmt.Print("   =>   ")
		llDraw(deleteDuplicates(l))
		fmt.Print("\n")
	}
}

// 86m Partition List
func Test86(t *testing.T) {
	type L = ListNode

	lists := []*L{
		{1, &L{4, &L{3, &L{2, &L{5, &L{Val: 2}}}}}},
		{2, &L{Val: 1}},
		{Val: 1},
	}
	for i, x := range []int{3, 2, 2} {
		l := lists[i]
		llDraw(l)
		fmt.Printf("   ==%d=>   ", x)
		llDraw(partition(l, x))
		fmt.Print("\n")
	}
}

// 92m Reverse Linked List II
func Test92(t *testing.T) {
	Reverse := func(head *ListNode) *ListNode {
		var prv *ListNode
		for n := head; n != nil; {
			n.Next, prv, n = prv, n, n.Next
		}
		return prv
	}

	type L = ListNode

	for _, l := range []*L{{1, &L{2, &L{3, &L{4, &L{Val: 5}}}}}, {Val: 1}, {1, &L{Val: 9}}} {
		llDraw(l)
		fmt.Print("  ==R=>  ")
		llDraw(Reverse(l))
		fmt.Print("\n")
	}
	log.Print("---")

	llDraw(reverseBetween(&L{1, &L{2, &L{3, &L{4, &L{Val: 5}}}}}, 2, 4))
	fmt.Print("\n")
	llDraw(reverseBetween(&L{Val: 5}, 1, 1))
	fmt.Print("\n")
}

// 3217m Delete Nodes from Linked List Present in Array
func Test3217(t *testing.T) {
	llDraw := func(n *ListNode) string {
		var s string
		for n != nil {
			if n.Next != nil {
				s += fmt.Sprintf("{%d *}->", n.Val)
			} else {
				s += fmt.Sprintf("{%d /}", n.Val)
			}
			n = n.Next
		}
		return s
	}

	type L = ListNode

	Vals := [][]int{{1, 2, 3}, {1}, {5}}
	for i, l := range []*L{
		{1, &L{2, &L{3, &L{4, &L{Val: 5}}}}},
		{1, &L{2, &L{1, &L{2, &L{1, &L{Val: 2}}}}}},
		{1, &L{2, &L{3, &L{Val: 4}}}},
	} {
		log.Printf("%s   =>   %s", llDraw(l), llDraw(modifiedList(Vals[i], l)))
	}
}
