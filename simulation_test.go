package gopher

import (
	"log"
	"testing"
)

// 6m Zigzag Conversion
func Test6(t *testing.T) {
	log.Print("PAHNAPLSIIGYIR ?= ", convert("PAYPALISHIRING", 3))
	log.Print("PINALSIGYAHRPI ?= ", convert("PAYPALISHIRING", 4))
	log.Print("A ?= ", convert("A", 1))

}

// 289m Game of Life
func Test289(t *testing.T) {
	gameOfLife([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
	gameOfLife([][]int{{1, 1}, {1, 0}})
}

// 840m Magic Squares In Grid
func Test840(t *testing.T) {
	log.Print("1 ?= ", numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
	log.Print("0 ?= ", numMagicSquaresInside([][]int{{8}}))
}

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
