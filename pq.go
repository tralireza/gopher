package gopher

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
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

// 373m Find K Pairs with Smallest Sums
type E373 struct{ v, x, y int } // Value, Index1, Index2
type PQ373 []E373

func (h PQ373) Len() int           { return len(h) }
func (h PQ373) Less(i, j int) bool { return h[i].v < h[j].v }
func (h PQ373) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ373) Push(x any)        { *h = append(*h, x.(E373)) }
func (h *PQ373) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	type E = E373
	type PQ = PQ373

	R := [][]int{}

	Mem := map[[2]int]bool{}
	h := PQ{}

	heap.Push(&h, E{nums1[0] + nums2[0], 0, 0})
	Mem[[2]int{0, 0}] = true

	for range k {
		e := heap.Pop(&h).(E)
		R = append(R, []int{nums1[e.x], nums2[e.y]})

		// try 2 minimum candidate after (x, y) -> 1: (x+1, y) and 2: (x, y+1)
		// ( if not already in Queue )
		e.x++
		if e.x < len(nums1) && !Mem[[2]int{e.x, e.y}] {
			heap.Push(&h, E{nums1[e.x] + nums2[e.y], e.x, e.y})
			Mem[[2]int{e.x, e.y}] = true
		}

		e.x--
		e.y++
		if e.y < len(nums2) && !Mem[[2]int{e.x, e.y}] {
			heap.Push(&h, E{nums1[e.x] + nums2[e.y], e.x, e.y})
			Mem[[2]int{e.x, e.y}] = true
		}
	}

	return R
}

// 407h Trapping Rain Water II
type Cell407 struct {
	height, r, c int
}
type PQ407 []Cell407

func (h PQ407) Len() int           { return len(h) }
func (h PQ407) Less(i, j int) bool { return h[i].height < h[j].height }
func (h PQ407) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ407) Push(x any)        { *h = append(*h, x.(Cell407)) }
func (h *PQ407) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	Rows, Cols := len(heightMap), len(heightMap[0])

	Vis := make([][]bool, Rows)
	for r := range Vis {
		Vis[r] = make([]bool, Cols)
	}

	pq := &PQ407{}
	for r := range Rows {
		heap.Push(pq, Cell407{heightMap[r][0], r, 0})
		heap.Push(pq, Cell407{heightMap[r][Cols-1], r, Cols - 1})

		Vis[r][0], Vis[r][Cols-1] = true, true
	}
	for c := range Cols {
		heap.Push(pq, Cell407{heightMap[0][c], 0, c})
		heap.Push(pq, Cell407{heightMap[Rows-1][c], Rows - 1, c})

		Vis[0][c], Vis[Rows-1][c] = true, true
	}

	log.Print(" -> ", pq)

	totalWater := 0
	Dir := []int{1, 0, -1, 0, 1}

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell407)

		for i := range 4 {
			r, c := cell.r+Dir[i], cell.c+Dir[i+1]
			if 0 <= r && r < Rows && 0 <= c && c < Cols && !Vis[r][c] {
				Vis[r][c] = true
				h := heightMap[r][c]
				if h < cell.height {
					totalWater += cell.height - h
				}
				heap.Push(pq, Cell407{max(h, cell.height), r, c})
			}
		}
	}

	return totalWater
}

// 632h Smallest Range Covering Elements from K Lists
type Value632 struct{ v, i, ls int } // Value, Index, List#
type PQ632 []Value632

