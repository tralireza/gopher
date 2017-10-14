package gopher

import (
	"log"
	"testing"
)

// 1395m Count Number of Teams
func Test1395(t *testing.T) {
	log.Print("3 ?= ", numTeams([]int{2, 5, 3, 4, 1}))
	log.Print("0 ?= ", numTeams([]int{2, 1, 3}))
	log.Print("4 ?= ", numTeams([]int{1, 2, 3, 4}))
}
