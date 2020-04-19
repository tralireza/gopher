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

// 503m Next Greater Element II
func nextGreaterElements(nums []int) []int {
	R := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		R[i] = -1

		for j := 1; j < len(nums); j++ {
			if nums[(i+j)%len(nums)] > nums[i] {
				R[i] = nums[(i+j)%len(nums)]
				break
			}
		}
	}

	return R
}

// 589 N-ary Tree Preorder Traversal
type Node589 struct {
	Val      int
	Children []*Node589
}

func preorder(root *Node589) []int {
	type N = Node589

	Iterative := func(root *N) []int {
		if root == nil {
			return []int{}
		}

		W := []int{}

		Q := []*N{root}
		var n *N
		for len(Q) > 0 {
			n, Q = Q[len(Q)-1], Q[:len(Q)-1]
			W = append(W, n.Val)

			for i := len(n.Children) - 1; i >= 0; i-- {
				if n.Children[i] != nil {
					Q = append(Q, n.Children[i])
				}
			}
		}

		return W
	}
	log.Print(":? ", Iterative(root))

	var Walk func(*N) []int
	Walk = func(n *N) []int {
		if n == nil {
			return []int{}
		}

		W := []int{n.Val}
		for _, c := range n.Children {
			W = append(W, Walk(c)...)
		}
		return W
	}

	return Walk(root)
}

// 844 Backtrace String Compare
func backspaceCompare(s string, t string) bool {
	Walk := func(s string) string {
		Q := []byte{}
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '#':
				if len(Q) > 0 {
					Q = Q[:len(Q)-1]
				}
			default:
				Q = append(Q, s[i])
			}
		}

		return string(Q)
	}

	return Walk(s) == Walk(t)
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

// 1021 Remove Outermost Parentheses
func removeOuterParentheses(s string) string {
	Prims := []string{}

	start, stack := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack++
		case ')':
			stack--
			if stack == 0 {
				Prims = append(Prims, s[start+1:i])
				start = i + 1
			}
		}
	}

	return strings.Join(Prims, "")
}

// 1081m Smallest Subsequence of Distinct Characters
func smallestSubsequence(s string) string {
	lPos := [26]int{}
	for i := range len(s) {
		lPos[s[i]-'a'] = i
	}

	Q := []byte{}

	for i := range len(s) {
		if slices.Index(Q, s[i]) == -1 {
			for len(Q) > 0 && s[i] < Q[len(Q)-1] && lPos[Q[len(Q)-1]-'a'] > i {
				Q = Q[:len(Q)-1]
			}
			Q = append(Q, s[i])
		}
	}

	return string(Q)
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

// 1475 Final Prices With a Special Discount in a Shop
func finalPrices(prices []int) []int {
	R := make([]int, len(prices))
	copy(R, prices)

	Q := []int{}
	for i := 0; i < len(prices); i++ {
		for len(Q) > 0 && prices[Q[len(Q)-1]] >= prices[i] {
			R[Q[len(Q)-1]] -= prices[i]
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, i)
	}

	return R
}

// 1910m Remove all Occurrences of a Substring
func removeOccurrences(s string, part string) string {
	S := []byte{}

	for i := 0; i < len(s); i++ {
		S = append(S, s[i])

		count := 0
		for t := len(S) - 1; t >= max(0, len(S)-len(part)); t-- {
			if part[len(part)-1-count] == S[t] {
				count++
				continue
			}
		}

		if count == len(part) {
			S = S[:len(S)-len(part)]
		}
	}

	return string(S)
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