func (h PQ632) Len() int           { return len(h) }
func (h PQ632) Less(i, j int) bool { return h[i].v < h[j].v } // MinHeap
func (h PQ632) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ632) Push(x any)        { *h = append(*h, x.(Value632)) }
func (h *PQ632) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func smallestRange(nums [][]int) []int {
	Q := PQ632{}

	curMax := -100_000
	for i := range nums {
		heap.Push(&Q, Value632{nums[i][0], 0, i})
		curMax = max(nums[i][0], curMax)
	}

	start, end := -100_000, 100_000
	for Q.Len() == len(nums) {
		log.Print(" -> ", Q)
		V := heap.Pop(&Q).(Value632)
		if V.v > curMax {
			curMax = V.v
		}

		if curMax-V.v < end-start {
			start, end = V.v, curMax
		}

		V.i++
		if V.i < len(nums[V.ls]) {
			V.v = nums[V.ls][V.i]
			heap.Push(&Q, V)

			curMax = max(V.v, curMax)
		}
	}

	return []int{start, end}
}

// 1046 Last Stone Weight
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		slices.Sort(stones)

		w, w0 := stones[len(stones)-1], stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if w > w0 {
			stones = append(stones, w-w0)
		}
	}

	return append(stones, 0)[0]
}

// 1405m Longest Happy String
type PQ1405 []Letter1405
type Letter1405 struct {
	chr   byte
	count int
}

func (o Letter1405) String() string { return fmt.Sprintf("{%q %d}", o.chr, o.count) }

func (h PQ1405) Len() int           { return len(h) }
func (h PQ1405) Less(i, j int) bool { return h[i].count > h[j].count }
func (h PQ1405) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ1405) Push(x any)        { *h = append(*h, x.(Letter1405)) }
func (h *PQ1405) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func longestDiverseString(a int, b int, c int) string {
	pq := PQ1405{}

	for i, count := range []int{a, b, c} {
		if count > 0 {
			heap.Push(&pq, Letter1405{'a' + byte(i), count})
		}
	}

	happy := []byte{}
	for pq.Len() > 0 {
		log.Print(" -> ", pq)
		l := heap.Pop(&pq).(Letter1405)

		if len(happy) >= 2 && l.chr == happy[len(happy)-1] && l.chr == happy[len(happy)-2] {
			if pq.Len() == 0 {
				return string(happy)
			}

			n := heap.Pop(&pq).(Letter1405)
			happy = append(happy, n.chr)
			n.count--
			if n.count > 0 {
				heap.Push(&pq, n)
			}

			heap.Push(&pq, l)
		} else {
			happy = append(happy, l.chr)
			l.count--
			if l.count > 0 {
				heap.Push(&pq, l)
			}
		}
	}

	return string(happy)
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

// 1792m Maximum Average Pass Ratio
type PQ1792 []Class1792
type Class1792 struct {
	r, rNext    float64 // Ratio, NextRatio
	pass, total int
}

func (o PQ1792) Len() int           { return len(o) }
func (o PQ1792) Less(i, j int) bool { return o[i].rNext-o[i].r > o[j].rNext-o[j].r }
func (o PQ1792) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

func (o *PQ1792) Push(_ any) {}
func (o PQ1792) Pop() any    { return Class1792{} }

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := PQ1792{}

	for _, v := range classes {
		pq = append(pq, Class1792{
			float64(v[0]) / float64(v[1]),
			float64(v[0]+1) / float64(v[1]+1),
			v[0],
			v[1],
		})
	}
	heap.Init(&pq)

	log.Print(" -> PQ :: ", pq)
	for extraStudents > 0 {
		pq[0].pass++
		pq[0].total++

		pq[0].r = float64(pq[0].pass) / float64(pq[0].total)
		pq[0].rNext = float64(pq[0].pass+1) / float64(pq[0].total+1) // Diff or r & rNext is Grain

		heap.Fix(&pq, 0)

		extraStudents--
		log.Print(" -> PQ :: ", pq)
	}

	mar := float64(0)
	for _, e := range pq {
		mar += float64(e.r)
	}

	return mar / float64(pq.Len())
}

// 1942m The Number of the Smallest Unoccupied Chair
type Chair1942 struct{ n, time int }
type PQ1942 []Chair1942

