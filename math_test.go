package gopher

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// 319m Bulb Switcher
func Test319(t *testing.T) {
	log.Print("1 ?= ", bulbSwitch(3))
	log.Print("0 ?= ", bulbSwitch(0))
	log.Print("1 ?= ", bulbSwitch(1))

	log.Print("3 ?= ", bulbSwitch(9))
	log.Print("5 ?= ", bulbSwitch(27))
}

// 326 Power of Three
func Test326(t *testing.T) {
	log.Print("true ? ", isPowerOfThree(27))
	log.Print("false ? ", isPowerOfThree(0))
	log.Print("false ? ", isPowerOfThree(-1))
	log.Print("true ? ", isPowerOfThree(1))
}

func Test335(t *testing.T) {
	for _, c := range []struct {
		rst      bool
		distance []int
	}{
		{true, []int{2, 1, 1, 2}},
		{false, []int{1, 2, 3, 4}},
		{true, []int{1, 1, 1, 2, 1}},

		{true, []int{1, 1, 2, 1, 1}},
		{true, []int{1, 1, 2, 2, 1, 1}},
		{false, []int{3, 3, 3, 2, 1, 1}},
	} {
		if c.rst != isSelfCrossing(c.distance) {
			t.FailNow()
		}
	}
}

// 342 Power of Four
func Test342(t *testing.T) {
	log.Print("true ?= ", isPowerOfFour(16))
	log.Print("false ?= ", isPowerOfFour(5))
	log.Print("true ?= ", isPowerOfFour(1))
}

// 598 Range Addition II
func Test598(t *testing.T) {
	log.Print("4 ?= ", maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
	log.Print("4 ?= ", maxCount(3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}))
	log.Print("9 ?= ", maxCount(3, 3, [][]int{}))
}

func Test838(t *testing.T) {
	for _, c := range []struct {
		rst, dominoes string
	}{
		{"RR.L", "RR.L"},
		{"LL.RR.LLRRLL..", ".L.R...LR..L.."},
	} {
		if c.rst != pushDominoes(c.dominoes) {
			t.FailNow()
		}
	}
}

// 908 Smallest Range I
func Test908(t *testing.T) {
	log.Print("0 ?= ", smallestRangeI([]int{1}, 0))
	log.Print("6 ?= ", smallestRangeI([]int{0, 10}, 2))
	log.Print("0 ?= ", smallestRangeI([]int{1, 3, 6}, 3))
}

func Test970(t *testing.T) {
	// 1 <= x, y <= 100
	// 0 <= bound <= 10^6

	for _, c := range []struct {
		rst         []int
		x, y, bound int
	}{
		{[]int{2, 3, 4, 5, 7, 9, 10}, 2, 3, 10},
		{[]int{2, 4, 6, 8, 10, 14}, 3, 5, 15},

		{[]int{2, 3, 5, 9}, 2, 1, 10},
	} {
		if !reflect.DeepEqual(c.rst, powerfulIntegers(c.x, c.y, c.bound)) {
			t.FailNow()
		}
		log.Printf(":: %v   <- [%d %d] %d", c.rst, c.x, c.y, c.bound)
	}
}

// 989 Add to Array-Form of Integer
func Test989(t *testing.T) {
	log.Print("[1 2 3 4] ?= ", addToArrayForm([]int{1, 2, 0, 0}, 34))
	log.Print("[4 5 5] ?= ", addToArrayForm([]int{2, 7, 4}, 181))
	log.Print("[1 0 2 1] ?= ", addToArrayForm([]int{2, 1, 5}, 806))
}

// 1154 Day of the Year
func Test1154(t *testing.T) {
	log.Print("9 ?= ", dayOfYear("2019-01-09"))
	log.Print("41 ?= ", dayOfYear("2019-02-10"))
}

func Test1295(t *testing.T) {
	for _, c := range []struct {
		rst  int
		nums []int
	}{
		{2, []int{12, 345, 2, 6, 7896}},
		{1, []int{555, 901, 482, 1771}},
	} {
		if c.rst != findNumbers(c.nums) {
			t.FailNow()
		}
	}
}

func Test1432(t *testing.T) {
	for _, c := range []struct {
		rst, num int
	}{
		{888, 555},
		{8, 9},

		{820000, 123456},
		{888, 111},         // 205/211
		{8808050, 1101057}, // 207/211

	} {
		log.Print("* ", c.num)
		if c.rst != maxDiff(c.num) {
			t.Errorf("~ %d", c.rst)
		}
	}
}

func Benchmark1432(b *testing.B) {
	for range b.N {
		maxDiff(18808050)
	}
}

func Benchmark1432_Strings(b *testing.B) {
	Strings := func(num int) int {
		vMin, vMax := strconv.Itoa(num), strconv.Itoa(num)

		for i := 0; i < len(vMax); i++ {
			if vMax[i] != '9' {
				vMax = strings.ReplaceAll(vMax, string(vMax[i]), "9")
				break
			}
		}

		for i := 0; i < len(vMin); i++ {
			if vMin[i] != '0' && vMin[i] != '1' {
				vMin = strings.ReplaceAll(vMin, string(vMin[i]), "0")
				break
			}
		}

		x, _ := strconv.Atoi(vMax)
		m, _ := strconv.Atoi(vMin)
		return x - m
	}

	for range b.N {
		Strings(18808050)
	}
}

