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

// 432h All O'one Data
func Test432(t *testing.T) {
	o := NewAllOne432()

	o.Inc("hello")
	o.Inc("hello")
	log.Print("hello ?= ", o.GetMaxKey())
	log.Print("hello ?= ", o.GetMinKey())
	o.Inc("leet")
	log.Print("hello ?= ", o.GetMaxKey())
	log.Print("leet ?= ", o.GetMinKey())
}

func Test460(t *testing.T) {
	o := NewLFUCache460(2)

	for _, c := range []struct {
		Key, Value int
		rst        any
	}{
		{1, 1, nil},
		{2, 2, nil},
		{1, -1, 1},
		{3, 3, nil},
		{2, -1, -1},
		{3, -1, 3},
		{4, 4, nil},
		{1, -1, -1},
		{3, 1, nil},
		{3, -1, 1},
		{4, -1, 4},
	} {
		switch c.Value {
		case -1:
			if o.Get(c.Key) != c.rst.(int) {
				t.FailNow()
			}
		default:
			o.Put(c.Key, c.Value)
		}
	}
}

// 641m Design Circular Deque
type DLCircularDeque641 struct {
	Head, Tail *LNode641
	Size, Cap  int
}
type LNode641 struct {
	Val        int
	Prev, Next *LNode641
}

func NewDlQ641(cap int) DLCircularDeque641 {
	return DLCircularDeque641{Cap: cap}
}

func (o *DLCircularDeque641) IsEmpty() bool { return o.Size == 0 }
func (o *DLCircularDeque641) IsFull() bool  { return o.Size == o.Cap }
func (o *DLCircularDeque641) InsertFront(v int) bool {
	if o.IsFull() {
		return false
	}

	n := &LNode641{Val: v}
	if o.Head != nil {
		o.Head, n.Next, o.Head.Prev = n, o.Head, n
	} else {
		o.Head, o.Tail = n, n
	}
	o.Size++
	return true
}
func (o *DLCircularDeque641) InsertLast(v int) bool {
	if o.IsFull() {
		return false
	}

	n := &LNode641{Val: v}
	if o.Tail != nil {
		o.Tail, n.Prev, o.Tail.Prev = n, o.Tail, n
	} else {
		o.Tail, o.Head = n, n
	}
	o.Size++
	return true
}
func (o *DLCircularDeque641) DeleteFront() bool {
	if o.IsEmpty() {
		return false
	}

	if o.Size == 1 {
		o.Head, o.Tail = nil, nil
	} else {
		o.Head = o.Head.Next
	}
	o.Size--
	return true
}
func (o *DLCircularDeque641) DeleteLast() bool {
	if o.IsEmpty() {
		return false
	}

	if o.Size == 1 {
		o.Head, o.Tail = nil, nil
	} else {
		o.Tail = o.Tail.Prev
	}
	o.Size--
	return true
}
func (o *DLCircularDeque641) GetFront() int {
	if o.IsEmpty() {
		return -1
	}
	return o.Head.Val
}
func (o *DLCircularDeque641) GetLast() int {
	if o.IsEmpty() {
		return -1
	}
	return o.Tail.Val
}

func Test641(t *testing.T) {
	type Q = DLCircularDeque641

	Draw := func(q *Q) {
		for n := q.Head; n != nil; n = n.Next {
			l, r := '/', '/'
			if n.Next != nil {
				r = '*'
			}
			if n.Prev != nil {
				l = '*'
			}
			fmt.Printf("{%c %d %c}", l, n.Val, r)
		}
		fmt.Print("\n")
	}

	aQ := NewArrayQ641(3) // array as backing store
	dlQ := NewDlQ641(3)   // doubly linked list as backing store

	for _, o := range []CircularDequer641{&aQ, &dlQ} {
		log.Print("true ?= ", o.InsertLast(1))
		log.Print("1 ?= ", o.GetFront())
		log.Print("1 ?= ", o.GetLast())
		log.Print("true ?= ", o.InsertLast(2))
		log.Print("true ?= ", o.InsertFront(3))
		log.Print("false ?= ", o.InsertFront(4))
		log.Print("2 ?= ", o.GetLast())
		log.Print("true ?= ", o.IsFull())
		log.Print("true ?= ", o.DeleteLast())
		log.Print("true ?= ", o.InsertFront(4))
		log.Print("4 ?= ", o.GetFront())
		log.Print("--")
	}

	Draw(&dlQ)
}