func (h PQ1942) Len() int { return len(h) }
func (h PQ1942) Less(i, j int) bool {
	if h[i].time == h[j].time {
		return h[i].n < h[j].n
	}
	return h[i].time < h[j].time
}
func (h PQ1942) Swap(i int, j int) { h[i], h[j] = h[j], h[i] }
func (h *PQ1942) Push(x any)       { *h = append(*h, x.(Chair1942)) }
func (h *PQ1942) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func smallestChair(times [][]int, targetFriend int) int {
	D := []int{}
	for f := range times {
		D = append(D, f)
	}

	slices.SortFunc(D, func(a, b int) int { return times[a][0] - times[b][0] })
	log.Print(" -> D :: ", D)

	Q, E := PQ1942{}, PQ1942{} // Occupied, Empty

	for _, f := range D {
		arrive, leave := times[f][0], times[f][1]

		log.Printf(" -> %2d %d %d :: %v", f, arrive, leave, Q)

		for Q.Len() > 0 && Q[0].time <= arrive {
			chair := heap.Pop(&Q).(Chair1942)
			chair.time = 0
			heap.Push(&E, chair)
		}

		if E.Len() == 0 {
			heap.Push(&E, Chair1942{Q.Len(), 0})
		}

		chair := heap.Pop(&E).(Chair1942)
		if f == targetFriend {
			return chair.n
		}

		chair.time = leave
		heap.Push(&Q, chair)
	}

	return -1
}

// 2182m Construct String With Repeat Limit
type PQ2182 []Char2182
type Char2182 struct {
	l     byte
	count int
}

func (h PQ2182) Len() int           { return len(h) }
func (h PQ2182) Less(i, j int) bool { return h[i].l > h[j].l } // MaxHeap
func (h PQ2182) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ2182) Push(x any) { *h = append(*h, x.(Char2182)) }
func (h *PQ2182) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func repeatLimitedString(s string, repeatLimit int) string {
	F := [26]int{}
	for i := 0; i < len(s); i++ {
		F[s[i]-'a']++
	}

	pq := PQ2182{}
	for l, count := range F {
		if count > 0 {
			pq = append(pq, Char2182{byte(l + 'a'), count})
		}
	}
	heap.Init(&pq)

	bfr := strings.Builder{}
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(Char2182)
		for range min(v.count, repeatLimit) {
			bfr.WriteByte(v.l)
		}
		v.count -= repeatLimit

		if v.count > 0 && pq.Len() > 0 {
			u := heap.Pop(&pq).(Char2182)
			bfr.WriteByte(u.l)
			u.count--
			if u.count > 0 {
				heap.Push(&pq, u)
			}

			heap.Push(&pq, v)
		}
	}

	return bfr.String()
}

// 2406m Divide Intervals Into Minimum Number of Groups
type PQ2406 struct{ sort.IntSlice }

func (h *PQ2406) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *PQ2406) Pop() any   { return 0 }

func minGroups(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	log.Print(" -> ", intervals)

	Q := PQ2406{}
	for _, e := range intervals {
		left, right := e[0], e[1]
		if Q.Len() > 0 && Q.IntSlice[0] < left {
			Q.IntSlice[0] = right
			heap.Fix(&Q, 0)
		} else {
			heap.Push(&Q, right)
		}
		log.Printf(" %v -> PQ :: %v", e, Q)
	}
	return Q.Len()
}

// 2530m Maximal Score After Applying K Operations
type PQ2530 struct{ sort.IntSlice }

func (h PQ2530) Less(i, j int) bool { return h.IntSlice[j] < h.IntSlice[i] }
func (h *PQ2530) Push(_ any)        {}
func (h *PQ2530) Pop() any          { return 0 }

func maxKelements(nums []int, k int) int64 {
	h := PQ2530{nums}
	heap.Init(&h)

	score := int64(0)
	for range k {
		log.Print(" -> ", h)

		score += int64(h.IntSlice[0])
		h.IntSlice[0] += 2
		h.IntSlice[0] /= 3
		heap.Fix(&h, 0)
	}
	return score
}

