package gopher

import (
	"log"
	"slices"
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
