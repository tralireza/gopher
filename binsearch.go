package gopher

import (
	"log"
	"math"
	"slices"
	"strconv"
)

// 153m Find Minimum in Rotated Sorted Array
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		log.Print(l, m, r, nums)

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

// 154h Find Minimum in Rotated Sorted Array II
func findMinII(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		log.Print(l, m, r, nums)

		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r++
		}

	}

	return nums[l]
}

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

// 492 Construct the Rectangle
func constructRectangle(area int) []int {
	x, w := 1, 1
	for (x+1)*(x+1) <= area {
		x++
		if area%x == 0 {
			w = x
		}
	}

	return []int{area / w, w}
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

	N := Value(n)

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

// 704 Binary Search
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else {
			if nums[m] == target {
				return m
			}
			r = m - 1
		}
	}
	return -1
}

// 1760m Minimum Limit of Balls in a Bag
func minimumSize(nums []int, maxOperations int) int {
	Possible := func(m int) bool {
		ops := 0
		for _, n := range nums {
			if n > m {
				ops += (n - 1) / m
			}
			if ops > maxOperations {
				return false
			}
		}
		return true
	}

	l, r := 1, slices.Max(nums)
	for l < r {
		m := l + (r-l)>>1
		log.Print(" -> ", l, m, r, Possible(m))

		if Possible(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// 1894m Find the Student that Will Replace the Chalk
func chalkReplacer(chalk []int, k int) int {
	pSum := make([]int, len(chalk))
	pSum[0] = chalk[0]
	for i, n := range chalk[1:] {
		pSum[i+1] = n + pSum[i]
	}
	k %= pSum[len(chalk)-1]

	l, r := 0, len(chalk)-1
	for l < r {
		m := l + (r-l)>>1
		if pSum[m] <= k {
			l = m + 1
		} else {
			r = m
		}
	}
	return r
}

// 2529 Maximum Count of Positive Integer and Negative Integer
func maximumCount(nums []int) int {
	BSLeft := func(t int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] < t {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}

	BSRight := func(t int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] > t {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	return max(len(nums)-BSRight(0), BSLeft(0))
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

// 3296m Minimum Number of Seconds to Make Mountain Height Zero
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	Check := func(m int) bool {
		hCur := mountainHeight

		for i, t := range workerTimes {
			x := 1
			for t <= m {
				hCur--
				if hCur == 0 {
					return true
				}
				x++
				t += x * workerTimes[i]
			}
		}

		return false
	}

	l, r := 0, math.MaxInt
	for l < r {
		m := l + (r-l)>>1
		if Check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l)
}
