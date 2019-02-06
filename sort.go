package gopher

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

// 56m Merge Intervals
func merge(intervals [][]int) [][]int {
	I := [][]int{}

	slices.SortFunc(intervals, func(x, y []int) int { return x[0] - y[0] })

	prv := intervals[0]
	for _, v := range intervals[1:] {
		if prv[1] >= v[0] { // merge
			prv[1] = max(prv[1], v[1])
		} else {
			I = append(I, prv)
			prv = v
		}
	}
	I = append(I, prv)

	return I
}

// 179m Largest Number
func largestNumber(nums []int) string {
	x := 0
	for _, n := range nums {
		l := 0
		for n > 0 {
			l++
			n /= 10
		}
		x = max(l, x)
	}

	S := []string{}
	for _, n := range nums {
		S = append(S, fmt.Sprintf("%-[1]*d", x, n))
	}

	slices.SortFunc(S, func(a, b string) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] == ' ' && b[i] == ' ' {
				return int(a[i-1]) - int(b[i-1])
			}
			if a[i] == b[i] {
				continue
			}

			if a[i] == ' ' {
				var x, y int
				fmt.Sscanf(strings.ReplaceAll(a[:i]+b, " ", ""), "%d", &x)
				fmt.Sscanf(strings.ReplaceAll(b+a[:i], " ", ""), "%d", &y)

				return x - y

			} else if b[i] == ' ' {
				var x, y int
				fmt.Sscanf(strings.ReplaceAll(a+b[:i], " ", ""), "%d", &x)
				fmt.Sscanf(strings.ReplaceAll(b[:i]+a, " ", ""), "%d", &y)

				return x - y
			}

			return int(a[i]) - int(b[i])
		}
		return 0
	})

	log.Printf(" -> %q", S)

	if S[len(S)-1] == "0" {
		return "0"
	}

	var v string
	for i := range S {
		v = strings.Replace(S[i], " ", "", -1) + v
	}
	return v
}

// 414 Third Maximum Number
func thirdMax(nums []int) int {
	f, s, t := 0, -1, 1 // First, Second & Third Max

	for i, n := range nums {
		switch {
		case n > nums[f]:
			f, s, t = i, f, s
		case n < nums[f] && (s == -1 || n > nums[s]):
			s, t = i, s
		case s > -1 && n < nums[s] && (t == -1 || n > nums[t]):
			t = i
		}
	}

	if t > -1 {
		return nums[t]
	}
	return nums[f]
}

// 539m Minimum Time Difference
func findMinDifference(timePoints []string) int {
	Ms := []int{}

	for _, t := range timePoints {
		hr, _ := strconv.Atoi(t[:2])
		mn, _ := strconv.Atoi(t[3:])
		Ms = append(Ms, 60*hr+mn)
	}

	slices.Sort(Ms)
	log.Print(Ms)

	mVal := 720
	for i := range len(Ms) - 1 {
		mVal = min(Ms[i+1]-Ms[i], mVal)
	}
	return min(1440-(Ms[len(Ms)-1]-Ms[0]), mVal)
}

// 905 Sort Array By Parity
func sortArrayByParity(nums []int) []int {
	R := make([]int, 0, len(nums))

	for _, n := range nums {
		if n&1 == 0 {
			R = append(R, n)
		}
	}
	for _, n := range nums {
		if n&1 == 1 {
			R = append(R, n)
		}
	}

	return R
}

// 912m Sort an Array
func sortArray(nums []int) []int {
	t := make([]int, len(nums)) // temporary merge storage
	copy(t, nums)

	var mSort func(s, e int, main, tmp []int)
	mSort = func(s, e int, main, tmp []int) {
		if e-s <= 1 {
			return
		}

		m := s + (e-s)>>1
		mSort(s, m, tmp, main)
		mSort(m, e, tmp, main)

		// Merge
		l, r := s, m
		for i := s; i < e; i++ {
			if l < m && (r >= e || tmp[l] <= tmp[r]) {
				main[i] = tmp[l]
				l++
			} else {
				main[i] = tmp[r]
				r++
			}
		}
	}

	mSort(0, len(nums), nums, t)
	return nums
}

// 922 Sort Array By Parity II
func sortArrayByParityII(nums []int) []int {
	R := make([]int, len(nums))

	e, o := 0, 1
	for _, n := range nums {
		switch n & 1 {
		case 0:
			R[e] = n
			e += 2
		case 1:
			R[o] = n
			o += 2
		}
	}

	return R
}

// 2191m Sort the Jumbled Numbers
func sortJumbled(mapping []int, nums []int) []int {
	// 0 <= nums[i] < 10^9
	Map := func(n int) int {
		m := 0
		for rdx := 1; n > 0; rdx *= 10 {
			m += mapping[n%10] * rdx
			n /= 10
		}
		return m
	}

	D := [][]int{}
	for i, n := range nums {
		D = append(D, []int{Map(n), nums[i]})
	}
	log.Print(" -> ", D)

	slices.SortFunc(D, func(x, y []int) int { return x[0] - y[0] })

	for i := range D {
		nums[i] = D[i][1]
	}
	return nums
}

// 2418 Sort the People
func sortPeople(names []string, heights []int) []string {
	type P struct {
		name   string
		height int
	}

	D := []*P{}
	for i := range names {
		D = append(D, &P{name: names[i], height: heights[i]})
	}
	slices.SortFunc(D, func(x, y *P) int { return y.height - x.height })

	for i := range D {
		names[i] = D[i].name
	}
	return names
}

// 2948m Make Lexicographically Smallest Array by Swapping Elements
func lexicographicallySmallestArray(nums []int, limit int) []int {
	Val := make([][2]int, 0, len(nums))
	for i, n := range nums {
		Val = append(Val, [2]int{i, n})
	}
	slices.SortFunc(Val, func(a, b [2]int) int { return a[1] - b[1] })
	Val = append(Val, [2]int{len(Val) + 1, math.MaxInt})

	log.Print(" -> ", Val)

	p := 0
	G := []int{Val[0][0]}
	R := make([]int, len(Val))

	for i := 1; i < len(Val); i++ {
		if Val[i][1] > Val[i-1][1]+limit {
			slices.Sort(G)

			var g int
			for len(G) > 0 {
				g, G = G[0], G[1:]

				R[g] = Val[p][1]
				p++
			}
		}

		G = append(G, Val[i][0])
	}

	return R[:len(R)-1]
}

// 3301m Maximize the Total Height of Unique Towers
func maximumTotalSum(maximumHeight []int) int64 {
	R := []int{}

	slices.SortFunc(maximumHeight, func(a, b int) int { return b - a })
	log.Print(" -> ", maximumHeight)

	R = append(R, maximumHeight[0])
	for i, h := range maximumHeight[1:] {
		R = append(R, min(h, R[i]-1))
	}

	t := int64(0)
	for _, h := range R {
		if h < 1 {
			return -1
		}
		t += int64(h)
	}
	return t
}
