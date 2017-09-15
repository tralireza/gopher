package gopher

import (
	"log"
	"testing"
)

// 2096m Step-By-Step Directions From a Binary Tree Node to Another
func Test2096(t *testing.T) {
	type T = TreeNode

	log.Print("UURL ?= ", getDirections(&T{5, &T{1, &T{Val: 3}, nil}, &T{2, &T{Val: 6}, &T{Val: 4}}}, 3, 6))
	log.Print("L ?= ", getDirections(&T{2, &T{Val: 1}, nil}, 2, 1))
}
