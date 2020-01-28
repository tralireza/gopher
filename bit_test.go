package gopher

import (
	"log"
	"math/rand/v2"
	"slices"
	"testing"
)

func Test493(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{2, []int{1, 3, 2, 3, 1}},
		{3, []int{2, 4, 3, 5, 1}},
	} {
		log.Print("* ", c.nums)
		if c.rst != reversePairs(c.nums) {
			t.Fail()
		}
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
