package gopher

import (
	"log"
	"testing"
)

// 2418 Sort the People
func Test2418(t *testing.T) {
	log.Print("[Mary Emma John] ?= ", sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	log.Print("[Bob Alice Bob] ?= ", sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}