func Test706(t *testing.T) {
	// 0 <= Key,Value <= 10^6
	o := NewMyHashMap706()
	o.Put(1, 1)
	o.Put(2, 2)
	if o.Get(1) != 1 {
		t.FailNow()
	}
	if o.Get(3) != -1 {
		t.FailNow()
	}

	o.Put(5, 5)
	o.Put(15, 15)
	o.Remove(5)
	if o.Get(15) != 15 {
		t.FailNow()
	}

	o.Put(15, 5)
	if o.Get(15) != 5 {
		t.FailNow()
	}

	log.Print(o)
}

// 725m Split Linked List in Parts
func Test725(t *testing.T) {
	type L = ListNode

	Draw := func(Seg []*ListNode) {
		fmt.Print("[")
		for _, l := range Seg {
			if l != nil {
				fmt.Print(" [")
				for l != nil {
					if l.Next != nil {
						fmt.Printf("{%d *}->", l.Val)
					} else {
						fmt.Printf("{%d /}", l.Val)
					}
					l = l.Next
				}
				fmt.Print("]")
			} else {
				fmt.Print(" []")
			}
		}
		fmt.Print(" ]\n")
	}

	Draw(splitListToParts(&L{1, &L{2, &L{Val: 3}}}, 5))
	Draw(splitListToParts(&L{1, &L{2, &L{3, &L{4, &L{5, &L{6, &L{7, &L{8, &L{9, &L{Val: 10}}}}}}}}}}, 3))
}

// 2326m Spiral Matrix IV
func Test2326(t *testing.T) {
	// 0 <= Node.Val <= 1000

	Trick := func(m, n int, head *ListNode) [][]int {
		M := make([][]int, m)
		for r := range M {
			M[r] = make([]int, n)
			for c := range M[r] {
				M[r][c] = -1
			}
		}

		d, Dir := 0, []int{0, 1, 0, -1, 0}

		r, c := 0, 0
		for head != nil {
			M[r][c] = head.Val
			head = head.Next

			rNxt, cNxt := r+Dir[d], c+Dir[d+1]
			if cNxt == n || rNxt == m || cNxt < 0 || M[rNxt][cNxt] != -1 {
				d = (d + 1) % 4
				r, c = r+Dir[d], c+Dir[d+1]
			} else {
				r, c = rNxt, cNxt
			}
		}

		return M
	}

	type L = ListNode

	for _, fn := range []func(int, int, *ListNode) [][]int{spiralMatrix, Trick} {
		log.Print("[[3 0 2 6 8] [5 0 -1 -1 1] [5 2 4 9 7]] ?= ", fn(3, 5, &L{3, &L{0, &L{2, &L{6, &L{8, &L{1, &L{7, &L{9, &L{4, &L{2, &L{5, &L{5, &L{Val: 0}}}}}}}}}}}}}))
		log.Print("[[0 1 2 -1]] ?= ", fn(1, 4, &L{0, &L{1, &L{Val: 2}}}))
		log.Print("--")
	}
}

// 2807m Insert Greatest Common Divisors in Linked List
func Test2807(t *testing.T) {
	llDraw := func(l *ListNode) string {
		var s string
		for l != nil {
			if l.Next != nil {
				s += fmt.Sprintf("{%d *}->", l.Val)
			} else {
				s += fmt.Sprintf("{%d /}", l.Val)
			}
			l = l.Next
		}
		return s
	}

	type L = ListNode

	for _, l := range []*L{
		{18, &L{6, &L{10, &L{Val: 3}}}},
		{Val: 7},
	} {
		log.Print(llDraw(l), "  =>  ", llDraw(insertGreatestCommonDivisors(l)))
	}
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