// 1780m Check if Number is a Sum of Powers of Three
func Test1780(t *testing.T) {
	// 1 <= N <= 10^7

	// O(logN)
	Ternary := func(n int) bool {
		for n > 0 {
			if n%3 == 2 {
				return false
			}
			n /= 3
		}

		return true
	}

	// O(2^log3(N))
	Recursive := func(n int) bool {
		var Search func(p, n int) bool
		Search = func(p, n int) bool {
			if n == 0 {
				return true
			}
			if n < p {
				return false
			}

			return Search(3*p, n) || Search(3*p, n-p)
		}

		return Search(1, n)
	}

	for _, f := range []func(int) bool{checkPowersOfThree, Ternary, Recursive} {
		log.Print("true ?= ", f(12))
		log.Print("true ?= ", f(91))
		log.Print("false ?= ", f(21))
		log.Print("--")
	}
}

func Test1922(t *testing.T) {
	for _, c := range []struct {
		rst int
		n   int64
	}{
		{5, 1},
		{400, 4},
		{564908303, 50},

		{711414395, 1000_000_000_000_000},
	} {
		if c.rst != countGoodNumbers(c.n) {
			t.FailNow()
		}
		log.Printf(":: %d   <- %d", c.rst, c.n)
	}
}

// 1998h GCD Sort of an Array
func Test1998(t *testing.T) {
	// 2 <= N_i <= 10^5, N.length <= 3*10^4

	log.Println("true ?= ", gcdSort([]int{7, 21, 3}))
	log.Println("false ?= ", gcdSort([]int{5, 2, 6, 2}))
	log.Println("true ?= ", gcdSort([]int{10, 5, 9, 3, 15}))
}

func Test2523(t *testing.T) {
	for _, c := range []struct {
		rst         []int
		left, right int
	}{
		{[]int{11, 13}, 10, 19},
		{[]int{-1, -1}, 4, 6},

		{[]int{29, 31}, 19, 31},
	} {
		log.Print("* ", c.left, c.right)
		if !reflect.DeepEqual(c.rst, closestPrimes(c.left, c.right)) {
			t.FailNow()
		}
	}
}

func Test2566(t *testing.T) {
	for _, c := range []struct {
		rst, num int
	}{
		{99009, 11891},
		{99, 90},

		{999, 999},
		{9, 0},
	} {
		if c.rst != minMaxDifference(c.num) {
			t.FailNow()
		}
	}
}

// 2579m Count Total Number of Colored Cells
func Test2579(t *testing.T) {
	log.Print("1 ?= ", coloredCells(1))
	log.Print("5 ?= ", coloredCells(2))
}

func Test2843(t *testing.T) {
	for _, c := range []struct {
		rst, low, high int
	}{
		{9, 1, 100},
		{4, 1200, 1230},
	} {
		if c.rst != countSymmetricIntegers(c.low, c.high) {
			t.FailNow()
		}
		log.Printf(":: %d   <- (%d, %d)", c.rst, c.low, c.high)
	}
}

func Test2894(t *testing.T) {
	for _, c := range []struct {
		rst, n, m int
	}{
		{19, 10, 3},
		{15, 5, 6},
		{-15, 5, 1},
	} {
		if c.rst != differenceOfSums(c.n, c.m) {
			t.FailNow()
		}
	}
}

func Test2929(t *testing.T) {
	for _, c := range []struct {
		rst      int64
		n, limit int
	}{
		{int64(50025003), 10001, 20001}, // (TLE) 500/958

		{int64(3), 5, 2},
		{int64(10), 3, 3},
	} {
		log.Print("* ", c.n, c.limit)
		if c.rst != distributeCandiesII(c.n, c.limit) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test3024(t *testing.T) {
	for _, c := range []struct {
		rst  string
		nums []int
	}{
		{"equilateral", []int{3, 3, 3}},
		{"scalene", []int{3, 4, 5}},
	} {
		if c.rst != triangleType(c.nums) {
			t.FailNow()
		}
	}
}

// 3151 Special Array I
func Test3151(t *testing.T) {
	log.Print("true ?= ", isArraySpecialI([]int{1}))
	log.Print("true ?= ", isArraySpecialI([]int{2, 1, 4}))
	log.Print("false ?= ", isArraySpecialI([]int{4, 3, 1, 6}))
}

func Test3272(t *testing.T) {
	for _, c := range []struct {
		rst  int64
		n, k int
	}{
		{27, 3, 5},
		{2, 1, 4},
		{2468, 5, 6},

		{9, 2, 1},
		{41457024, 10, 1},
	} {
		if c.rst != countGoodIntegers(c.n, c.k) {
			t.FailNow()
		}
		log.Printf(":: %d   <- %d %d", c.rst, c.n, c.k)
	}
}

// 3312h Sorted GCD Pair Queries
func Test3312(t *testing.T) {
	// 1 <= N_i <= 5*10^4, N.length <= 10^5

	log.Print("[1 2 2] ?= ", gcdValues([]int{2, 3, 4}, []int64{0, 2, 2}))
	log.Print("[4 2 1 1] ?= ", gcdValues([]int{4, 4, 2, 1}, []int64{5, 3, 1, 0}))
	log.Print("[2 2] ?= ", gcdValues([]int{2, 2}, []int64{0, 0}))
}
