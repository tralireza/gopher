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

// 86m Partition List
func partition(head *ListNode, x int) *ListNode {
	lh, gh := &ListNode{}, &ListNode{}

	l, g := lh, gh
	for n := head; n != nil; n = n.Next {
		if n.Val < x {
			l.Next = n
			l = n
		} else {
			g.Next = n
			g = n
		}
	}
	l.Next, g.Next = gh.Next, nil

	return lh.Next
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

// 725m Split Linked List in Parts
func splitListToParts(head *ListNode, k int) []*ListNode {
	lZ := 0
	for n := head; n != nil; n = n.Next {
		lZ++
	}

	Seg := []*ListNode{}

	n := head
	for i := range k {
		extra := 1
		if i >= lZ%k {
			extra = 0
		}

		h := &ListNode{Next: n}
		p := h
		for range lZ/k + extra {
			if n == nil {
				break
			}
			p.Next, p = n, n
			n = n.Next
		}
		p.Next = nil

		Seg = append(Seg, h.Next)
	}

	return Seg
}

// 2326m Spiral Matrix IV
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	M := make([][]int, m)
	for r := range M {
		M[r] = make([]int, n)
		for c := range M[r] {
			M[r][c] = -1
		}
	}

	const (
		RIGHT = iota
		DOWN
		LEFT
		UP
	)

	Dir := []int{0, 1, 0, -1, 0}

	rX, cX, rY, cY := 0, 0, m-1, n-1 // X: top-left, Y: bottom-right

	r, c := 0, 0
	d := RIGHT
	rX++

	for head != nil {
		M[r][c] = head.Val

		r, c = r+Dir[d], c+Dir[d+1]

		if d == RIGHT && c > cY {
			d = DOWN
			r, c = r+1, cY
			cY--
		} else if d == DOWN && r > rY {
			d = LEFT
			r, c = rY, c-1
			rY--
		} else if d == LEFT && c < cX {
			d = UP
			r, c = r-1, cX
			cX++
		} else if d == UP && r < rX {
			d = RIGHT
			r, c = rX, c+1
			rX++
		}

		head = head.Next
	}

	return M
}

// 3217m Delete Nodes from Linked List Present in Array
func modifiedList(nums []int, head *ListNode) *ListNode {
	Vals := map[int]struct{}{}
	for _, n := range nums {
		Vals[n] = struct{}{}
	}

	h := &ListNode{Next: head}
	prv := h
	for n := head; n != nil; n = n.Next {
		if _, ok := Vals[n.Val]; ok {
			prv.Next = n.Next
		} else {
			prv = n
		}
	}
	return h.Next
}
