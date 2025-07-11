package gopher

import (
	"container/list"
	"log"
	"math"
	"slices"
	"strconv"
)

// 153m Find Minimum in Rotated Sorted Array
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		log.Print(l, m, r, nums)

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

// 154h Find Minimum in Rotated Sorted Array II
func findMinII(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		log.Print(l, m, r, nums)

		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r++
		}

	}

	return nums[l]
}

// 274m H-Index
func hIndex(citations []int) int {
	slices.Sort(citations)

	Check := func(m int) int {
		x := 0
		for i := 0; i < len(citations); i++ {
			if citations[i] >= m {
				x++
			}
		}
		return x
	}

	l, r := 0, len(citations)
	var h int
	for l <= r {
		m := l + (r-l)>>1

		v := Check(m)

		log.Print(l, m, r, " :: ", v)

		if v >= m {
			l = m + 1
			h = m
		} else {
			r = m - 1
		}
	}
	return h
}

// 315h Count of Smaller Numbers After Self
type BIT315 []int // "BIT: Binary Index Tree" storage

func (t BIT315) Update(i, Val int) {
	for i < len(t) {
		t[i] += Val
		i += i & -i
	}
}
func (t BIT315) Query(i int) int {
	v := 0
	for i > 0 {
		v += t[i]
		i -= i & -i
	}
	return v
}

type SegmentTree315 struct {
	tree []int
}

func NewSegmentTree315(size int) SegmentTree315 {
	return SegmentTree315{
		tree: make([]int, 4*size),
	}
}
func (t *SegmentTree315) Update(v, p, Left, Right int) {
	if Left == Right {
		t.tree[v]++
		return
	}

	Mid := Left + (Right-Left)>>1
	if p <= Mid {
		t.Update(2*v, p, Left, Mid) // left child "Seg"
	} else {
		t.Update(2*v+1, p, Mid+1, Right) // right child "Seg"
	}
	t.tree[v] = t.tree[2*v] + t.tree[2*v+1]
}
func (t *SegmentTree315) Query(v, qryLeft, qryRight, Left, Right int) int {
	if qryLeft > qryRight {
		return 0
	}
	if qryLeft == Left && qryRight == Right {
		return t.tree[v]
	}

	Mid := Left + (Right-Left)>>1
	lVal := t.Query(2*v, qryLeft, min(Mid, qryRight), Left, Mid)
	rVal := t.Query(2*v+1, max(Mid+1, qryLeft), qryRight, Mid+1, Right)
	return lVal + rVal
}

func countSmaller(nums []int) []int {
	D := make([][3]int, len(nums))
	for i, n := range nums {
		D[i] = [3]int{n, i, 0}
	}

	B := make([][3]int, len(nums))
	copy(B, D)

	var MergeSort func(data, bfr [][3]int, l, r int)
	MergeSort = func(data, bfr [][3]int, l, r int) {
		if l >= r {
			return
		}

		m := l + (r-l)>>1
		MergeSort(bfr, data, l, m)
		MergeSort(bfr, data, m+1, r)

		smaller := 0
		p, q, x := l, m+1, l
		for ; p <= m && q <= r; x++ {
			if bfr[p][0] <= bfr[q][0] {
				data[x] = bfr[p]
				data[x][2] += smaller
				p++
			} else {
				data[x] = bfr[q]
				smaller++
				q++
			}
		}
		for ; p <= m; x++ {
			data[x] = bfr[p]
			data[x][2] += smaller
			p++
		}
		for ; q <= r; x++ {
			data[x] = bfr[q]
			q++
		}
	}

	MergeSort(D, B, 0, len(nums)-1)

	R := make([]int, len(D))
	for _, v := range D {
		R[v[1]] = v[2]
	}
	return R
}

// 492 Construct the Rectangle
func constructRectangle(area int) []int {
	x, w := 1, 1
	for (x+1)*(x+1) <= area {
		x++
		if area%x == 0 {
			w = x
		}
	}

	return []int{area / w, w}
}

