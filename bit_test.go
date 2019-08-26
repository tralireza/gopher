package gopher

import "testing"

func Test2179(t *testing.T) {
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