// 2558 Take Gifts From the Richest Pile
type PQ2558 struct{ sort.IntSlice }

func (h PQ2558) Less(i, j int) bool { return h.IntSlice[j] < h.IntSlice[i] }
func (h *PQ2558) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *PQ2558) Pop() any {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func pickGifts(gifts []int, k int) int64 {
	h := PQ2558{IntSlice: gifts}

	for range k {
		v := heap.Pop(&h).(int)
		heap.Push(&h, int(math.Sqrt(float64(v))))
	}

	t := int64(0)
	for _, g := range h.IntSlice {
		t += int64(g)
	}
	return t
}

// 2818h Apply Operations to Maximize Score
type PQ2818 [][2]int

func (h PQ2818) Len() int { return len(h) }
func (h PQ2818) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1] // Smaller Index
	}
	return h[i][0] >= h[j][0] // Larger N
}
func (h PQ2818) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PQ2818) Push(x any)   { *h = append(*h, x.([2]int)) }
func (h *PQ2818) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func maximumScore(nums []int, k int) int {
	xNum := slices.Max(nums)
	Sieve := make([]int, xNum+1)
	for p := 2; p < len(Sieve); p++ {
		if Sieve[p] == 0 {
			for m := p; m < len(Sieve); m += p {
				Sieve[m]++
			}
		}
	}
	log.Print("Prime Factors (distinct): ", Sieve)

	Prev, Next := make([]int, len(nums)), make([]int, len(nums))
	for i := range len(nums) {
		Prev[i], Next[i] = -1, len(nums)
	}

	S := []int{}
	for i, n := range nums {
		for len(S) > 0 && Sieve[nums[S[len(S)-1]]] < Sieve[n] {
			top := S[len(S)-1]
			S = S[:len(S)-1]

			Next[top] = i
		}
		if len(S) > 0 {
			Prev[i] = S[len(S)-1]
		}

		S = append(S, i)
	}

	log.Print("Prev | Next: ", Prev, Next)

	pq := PQ2818{}
	for i, n := range nums {
		heap.Push(&pq, [2]int{n, i})
	}

	log.Print("PQ: ", pq)

	const M = 1e9 + 7
	Power := func(b, e int) int {
		p := 1
		for e > 0 {
			if e&1 == 1 {
				p = (p * b) % M
			}
			b = (b * b) % M
			e >>= 1
		}

		return p
	}

	score := 1
	for k > 0 {
		e := heap.Pop(&pq).([2]int)
		n, i := e[0], e[1]

		log.Print("-> ", e)

		subarrays := (i - Prev[i]) * (Next[i] - i)
		ops := min(k, subarrays)
		k -= ops

		score = (Power(n, ops) * score) % M
	}

	return score
}

// 2940m Find Building Where Alice and Bob Can Meet
type PQ2940 [][2]int

func (h PQ2940) Len() int           { return len(h) }
func (h PQ2940) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h PQ2940) Swap(i int, j int)  { h[i], h[j] = h[j], h[i] }

func (h *PQ2940) Push(x any) { *h = append(*h, x.([2]int)) }
func (h *PQ2940) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	R := make([]int, len(queries))
	for i := range R {
		R[i] = -1
	}

	Q := make([][][2]int, len(heights))
	for i, query := range queries {
		a, b := query[0], query[1] // Alice, Bob
		if a > b {
			a, b = b, a
		}

		if heights[b] > heights[a] || a == b {
			R[i] = b
		} else {
			Q[b] = append(Q[b], [2]int{heights[a], i})
		}
	}

	log.Print(" -> Q :: ", Q)

	pq := PQ2940{}
	for i, h := range heights {
		for pq.Len() > 0 && pq[0][0] < h {
			e := heap.Pop(&pq).([2]int)
			R[e[1]] = i
		}

		for _, e := range Q[i] {
			heap.Push(&pq, e)
		}
	}

	return R
}

