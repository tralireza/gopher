package gopher

import (
	"container/heap"
)

// 239h Sliding Window Maximum
type E239 struct{ v, i int }
type PQ239 []E239

func (h PQ239) Len() int           { return len(h) }
func (h PQ239) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h PQ239) Less(i, j int) bool { return h[j].v < h[i].v }
func (h *PQ239) Push(x any)        { *h = append(*h, x.(E239)) }
func (h *PQ239) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	type E = E239
	type PQ = PQ239

	h := PQ{}

	R := []int{}
	for i := 0; i < len(nums); i++ {
		heap.Push(&h, E{nums[i], i})
		if i+1 >= k {
			for h[0].i <= i-k {
				heap.Pop(&h)
			}
			R = append(R, h[0].v)
		}
	}
	return R
}

// 1508m Range Sum of Sorted Subarray Sums
type E1508 struct{ n, i int }
type PQ1508 []E1508

func (h PQ1508) Len() int           { return len(h) }
func (h PQ1508) Less(i, j int) bool { return h[i].n < h[j].n }
func (h PQ1508) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ1508) Push(x any)        { *h = append(*h, x.(E1508)) }
func (h *PQ1508) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func rangeSum(nums []int, n int, left int, right int) int {
	type E = E1508
	type PQ = PQ1508

	h := PQ{}
	for i, n := range nums {
		h = append(h, E{n, i})
	}
	heap.Init(&h)

	x := 0
	for i := range right {
		e := heap.Pop(&h).(E)

		if i+1 >= left {
			x += e.n
			x %= 1e9 + 7
		}

		e.i++
		if e.i < len(nums) {
			e.n += nums[e.i]
			heap.Push(&h, e)
		}
	}
	return x
}
