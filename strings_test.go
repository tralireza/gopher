package gopher

import (
	"log"
	"testing"
)

// 1813m Sentence Similarity III
func Test1813(t *testing.T) {
	log.Print("true ?= ", areSentencesSimilar("Hello Jane", "Hello my name is Jane"))
	log.Print("false ?= ", areSentencesSimilar("of", "A lot of words"))
	log.Print("true ?= ", areSentencesSimilar("Eating right now", "Eating"))
}
