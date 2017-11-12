package gopher

import (
	"log"
	"testing"
)

// 885m Spiral Matrix III
func Test885(t *testing.T) {
	log.Print(" ?= ", spiralMatrixIII(1, 4, 0, 0))
	log.Print(" ?= ", spiralMatrixIII(5, 6, 1, 4))
}

// 1380 Lucky Numbers in a Matrix
func Test1380(t *testing.T) {
	log.Print("[15] ?= ", luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	log.Print("[12] ?= ", luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
	log.Print("[7] ?= ", luckyNumbers([][]int{{7, 8}, {1, 2}}))
}
