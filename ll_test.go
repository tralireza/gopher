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
