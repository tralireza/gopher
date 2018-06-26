package gopher

import (
	"log"
	"testing"
)

// 412 Fizz Buzz
func Test412(t *testing.T) {
	log.Print(" ?= ", fizzBuzz(3))
	log.Print(" ?= ", fizzBuzz(5))
	log.Print(" ?= ", fizzBuzz(15))
}

// 1813m Sentence Similarity III
func Test1813(t *testing.T) {
	log.Print("true ?= ", areSentencesSimilar("Hello Jane", "Hello my name is Jane"))
	log.Print("false ?= ", areSentencesSimilar("of", "A lot of words"))
	log.Print("true ?= ", areSentencesSimilar("Eating right now", "Eating"))
}

// 2405m Optimal Partition of String
func Test2405(t *testing.T) {
	log.Print("4 ?= ", partitionString("abacaba"))
	log.Print("6 ?= ", partitionString("ssssss"))
}
