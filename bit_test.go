package gopher

import (
	"log"
	"math/rand/v2"
	"slices"
	"testing"
)

func Test493(t *testing.T) {
	BruteForce := func(nums []int) int {
		count := 0
		for i := 0; i < len(nums); i++ {
			for j := 0; j < i; j++ {
				if nums[j] > 2*nums[i] {
					count++
				}
			}
		}

		return count
	}

	MergeSort := func(nums []int) int {
		Merge := func(l, m, r int) {
			lArr, rArr := make([]int, m-l+1), make([]int, r-m)
			copy(lArr, nums[l:m+1])
			copy(rArr, nums[m+1:r+1])

			left, right := 0, 0
			for i := l; i <= r; i++ {
				if right == len(rArr) || left < len(lArr) && lArr[left] <= rArr[right] {
					nums[i] = lArr[left]
					left++
				} else {
					nums[i] = rArr[right]
					right++
				}
			}
		}

		var MSort func(l, r int) int
		MSort = func(l, r int) int {
			if l < r {
				m := l + (r-l)>>1
				count := MSort(l, m) + MSort(m+1, r)

				j := m + 1
				for i := l; i <= m; i++ {
					for j <= r && nums[i] > 2*nums[j] {
						j++
					}
					count += j - (m + 1)
				}

				Merge(l, m, r)

				return count
			}

			return 0
		}

		return MSort(0, len(nums)-1)
	}

	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{2, []int{1, 3, 2, 3, 1}},
		{3, []int{2, 4, 3, 5, 1}},

		{0, []int{2147483647, 2147483647, 2147483647, 2147483647, 2147483647, 2147483647}},
	} {
		log.Print("* ", c.nums)
		if c.rst != reversePairs(c.nums) {
			t.Fail()
		}
		log.Print(":: ", c.rst)
		log.Printf(":: %d (Brute Force)", BruteForce(c.nums))
		log.Printf(":: %d (Merge Sort)", MergeSort(c.nums))
	}
}

func Test2179(t *testing.T) {
	nums := []int{}
	for range 10 {
		nums = append(nums, rand.IntN(15)+1)
	}
	lFt := NewFenwick2197(slices.Max(nums) + 1)
	rFt := NewFenwick2197(slices.Max(nums) + 1)
	for _, n := range nums {
		rFt.Update(n, 1)
	}
	log.Print(nums)
	for i, n := range nums {
		rFt.Update(n, -1)
		log.Printf("<n %d   |n: %2d|   >n %d", lFt.Query(n-1), n, (len(nums)-1-i)-rFt.Query(n))
		lFt.Update(n, 1)
	}
	log.Print("---")

	for _, c := range []struct {
		rst          int64
		nums1, nums2 []int
	}{
		{1, []int{2, 0, 1, 3}, []int{0, 1, 2, 3}},
		{4, []int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}},
	} {
		if c.rst != goodTriplets(c.nums1, c.nums2) {
			t.FailNow()
		}
	}
}

func Test3480(t *testing.T) {
	for _, c := range []struct {
		rst               int64
		n                 int
		conflicting_pairs [][]int
	}{
		{9, 4, [][]int{{2, 3}, {1, 4}}},
		{12, 5, [][]int{{1, 2}, {2, 5}, {3, 5}}},
	} {
		log.Print("* ", c.n, c.conflicting_pairs)
		if c.rst != maxSubarrays(c.n, c.conflicting_pairs) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}