// 564h Find the Closest Palindrome
func nearestPalindromic(n string) string {
	Value := func(s string) int {
		v := 0
		for i := 0; i < len(s); i++ {
			v = v*10 + int(s[i]-'0')
		}
		return v
	}

	Palin := func(v int) int {
		s := strconv.Itoa(v)
		l, r := (len(s)-1)/2, len(s)/2
		bs := []byte(s)
		for l >= 0 {
			bs[r] = bs[l]
			l--
			r++
		}
		return Value(string(bs))
	}

	N := Value(n)

	Next := func() int {
		var v int
		l, r := N, math.MaxInt
		for l <= r {
			m := l + (r-l)>>1
			p := Palin(m)
			if p > N {
				v = p
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return v
	}

	Prev := func() int {
		var v int
		l, r := 0, N
		for l <= r {
			m := l + (r-l)>>1
			p := Palin(m)
			if p < N {
				v = p
				l = m + 1
			} else {
				r = m - 1
			}
		}
		return v
	}

	prev, next := Prev(), Next()
	log.Print(prev, " <  N: ", N, "  < ", next)

	if N-prev <= next-N {
		return strconv.Itoa(prev)
	}
	return strconv.Itoa(next)
}

// 704 Binary Search
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else {
			if nums[m] == target {
				return m
			}
			r = m - 1
		}
	}
	return -1
}

