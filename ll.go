package gopher

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
