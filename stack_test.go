package gopher

import (
	"log"
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

// 921m Minimum Add to Make Parentheses Valid
func Test921(t *testing.T) {
	log.Print("1 ?= ", minAddToMakeValid("())"))
	log.Print("3 ?= ", minAddToMakeValid("((("))
}

// 962m Maximum Width Ramp
func Test962(t *testing.T) {
	log.Print("4 ?= ", maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	log.Print("7 ?= ", maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
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

// 1963m Minimum Number of Swaps to Make the String Balanced
func Test1963(t *testing.T) {
	log.Print("1 ?= ", minSwapsToBalance("][]["))
	log.Print("2 ?= ", minSwapsToBalance("]]][[["))
	log.Print("0 ?= ", minSwapsToBalance("[]"))
}
