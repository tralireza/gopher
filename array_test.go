package gopher

import (
	"log"
	"testing"
)

// 485 Max Consecutive Ones
func Test485(t *testing.T) {
	log.Print("3 ?= ", findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	log.Print("2 ?= ", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

// 1437 Check If All 1's Are at Least Length K Places Away
func Test1437(t *testing.T) {
	log.Print("true ?= ", kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	log.Print("false ?= ", kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

// 1752 Check If Array Is Sorted and Rotated
func Test1752(t *testing.T) {
	log.Print("true ?= ", check([]int{3, 4, 5, 1, 2}))
	log.Print("false ?= ", check([]int{2, 1, 3, 4}))
	log.Print("true ?= ", check([]int{1, 2, 3}))
}

func Test3169(t *testing.T) {
	for _, c := range []struct {
		rst, days int
		meetings  [][]int
	}{
		{2, 10, [][]int{{5, 7}, {1, 3}, {9, 10}}},
		{1, 5, [][]int{{2, 4}, {1, 3}}},
		{0, 6, [][]int{{1, 6}}},
	} {
		rst, days, meetings := c.rst, c.days, c.meetings
		if rst != countDays(days, meetings) {
			t.FailNow()
		}
		log.Printf(":: %d <- %d / %v", rst, days, meetings)
	}
}
