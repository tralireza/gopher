package gopher

import (
	"log"
	"slices"
	"strings"
)

// 71m Simplify Path
func simplifyPath(path string) string {
	P := strings.Split(path, "/")

	log.Printf(" -> %q", P)

	Q := []string{}
	for _, p := range P {
		switch p {
		case "":
		case ".":
		case "..":
			if len(Q) > 0 {
				Q = Q[:len(Q)-1]
			}
		default:
			Q = append(Q, p)
		}
	}

	log.Printf(" -> %q", Q)

	return "/" + strings.Join(Q, "/")
}

// 316m Remove Duplicate Letters
func removeDuplicateLetters(s string) string {
	Q := []byte{}

	lPos := map[byte]int{}
	for i := range len(s) {
		lPos[s[i]] = i
	}

	Seen := map[byte]bool{}

	for i := range len(s) {
		if Seen[s[i]] {
			continue
		}

		for len(Q) > 0 && s[i] < Q[len(Q)-1] && lPos[Q[len(Q)-1]] > i {
			Seen[Q[len(Q)-1]] = false
			Q = Q[:len(Q)-1]
		}

		if !Seen[s[i]] {
			Q = append(Q, s[i])
			Seen[s[i]] = true
		}
	}

	return string(Q)
}

// 921m Minimum Add to Make Parentheses Valid
func minAddToMakeValid(s string) int {
	qSize, ops := 0, 0

	for _, l := range s {
		if l == '(' {
			qSize++ // .Push()
		} else { // )
			if qSize > 0 {
				qSize-- // .Pop()
			} else {
				ops++
			}
		}
	}

	return qSize + ops
}

// 962m Maximum Width Ramp
func maxWidthRamp(nums []int) int {
	indices := make([]int, len(nums))
	for i := range indices {
		indices[i] = i
	}

	slices.SortFunc(indices, func(x, y int) int {
		if nums[x] == nums[y] {
			return x - y
		}
		return nums[x] - nums[y]
	})

	log.Print(" -> ", indices)

	xWid := 0
	i := len(nums)
	for _, j := range indices { // Kadane's
		xWid = max(j-i, xWid)
		i = min(j, i)
	}
	return xWid
}

// 1381m Design a Stack with Increment Operation
type CustomStack1381 struct {
	Q, I []int
	top  int
}

func Constructor1381(maxSize int) CustomStack1381 {
	return CustomStack1381{
		make([]int, maxSize), make([]int, maxSize),
		-1,
	}
}

func (o *CustomStack1381) Push(x int) {
	if o.top+1 == len(o.Q) {
		return
	}
	o.top++
	o.Q[o.top] = x
}

func (o *CustomStack1381) Pop() int {
	if o.top == -1 {
		return -1
	}

	v := o.Q[o.top] + o.I[o.top]

	if o.top > 0 {
		o.I[o.top-1] += o.I[o.top]
	}
	o.I[o.top] = 0

	o.top--
	return v
}

func (o *CustomStack1381) Inc(k, v int) {
	if o.top == -1 {
		return
	}
	o.I[min(o.top, k-1)] = v
}

// 1963m Minimum Number of Swaps to Make the String Balanced
func minSwapsToBalance(s string) int {
	qSize := 0
	u := 0

	for _, l := range s {
		if l == '[' {
			qSize++
		} else { // l == ']'
			if qSize > 0 {
				qSize-- // .Pop()
			} else {
				u++
			}
		}
	}

	return (u + 1) / 2
}