// 793h Preimage Size of Factorial Zeroes Function
// 0 <= k <= 10^9
func preimageSizeFZF(k int) int {
	fzf := func(x int) int {
		fzCount := 0
		for x/5 > 0 {
			fzCount += x / 5
			x /= 5
		}
		return fzCount
	}

	f, x := 1, 0
	for x <= 15 {
		log.Printf("-> %d! = %d  fzf: %d", x, f, fzf(x))
		x++
		f *= x
	}
	log.Printf("-> fzf(%d) = %d", 52, fzf(52))
	log.Printf("-> fzf(1e9) = %d", fzf(int(1e9)))

	lSearch := func(target int) int { // Left-Most BinSearch
		l, r := 0, math.MaxInt
		for l < r {
			m := l + (r-l)>>1
			if fzf(m) < target {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}
	rSearch := func(target int) int { // Right-Most BinSearch
		l, r := 0, math.MaxInt
		for l < r {
			m := l + (r-l)>>1
			if fzf(m) > target {
				r = m
			} else {
				l = m + 1
			}
		}
		return r - 1
	}

	l, r := lSearch(k), rSearch(k)
	log.Printf("-> %d %d   lSearch(+1): %d", l, r, lSearch(k+1))

	return r - l + 1
}

// 1351 Count Negative Numbers in a Sorted Matrix
func countNegatives(grid [][]int) int {
	BSRight := func(nums []int, t int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if -nums[m] > t {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	negs := 0
	for _, row := range grid {
		negs += len(row) - BSRight(row, 0)
	}

	return negs
}

// 1498m Number of Subsequences That Satisfy the Given Sum Condition
func numSubseq(nums []int, target int) int {
	const M = int(1e9) + 7

	mPower := func(base, exponent int) int {
		power := 1
		for exponent > 0 {
			if exponent&1 == 1 {
				power = (base * power) % M
			}
			base = (base * base) % M
			exponent >>= 1
		}
		return power
	}

	TwoPointers := func(nums []int, target int) int {
		Powers := make([]int, len(nums))
		Powers[0] = 1
		for i := range Powers[:len(Powers)-1] {
			Powers[i+1] = (2 * Powers[i]) % M
		}
		log.Print("-> Mod Powers: ", Powers)

		slices.Sort(nums)

		count := 0
		l, r := 0, len(nums)-1
		for l <= r {
			if nums[l]+nums[r] <= target {
				count = (count + Powers[r-l]) % M
				l++
			} else {
				r--
			}
		}

		return count
	}
	log.Print(":? Two Pointers -> ", TwoPointers(nums, target))

	count := 0
	slices.Sort(nums)
	for p := 0; p < len(nums); p++ {
		l, r := p, len(nums)-1
		for l <= r {
			m := l + (r-l)>>1
			if nums[m] <= target-nums[p] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

		if r >= p {
			count += mPower(2, r-p)
			count %= M
		}
	}

	log.Print(":? ", count)
	return count
}

// 1760m Minimum Limit of Balls in a Bag
func minimumSize(nums []int, maxOperations int) int {
	Possible := func(m int) bool {
		ops := 0
		for _, n := range nums {
			if n > m {
				ops += (n - 1) / m
			}
			if ops > maxOperations {
				return false
			}
		}
		return true
	}

	l, r := 1, slices.Max(nums)
	for l < r {
		m := l + (r-l)>>1
		log.Print(" -> ", l, m, r, Possible(m))

		if Possible(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// 1894m Find the Student that Will Replace the Chalk
func chalkReplacer(chalk []int, k int) int {
	pSum := make([]int, len(chalk))
	pSum[0] = chalk[0]
	for i, n := range chalk[1:] {
		pSum[i+1] = n + pSum[i]
	}
	k %= pSum[len(chalk)-1]

	l, r := 0, len(chalk)-1
	for l < r {
		m := l + (r-l)>>1
		if pSum[m] <= k {
			l = m + 1
		} else {
			r = m
		}
	}
	return r
}

// 2040h Kth Smallest Product of Two Sorted Array
func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	countSmaller := func(v int64) int64 {
		count := int64(0)
		for _, n := range nums1 {
			l, r := 0, len(nums2)-1
			for l <= r {
				m := l + (r-l)>>1
				p := int64(n) * int64(nums2[m])
				if n >= 0 && p <= v || n < 0 && p > v {
					l = m + 1
				} else {
					r = m - 1
				}
			}

			if n >= 0 {
				count += int64(l)
			} else {
				count += int64(len(nums2) - l)
			}
		}

		return count
	}

	l, r := -int64(1e10), int64(1e10)
	for l <= r {
		m := l + (r-l)>>1

		log.Print("-> ", l, m, r)

		if countSmaller(m) < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	log.Print(":: ", l)
	return l
}

// 2071h Maximum Number of Tasks You Can Assign
func maxTaskAssign(tasks, workers []int, pills, strength int) int {
	slices.Sort(tasks)
	slices.Sort(workers)

	Check := func(tasks, workers []int, pills int) bool {
		Q := list.New()

		t := 0
		for _, wkr := range workers {
			for t < len(tasks) && wkr+strength >= tasks[t] {
				Q.PushBack(tasks[t])
				t++
			}

			if Q.Len() == 0 {
				return false
			}

			if Q.Front().Value.(int) <= wkr {
				Q.Remove(Q.Front())
			} else {
				if pills == 0 {
					return false
				}
				pills--
				Q.Remove(Q.Back())
			}
		}

		return true
	}

	l, r := 0, min(len(tasks), len(workers))
	for l < r {
		m := l + (r-l+1)>>1
		log.Print("-> ", l, m, r)

		if Check(tasks[:m], workers[len(workers)-m:], pills) {
			l = m
		} else {
			r = m - 1
		}
	}

	log.Print("<- ", l, r)
	return l
}

// 2226m Maximum Candies Allocated to K Children
func maximumCandies(candies []int, k int64) int {
	Possible := func(m int) bool {
		if m == 0 {
			return true
		}

		t := int64(0)
		for _, c := range candies {
			t += int64(c / m)
		}
		return t >= k
	}

	l, r := 0, slices.Max(candies)
	for l <= r {
		m := l + (r-l)>>1
		log.Print(l, m, r)
		if Possible(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

// 2529 Maximum Count of Positive Integer and Negative Integer
func maximumCount(nums []int) int {
	BSLeft := func(t int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] < t {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}

	BSRight := func(t int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] > t {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	return max(len(nums)-BSRight(0), BSLeft(0))
}

// 2560m House Robber IV
func minCapability(nums []int, k int) int {
	l, r := slices.Min(nums), slices.Max(nums)
	for l < r {
		m := l + (r-l)>>1

		steals := 0
		for p := 0; p < len(nums); p++ {
			if nums[p] <= m {
				steals++
				p++
			}
		}

		if steals >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// 2563m Count the Number of Fair Pairs
func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)

	BSearch := func(l, r int, target int) int {
		for l < r {
			m := l + (r-l)>>1
			if nums[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}

	count := 0
	for i, n := range nums {
		count += BSearch(i+1, len(nums), upper-n+1) - BSearch(i+1, len(nums), lower-n)
	}
	log.Print(":: BSearch: ", count)

	Less := func(target int) int64 {
		pairs := int64(0)

		l, r := 0, len(nums)-1
		for l < r {
			log.Print("-> ", l, r)

			if nums[l]+nums[r] < target {
				pairs += int64(r - l)
				l++
			} else {
				r--
			}
		}

		return pairs
	}

	return Less(upper+1) - Less(lower)
}

// 2594m Minimum Time to Repair Cars
func repairCars(ranks []int, cars int) int64 {
	l, r := int64(1), int64(slices.Min(ranks))*int64(cars)*int64(cars)

	for l < r {
		m := l + (r-l)>>1

		repairs := 0
		for _, r := range ranks {
			repairs += int(math.Sqrt(float64(m / int64(r))))
		}

		if repairs < cars {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

// 2616m Minimize the Maximum Difference of Pairs
func minimizeMax(nums []int, p int) int {
	slices.Sort(nums)

	CountPairs := func(m int) int {
		count := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= m {
				count++
				i++
			}
		}

		return count
	}

	l, r := 0, nums[len(nums)-1]-nums[0]
	for l < r {
		m := l + (r-l)>>1
		if CountPairs(m) >= p {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

// 3224m Minimum Array Changes to Make Difference Equal
func minChanges(nums []int, k int) int {
	M := map[int]int{}
	Diffs := make([]int, 0, len(nums)/2)

	l, r := 0, len(nums)-1
	for l < r {
		A, a := nums[l], nums[r]
		if a > A {
			A, a = a, A
		}

		M[A-a]++

		// maximum difference of "pair" elements that can be fixed by one operation
		// ... by setting either: a to 0 or A to k
		Diffs = append(Diffs, max(A, k-a))

		l++
		r--
	}

	log.Print("Difference Frequency -> ", M)

	slices.Sort(Diffs)
	log.Print("(One Operation) Maximum Difference -> ", Diffs)

	minOps := math.MaxInt
	for x, f := range M {
		l, r := 0, len(Diffs)-1
		for l < r {
			m := l + (r-l)>>1
			if Diffs[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		minOps = min(minOps, len(nums)/2-f+l)
	}
	return minOps
}

// 3296m Minimum Number of Seconds to Make Mountain Height Zero
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	Check := func(m int) bool {
		hCur := mountainHeight

		for i, t := range workerTimes {
			x := 1
			for t <= m {
				hCur--
				if hCur == 0 {
					return true
				}
				x++
				t += x * workerTimes[i]
			}
		}

		return false
	}

	l, r := 0, math.MaxInt
	for l < r {
		m := l + (r-l)>>1
		if Check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l)
}

// 3356m Zero Array Transformation II
func minZeroArray(nums []int, queries [][]int) int {
	Possible := func(m int) bool {
		M := make([]int, len(nums)+1)
		for k := 0; k < m; k++ {
			qry := queries[k]
			M[qry[0]] += qry[2]
			M[qry[1]+1] -= qry[2]
		}

		log.Print("-> ", m, M)

		tSum := 0
		for x := range nums {
			tSum += M[x]
			if nums[x] > tSum {
				return false
			}
		}
		return true
	}

	BSearch := func() int {
		l, r := 0, len(queries)
		for l <= r {
			m := l + (r-l)>>1
			if Possible(m) {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return l
	}

	if !Possible(len(queries)) {
		return -1
	}
	return BSearch()
}
