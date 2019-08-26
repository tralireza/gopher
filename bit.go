package gopher

import "log"

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
	N := len(nums1) & len(nums2)

	Map, rMap := make([]int, N), make([]int, N)
	for i, n := range nums1 {
		Map[n] = i
	}
	for i, n := range nums2 {
		rMap[Map[n]] = i
	}

	log.Print("-> ", Map, rMap)

	fenwick := NewFenwick2197(N + 1)

	count := int64(0)
	for n := range N {
		j := rMap[n]

		left := fenwick.Query(j + 1)
		fenwick.Update(j+1, 1)

		right := N - 1 - j - (n - left)

		count += int64(left) * int64(right)
	}

	return count
}
