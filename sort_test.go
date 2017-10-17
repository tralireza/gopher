package gopher

import (
	"log"
	"slices"
	"testing"
)

// 56m Merge Intervals
func Test56(t *testing.T) {
	log.Print("[[1 6] [8 10] [15 18]] ?= ", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	log.Print("[[1 5]] ?= ", merge([][]int{{1, 4}, {4, 5}}))
}

// 912m Sort an Array
func Test912(t *testing.T) {
	log.Print(" ?= ", sortArray([]int{5, 2, 3, 1}))
	log.Print(" ?= ", sortArray([]int{5, 1, 1, 2, 0, 0}))
}

// 2191m Sort the Jumbled Numbers
func Test2191(t *testing.T) {
	BucketSort := func(mapping []int, nums []int) []int {
		K := 0
		for x := slices.Max(nums); x > 0; x /= 10 {
			K++
		}

		Buckets := make([][]int, 10)
		Buckets[0] = nums

		rdx := 1
		for range K {
			Tmps := make([][]int, 10)
			for b := range Buckets {
				for _, n := range Buckets[b] {
					if n < rdx && n != 0 {
						Tmps[0] = append(Tmps[0], n)
						continue
					}

					digit := (n / rdx) % 10
					md := mapping[digit]
					Tmps[md] = append(Tmps[md], n)
				}
			}
			rdx *= 10
			Buckets = Tmps
		}

		i := 0
		for b := range Buckets {
			for _, n := range Buckets[b] {
				nums[i] = n
				i++
			}
		}
		return nums
	}

	for _, f := range []func([]int, []int) []int{sortJumbled, BucketSort} {
		log.Print("[338 38 991] ?= ", f([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}))
		log.Print("[123 456 789] ?= ", f([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}))
		log.Print("--")
	}
}

// 2418 Sort the People
func Test2418(t *testing.T) {
	QuickSort := func(names []string, heights []int) []string {
		// worse-case: O(N^2)
		var qSort func(s, e int)
		qSort = func(s, e int) {
			if s >= e {
				return
			}

			// Pivot Value: element at first index :: PivotVal <- Arrry[start]
			v, p := heights[s], s
			for i := s + 1; i <= e; i++ {
				if heights[i] > v {
					heights[i], heights[p] = heights[p], heights[i]
					names[i], names[p] = names[p], names[i]
					p++
				}
			}

			qSort(s, p-1)
			qSort(p+1, e)
		}
		qSort(0, len(heights)-1)

		return names
	}

	MergeSort := func(names []string, heights []int) []string {
		// worse-case: O(NlogN)

		// Heights/Names merge temporary storage
		th, tn := make([]int, len(heights)), make([]string, len(heights))
		for i := range heights {
			th[i], tn[i] = heights[i], names[i]
		}

		var mSort func(heights []int, names []string, s, e int, th []int, tn []string)
		mSort = func(heights []int, names []string, s, e int, th []int, tn []string) {
			// s <= i < e <=> N: [s...e)
			if e-s <= 1 {
				return
			}

			m := s + (e-s)>>1
			mSort(heights, names, s, m, th, tn)
			mSort(heights, names, m, e, th, tn)

			// Merge
			l, r := s, m
			for i := s; i < e; i++ {
				if l < m && (r >= e || th[l] >= th[r]) {
					heights[i], names[i] = th[l], tn[l]
					l++
				} else {
					heights[i], names[i] = th[r], tn[r]
					r++
				}
			}
		}
		mSort(heights, names, 0, len(heights), th, tn)

		return names
	}

	WithIndex := func(names []string, heights []int) []string {
		type P struct{ i, h int }

		D := []P{}
		for i := range names {
			D = append(D, P{i: i, h: heights[i]})
		}
		slices.SortFunc(D, func(x, y P) int { return y.h - x.h })

		R := []string{}
		for i := range D {
			R = append(R, names[D[i].i])
		}
		return R
	}

	for _, f := range []func([]string, []int) []string{sortPeople, WithIndex, QuickSort, MergeSort} {
		log.Print("[Mary Emma John] ?= ", f([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
		log.Print("[Bob Alice Bob] ?= ", f([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
		log.Print("[A B C D E F G] ?= ", f([]string{"A", "F", "G", "B", "D", "C", "E"}, []int{7, 2, 1, 6, 4, 5, 3}))
		log.Print("--")
	}
}
