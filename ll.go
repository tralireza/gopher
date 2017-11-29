package gopher

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 23h Merge k Sorted Lists
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	m := len(lists) / 2
	l, r := mergeKLists(lists[:m]), mergeKLists(lists[m:])

	h := &ListNode{}

	p := h
	for l != nil || r != nil {
		if l != nil && (r == nil || l.Val <= r.Val) {
			p.Next, l = l, l.Next
		} else {
			p.Next, r = r, r.Next
		}
		p = p.Next
	}

	return h.Next
}

// 25h Reverse Nodes in k-Group
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n, l := head, 0
	for n != nil && l < k {
		l++
		n = n.Next
	}
	if l < k {
		return head
	}

	r := reverseKGroup(n, k)

	n, Q := head, []*ListNode{}
	for range k {
		Q = append(Q, n)
		n = n.Next
	}

	h := &ListNode{}
	p := h
	for len(Q) > 0 {
		p.Next = Q[len(Q)-1]
		Q = Q[:len(Q)-1]
		p = p.Next
	}
	p.Next = r
	return h.Next
}

// 61m Rotate Right
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l := 0
	for n := head; n != nil; n = n.Next {
		l++
	}

	// edge cases ...
	if l < 2 {
		return head
	}
	k %= l
	if k == 0 {
		return head
	}

	var p *ListNode
	n := head
	for ; l-k > 0; n = n.Next {
		p = n
		l--
	}
	ph := head
	head, p.Next = n, nil
	for ; n != nil; n = n.Next {
		p = n
	}
	p.Next = ph
	return head
}

// 82m Remove Duplicates from Sorted List II
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	v := head.Val
	if v == head.Next.Val { // head has a duplicate at .Next ...
		n := head
		for ; n != nil && n.Val == v; n = n.Next {
		}
		return deleteDuplicates(n)
	}

	// head didn't have a duplicate ...
	head.Next = deleteDuplicates(head.Next)
	return head
}

// 92m Reverse Linked List II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	h := &ListNode{Next: head}

	n := h
	for range left - 1 {
		n = n.Next
	}
	l := n
	n = n.Next

	log.Print(" -- left (prv): ", l)

	var prv *ListNode
	for range right - left + 1 { // reverse #N nodes
		n.Next, prv, n = prv, n, n.Next
	}

	log.Print(" -- right (next): ", n)

	l.Next, l.Next.Next = prv, n
	return h.Next
}
