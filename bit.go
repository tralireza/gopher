package gopher

import (
	"log"
	"slices"
)

// 493h Reverse Pairs
func reversePairs(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	slices.Sort(sorted)

	lBSearch := func(t int) int {
		l, r := 0, len(sorted)
		for l < r {
			m := l + (r-l)>>1
			if sorted[m] < t {
				l = m + 1
			} else {
				r = m
			}
		}

		return l
	}

	fw := make([]int, len(nums)+1)
	Update := func(i, diff int) {
		for i > 0 {
			fw[i] += diff
			i -= i & (^i + 1)
		}
	}
	Query := func(i int) int {
		v := 0
		for i < len(fw) {
			v += fw[i]
			i += i & (^i + 1)
		}
		return v
	}

	count := 0
	for _, n := range nums {
		log.Print("-> Fenwick: ", fw)

		count += Query(lBSearch(2*n+1) + 1)
		Update(lBSearch(n)+1, 1)
	}

	return count
}

// 2179h Count Good Triplets in an Array
func NewFenwick2197(size int) BIT2197 {
	return make(BIT2197, size)
}

type BIT2197 []int

func (fw *BIT2197) Update(i int, diff int) {
	for i < len(*fw) {
		(*fw)[i] += diff
		i += i & -i
	}
}
func (fw *BIT2197) Query(i int) int {
	v := 0
	for i > 0 {
		v += (*fw)[i]
		i -= i & -i
	}
	return v
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	log.Print("** ", nums1, nums2)

	N := len(nums1) & len(nums2) // same length

	Map := make([]int, N)
	{
		Tmp := make([]int, N)
		for i, n := range nums2 {
			Tmp[n] = i
		}
		for i, n := range nums1 {
			Map[Tmp[n]] = i
		}
	}

	fenwick := NewFenwick2197(N + 1)

	count := int64(0)
	for i := range N {
		j := Map[i] // j_2: Post of n in nums2

		left := fenwick.Query(j + 1)
		fenwick.Update(j+1, 1)

		right := N - 1 - j - (i - left)

		log.Printf("-> [j1: %d  |n: %d|  j2: %d]  %d*%d", i, nums1[i], j, left, right)

		count += int64(left) * int64(right)
	}

	return count
}
