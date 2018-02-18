package gopher

import (
	"log"
	"math"
	"slices"
	"strconv"
)

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

// 564h Find the Closest Palindrome
func nearestPalindromic(n string) string {
	Value := func(s string) int {
		v := 0
		for i := 0; i < len(s); i++ {
			v = v*10 + int(s[i]-'0')
		}
		return v
	}

	N := Value(n)

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
