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
