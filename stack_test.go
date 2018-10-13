package gopher

import (
	"bytes"
	"log"
	"slices"
	"testing"
)

// 71m Simplify Path
func Test71(t *testing.T) {
	log.Print("/home ?= ", simplifyPath("/home/"))
	log.Print("/home/foo ?= ", simplifyPath("/home//foo/"))
	log.Print("/home/user/Pictures ?= ", simplifyPath("/home/user/Documents/../Pictures"))
	log.Print("/ ?= ", simplifyPath("/../"))
	log.Print("/.../b/d ?= ", simplifyPath("/.../a/../b/c/../d/./"))
}

// 316m Remove Duplicate Letters
func Test316(t *testing.T) {
	log.Print(" -> b ? abcd :: ", slices.Index([]byte("abcd"), 'b'))
	log.Print(" -> b ? abcd :: ", bytes.Index([]byte("abcd"), []byte("b")))

	log.Print("abc ?= ", removeDuplicateLetters("bcabc"))
	log.Print("acdb ?= ", removeDuplicateLetters("cbacdcbc"))
}

// 503m Next Greater Element II
func Test503(t *testing.T) {
	log.Print("[2 -1 2] ?= ", nextGreaterElements([]int{1, 2, 1}))
	log.Print("[2 3 4 -1 4] ?= ", nextGreaterElements([]int{1, 2, 3, 4, 3}))
}

// 921m Minimum Add to Make Parentheses Valid
func Test921(t *testing.T) {
	log.Print("1 ?= ", minAddToMakeValid("())"))
	log.Print("3 ?= ", minAddToMakeValid("((("))
}

// 962m Maximum Width Ramp
func Test962(t *testing.T) {
	TwoPointers := func(nums []int) int {
		N := len(nums)

		rightMax := make([]int, N)
		rightMax[N-1] = nums[N-1]
		for i := N - 1 - 1; i >= 0; i-- {
			rightMax[i] = max(rightMax[i+1], nums[i])
		}

		log.Print(" -> rightMax :: ", rightMax)

		xWid := 0

		l := 0
		for r := range nums {
			for l < r && nums[l] > rightMax[r] {
				l++
			}
			xWid = max(r-l, xWid)
		}

		return xWid
	}

	MonotonicStack := func(nums []int) int {
		Q := []int{0}
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[Q[len(Q)-1]] {
				Q = append(Q, i)
			}
		}

		xWid := 0
		for j := len(nums) - 1; j >= 0; j-- {
			for len(Q) > 0 && nums[j] >= nums[Q[len(Q)-1]] {
				xWid = max(xWid, j-Q[len(Q)-1])
				Q = Q[:len(Q)-1]
			}
		}
		return xWid
	}

	for _, fn := range []func([]int) int{maxWidthRamp, TwoPointers, MonotonicStack} {
		log.Print("4 ?= ", fn([]int{6, 0, 8, 2, 1, 5}))
		log.Print("7 ?= ", fn([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
		log.Print("--")
	}
}

// 1081m Smallest Subsequence of Distinct Characters
func Test1081(t *testing.T) {
	log.Print("abc ?= ", smallestSubsequence("bcabc"))
	log.Print("acdb ?= ", smallestSubsequence("cbacdcbc"))
}

// 1381m Design a Stack with Increment Operation
func Test1381(t *testing.T) {
	o := Constructor1381(3)

	o.Push(1)
	o.Push(2)
	log.Print("2 ?= ", o.Pop())
	o.Push(2)
	o.Push(3)
	o.Push(4)
	o.Inc(5, 100)
	o.Inc(2, 100)
	log.Print("103 ?= ", o.Pop())
	log.Print("202 ?= ", o.Pop())
	log.Print("201 ?= ", o.Pop())
	log.Print("-1 ?= ", o.Pop())
}

// 1475 Final Prices With a Special Discount in a Shop
func Test1475(t *testing.T) {
	log.Print(" ?= ", finalPrices([]int{8, 4, 6, 2, 3}))
	log.Print(" ?= ", finalPrices([]int{1, 2, 3, 4, 5}))
	log.Print(" ?= ", finalPrices([]int{10, 1, 1, 6}))
}

// 1963m Minimum Number of Swaps to Make the String Balanced
func Test1963(t *testing.T) {
	log.Print("1 ?= ", minSwapsToBalance("][]["))
	log.Print("2 ?= ", minSwapsToBalance("]]][[["))
	log.Print("0 ?= ", minSwapsToBalance("[]"))
}
