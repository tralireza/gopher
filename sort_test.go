package gopher

import (
	"log"
	"slices"
	"testing"
)

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

		var mSort func(s, e int)
		mSort = func(s, e int) {
			if s >= e {
				return
			}

			m := s + (e-s)>>1
			mSort(s, m)
			mSort(m+1, e)

			// Merge
			l, r := s, m+1
			for i := s; i <= e; i++ {
				if l <= m && (r > e || heights[l] >= heights[r]) {
					th[i], tn[i] = heights[l], names[l]
					l++
				} else {
					th[i], tn[i] = heights[r], names[r]
					r++
				}
			}
			copy(heights[s:e+1], th[s:e+1])
			copy(names[s:e+1], tn[s:e+1])
		}
		mSort(0, len(heights)-1)

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
