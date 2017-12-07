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

// 3256h Maximum Value Sum by Placing Three Rooks I
type E3256 struct{ col, score int }
type PQ3256 []E3256

func (h PQ3256) Len() int           { return len(h) }
func (h PQ3256) Less(i, j int) bool { return h[j].score < h[i].score }
func (h PQ3256) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ3256) Push(_ any)        {}
func (h *PQ3256) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func maximumValueSum(board [][]int) int64 {
	type E = E3256
	type PQ = PQ3256

	Rows := len(board)

	D := make([][3]E, Rows)
	for r := range D {
		h := PQ{}
		for c, s := range board[r] {
			h = append(h, E{c, s})
		}
		heap.Init(&h)
		for i := range 3 {
			D[r][i] = heap.Pop(&h).(E)
		}
	}

	xScore := -3000_000_000
	for r := 0; r < Rows; r++ {
		for r2 := r + 1; r2 < Rows; r2++ {
			for r3 := r2 + 1; r3 < Rows; r3++ {

				for _, e := range D[r] {
					for _, e2 := range D[r2] {
						for _, e3 := range D[r3] {
							if e.col != e2.col && e.col != e3.col && e2.col != e3.col { // different columns...
								xScore = max(xScore, e.score+e2.score+e3.score)
							}
						}
					}
				}

			}
		}
	}
	return int64(xScore)
}