// 3066m Minimum Operations to Exceed Threshold Value II
type PQ3066 struct{ sort.IntSlice }

func (h *PQ3066) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *PQ3066) Pop() any {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return x
}

func minOperationsII(nums []int, k int) int {
	pq := PQ3066{nums}
	heap.Init(&pq)

	ops := 0
	for pq.IntSlice[0] < k {
		log.Print("-> ", pq)
		ops++

		x, y := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		heap.Push(&pq, min(x, y)*2+max(x, y))
	}

	return ops
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

// 3257h Maximum Value Sum by Placing Three Rooks II
func maximumValueSum2(board [][]int) int64 {
	Rows, Cols := len(board), len(board[0])

	D := make([][3]int, 0, Rows*Cols)
	for r := range Rows {
		for c := range Cols {
			D = append(D, [3]int{board[r][c], r, c})
		}
	}

	slices.SortFunc(D, func(x, y [3]int) int { return y[0] - x[0] })
	log.Print(D)

	xScore := math.MinInt
	for start := range 5 {
		r, c := D[start][1], D[start][2]

		for x := start + 1; x < 3*max(Rows, Cols); x++ {
			r2, c2 := D[x][1], D[x][2]

			for y := x + 1; y < 3*max(Rows, Cols); y++ {
				r3, c3 := D[y][1], D[y][2]

				if r != r2 && r != r3 && r2 != r3 && c != c2 && c != c3 && c2 != c3 {
					xScore = max(D[start][0]+D[x][0]+D[y][0], xScore)
				}
			}
		}
	}
	return int64(xScore)
}

// 3264 Final Array State After K Multiplication Operations I
type PQ3264 []State3264
type State3264 struct {
	v, i int
}

func (h PQ3264) Len() int { return len(h) }
func (h PQ3264) Less(i, j int) bool {
	if h[i].v == h[j].v {
		return h[i].i < h[j].i
	}
	return h[i].v < h[j].v
}
func (h PQ3264) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PQ3264) Push(_ any)   {}
func (h *PQ3264) Pop() any     { return State3264{} }

func getFinalState(nums []int, k int, multiplier int) []int {
	pq := PQ3264{}
	for i, v := range nums {
		pq = append(pq, State3264{v: v, i: i})
	}
	heap.Init(&pq)

	for k > 0 {
		pq[0].v *= multiplier
		heap.Fix(&pq, 0)

		k--
	}

	for _, s := range pq {
		nums[s.i] = s.v
	}

	return nums
}

// 3362m Zero Array Transformation II
type PQ3362 struct{ sort.IntSlice }

func (o PQ3362) Less(i, j int) bool { return o.IntSlice[i] > o.IntSlice[j] }

func (o *PQ3362) Push(v any) { o.IntSlice = append(o.IntSlice, v.(int)) }
func (o *PQ3362) Pop() any {
	v := o.IntSlice[o.Len()-1]
	o.IntSlice = o.IntSlice[:o.Len()-1]
	return v
}

func maxRemoval(nums []int, queries [][]int) int {
	slices.SortFunc(queries, func(x, y []int) int { return x[0] - y[0] })
	log.Print("-> Q ", queries)

	pq := PQ3362{}

	diff, Diff := 0, make([]int, len(nums)+1)
	q := 0
	for i, n := range nums {
		diff += Diff[i]

		for q < len(queries) && queries[q][0] == i {
			heap.Push(&pq, queries[q][1])
			q++
		}

		for diff < n && pq.Len() > 0 && pq.IntSlice[0] >= i {
			diff++
			Diff[heap.Pop(&pq).(int)+1]--
		}

		if diff < n {
			return -1
		}
	}

	log.Print("-> PQ ", pq)
	return pq.Len()
}
