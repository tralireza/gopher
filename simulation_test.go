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

// 68h Text Justification
func Test68(t *testing.T) {
	log.Printf("-> %q", fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	log.Printf("-> %q", fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	log.Printf("-> %q", fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}

// 289m Game of Life
func Test289(t *testing.T) {
	gameOfLife([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
	gameOfLife([][]int{{1, 1}, {1, 0}})
}

// 592m Fraction Addition and Subtraction
func Test592(t *testing.T) {
	log.Print("0/1 ?= ", fractionAddition("-1/2+1/2"))
	log.Print("1/3 ?= ", fractionAddition("-1/2+1/2+1/3"))
	log.Print("-1/6 ?= ", fractionAddition("1/3-1/2"))
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

// 860 Lemonade Change
func Test860(t *testing.T) {
	log.Print("true ?= ", lemonadeChange([]int{5, 5, 5, 10, 20}))
	log.Print("false ?= ", lemonadeChange([]int{5, 5, 10, 10, 20}))
}

// 1380 Lucky Numbers in a Matrix
func Test1380(t *testing.T) {
	log.Print("[15] ?= ", luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	log.Print("[12] ?= ", luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
	log.Print("[7] ?= ", luckyNumbers([][]int{{7, 8}, {1, 2}}))
}

// 2022 Convert 1D Array Into 2D Array
func Test2022(t *testing.T) {
	log.Print("[[1 2] [3 4]] ?= ", construct2DArray([]int{1, 2, 3, 4}, 2, 2))
	log.Print("[[1 2 3]] ?= ", construct2DArray([]int{1, 2, 3}, 1, 3))
	log.Print("[] ?= ", construct2DArray([]int{1, 2}, 1, 1))
}

// 3001m Minimum Moves to Capture the Queen
func Test3001(t *testing.T) {
	log.Print("2 ?= ", minMovesToCaptureTheQueen(1, 1, 8, 8, 2, 3))
	log.Print("1 ?= ", minMovesToCaptureTheQueen(5, 3, 3, 4, 5, 2))
}
