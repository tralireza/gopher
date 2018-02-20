package gopher

import (
	"log"
	"testing"
)

// 212h Word Search II
func Test212(t *testing.T) {
	log.Printf(`["oath" "eat"] ?= %q`, findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
	log.Printf(`[] ?= %q`, findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcd"}))
}
