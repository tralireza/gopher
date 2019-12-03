package gopher

import (
	"container/list"
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

// 432h All O'one Data Structure
type AllOne432 struct {
	Nodes      map[string]*LNode432
	Head, Tail *LNode432
}

func NewAllOne432() AllOne432 {
	o := AllOne432{
		Nodes: map[string]*LNode432{},
		Head:  &LNode432{}, Tail: &LNode432{},
	}
	o.Head.Next, o.Tail.Prev = o.Tail, o.Head
	return o
}

type LNode432 struct {
	Count      int
	Keys       map[string]struct{}
	Prev, Next *LNode432
}

func (o *AllOne432) Inc(key string) {
	if _, ok := o.Nodes[key]; !ok {
		c := o.Head.Next
		if c == o.Tail || c.Count != 1 {
			n := &LNode432{Count: 1, Keys: map[string]struct{}{}}
			n.Keys[key] = struct{}{}
			n.Next, n.Prev = c, o.Head
			o.Head.Next = n
			c.Prev = n
			o.Nodes[key] = n
		} else {
			c.Keys[key] = struct{}{}
			o.Nodes[key] = c
		}

		return
	}

	c := o.Nodes[key]
	delete(c.Keys, key)

	if c.Next == o.Tail || c.Next.Count != c.Count+1 {
		n := &LNode432{Count: c.Count + 1, Keys: map[string]struct{}{}}
		n.Keys[key] = struct{}{}
		n.Next, n.Prev = c.Next, c
		c.Next, c.Next.Prev = n, n
		o.Nodes[key] = n
	} else {
		c.Next.Keys[key] = struct{}{}
		o.Nodes[key] = c.Next
	}

	if len(c.Keys) == 0 {
		c.Prev.Next, c.Next.Prev = c.Next, c.Prev
	}
}

func (o *AllOne432) Dec(key string) {
	if _, ok := o.Nodes[key]; !ok {
		return
	}

	n := o.Nodes[key]
	delete(n.Keys, key)

	if n.Count == 1 {
		delete(o.Nodes, key)
	} else {
		if n.Prev == o.Head || n.Prev.Count != n.Count-1 {
			p := &LNode432{Count: n.Count - 1, Keys: map[string]struct{}{}}
			p.Keys[key] = struct{}{}
			p.Next, p.Prev = n, n.Prev
			n.Prev.Next, n.Prev = p, p
			o.Nodes[key] = p
		} else {
			n.Prev.Keys[key] = struct{}{}
			o.Nodes[key] = n.Prev
		}
	}

	if len(n.Keys) == 0 {
		n.Prev.Next, n.Next.Prev = n.Next, n.Prev
	}
}

func (o *AllOne432) GetMaxKey() string {
	if o.Tail.Prev == o.Head {
		return ""
	}

	for k := range o.Tail.Prev.Keys {
		return k
	}
	return ""
}

func (o *AllOne432) GetMinKey() string {
	if o.Head.Next == o.Tail {
		return ""
	}

	for k := range o.Head.Next.Keys {
		return k
	}
	return ""
}

// 460h LFU Cache
type LFUCache460 struct {
	Cap  int
	fMin int

	Nodes map[int]*list.Element
	LFU   map[int]*list.List
}

type Node460 struct {
	Key, Val int
	frq      int
}

func NewLFUCache460(capacity int) LFUCache460 {
	return LFUCache460{
		Cap:   capacity,
		Nodes: map[int]*list.Element{},
		LFU:   map[int]*list.List{},
	}
}

func (o *LFUCache460) Get(key int) int {
	if _, ok := o.Nodes[key]; !ok {
		return -1
	}

	lNode := o.Nodes[key]
	n := lNode.Value.(*Node460)

	ls := o.LFU[n.frq]
	ls.Remove(lNode)

	if ls.Len() == 0 {
		delete(o.LFU, n.frq)

		if o.fMin == n.frq {
			o.fMin++
		}
	}

	n.frq++
	if _, ok := o.LFU[n.frq]; !ok {
		o.LFU[n.frq] = list.New()
	}

	{
		ls := o.LFU[n.frq]
		lNode = ls.PushFront(&Node460{
			Key: n.Key,
			Val: n.Val,
			frq: n.frq,
		})
		o.Nodes[key] = lNode
	}

	return n.Val
}

func (o *LFUCache460) Put(key, value int) {
	if _, ok := o.Nodes[key]; ok {
		o.Get(key)

		lNode := o.Nodes[key]
		n := lNode.Value.(*Node460)
		n.Val = value
		return
	}

	if len(o.Nodes) == o.Cap {
		ls := o.LFU[o.fMin]

		lNode := ls.Back()
		ls.Remove(lNode)

		if ls.Len() == 0 {
			delete(o.LFU, o.fMin)
		}

		n := lNode.Value.(*Node460)
		delete(o.Nodes, n.Key)
	}

	o.fMin = 1
	if _, ok := o.LFU[1]; !ok {
		o.LFU[1] = list.New()
	}

	o.Nodes[key] = o.LFU[1].PushFront(&Node460{
		Key: key,
		Val: value,
		frq: 1,
	})

	log.Print("-> ", o)
}

// 641m Design Circular Deque
type CircularDequer641 interface {
	IsEmpty() bool
	IsFull() bool
	InsertFront(v int) bool
	InsertLast(v int) bool
	DeleteFront() bool
	DeleteLast() bool
	GetFront() int
	GetLast() int
}

type CircularDeque641 struct {
	vals        []int
	front, last int
	size, cap   int
}

func NewArrayQ641(cap int) CircularDeque641 {
	return CircularDeque641{
		vals: make([]int, cap),
		cap:  cap,
		last: cap - 1,
	}
}
func (o *CircularDeque641) IsEmpty() bool { return o.size == 0 }
func (o *CircularDeque641) IsFull() bool  { return o.size == o.cap }
func (o *CircularDeque641) InsertFront(v int) bool {
	if o.IsFull() {
		return false
	}

	o.front = (o.front - 1 + o.cap) % o.cap
	o.vals[o.front] = v

	o.size++
	return true
}
func (o *CircularDeque641) InsertLast(v int) bool {
	if o.IsFull() {
		return false
	}

	o.last = (o.last + 1) % o.cap
	o.vals[o.last] = v

	o.size++
	return true
}
func (o *CircularDeque641) DeleteFront() bool {
	if o.IsEmpty() {
		return false
	}

	o.front = (o.front + 1) % o.cap

	o.size--
	return true
}
func (o *CircularDeque641) DeleteLast() bool {
	if o.IsEmpty() {
		return false
	}

	o.last = (o.last - 1 + o.cap) % o.cap

	o.size--
	return true
}
func (o *CircularDeque641) GetFront() int {
	if o.IsEmpty() {
		return -1
	}
	return o.vals[o.front]
}
func (o *CircularDeque641) GetLast() int {
	if o.IsEmpty() {
		return -1
	}
	return o.vals[o.last]
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
func spiralMatrix(m, n int, head *ListNode) [][]int {
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

// 2807m Insert Greatest Common Divisors in Linked List
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	a, b := head.Val, head.Next.Val
	if b > a {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}

	head.Next = &ListNode{a, insertGreatestCommonDivisors(head.Next)}
	return head
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
