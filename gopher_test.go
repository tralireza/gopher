package gopher

import (
	"bytes"
	"container/heap"
	"container/list"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"

	"golang.org/x/net/html"
)

func TestByteCounter(t *testing.T) {
	var v ByteCounter

	fmt.Fprintf(&v, "Hello Gopher!")
	if int(v) != len("Hello Gopher!") {
		t.Fail()
	}

	var _ fmt.Stringer = &v
	var _ fmt.Stringer = new(ByteCounter)
	var _ fmt.Stringer = (*ByteCounter)(nil)

	log.Printf(" -> %v", v)
}

func TestInterface(t *testing.T) {
	var z fmt.Stringer
	log.Printf("%T", z)

	var v fmt.Stringer = new(ByteCounter)
	log.Printf("fmt.Stringer -> Dynamic Type: %T    Dynamic Value: %[1]v", v)

	// v.Write([]byte{}) -> compile error
	if v, ok := v.(io.Writer); ok {
		log.Print("👍")
		v.Write([]byte("Hello Gopher!"))
		log.Printf("io.Writer -> %T %[1]v", v)
	}
}

func TestSortInterface(t *testing.T) {
	PrintPlanets()
}

func TestTypeSwitch(t *testing.T) {
	for _, x := range []interface{}{nil, uint(5), 0, true, float64(2.7182818), "ID-X10X", [...]int{0, 0}} {
		log.Printf("%T: %[1]v -> %s", x, SqlQuote(x))
	}
}

func TestErrors(t *testing.T) {
	log.Print("? ", errors.New("UserError") == errors.New("UserError"))

	type UserError struct{ msg string }
	e1, e2 := UserError{"Msg"}, UserError{"Msg"}

	log.Print("? ", e1 == e2)
	log.Print("? ", &e1 == &e2)

	log.Printf(" -> %T", fmt.Errorf("%s", "UserError"))

	var errors = [...]string{
		1: "error 1",
		3: "error 3",
		4: "error 4",
		9: "error 9",
	}
	log.Printf("%T -> %[1]q", errors)
	log.Print("? ", errors[2])

	_, err := os.Open("/fs/dir/file")
	log.Printf("%T -> %#[1]v", err)
	log.Printf("-> %T", err.(*fs.PathError))
}

func TestClosure(t *testing.T) {
	log.Print("+ ", Fib(9))
	log.Print("+ ", Fib(45))
}

func TestChannelWG(t *testing.T) {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(cin <-chan int) {
		defer wg.Done()

		for v := range cin {
			log.Print("Reading...", v)
		}
	}(ch)

	go func(cout chan<- int) {
		for range 5 {
			cout <- rand.Intn(3)
		}
		close(cout)
	}(ch)

	wg.Wait()
}

func TestChannel(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%T -> %[1]v", err)
		}
	}()

	var ch chan int
	log.Print("? ", ch == nil)
	log.Printf(" :: '%T'   R: '%v' '%v'", ch, reflect.TypeOf(ch), reflect.ValueOf(ch).Type())

	ch1, ch2 := make(<-chan int), make(chan<- int, 1)
	// log.Print("? ", ch1 == ch2) -> compile error
	log.Printf("| %T | %T |", ch1, ch2)
	ch2 <- 0
	close(ch2)

	chn := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case chn <- i:
		case n := <-chn:
			log.Print(n)
		}
	}
	close(chn)

	close(ch)
}

func TestBits(t *testing.T) {
	func() {
		nums := []int{-3, -3, -3, 1, 1, 1, 4, 4, 4, -9}
		x := 0
		for p := 31; p >= 0; p-- {
			b := 0
			for _, n := range nums {
				b += (n >> p) & 1
			}
			x |= (b % 3) << p
		}
		log.Print(nums, " -> -9 ?= ", int(int32(x)))

		y, z := int32(x), int(int32(x))
		log.Printf("| %d | %d | %d |", x, y, z)
	}()

	func() {
		nums := []int{-3, -3, 1, 1, -2, 8}
		xy := 0
		for _, n := range nums {
			xy ^= n
		}

		x, y := 0, 0
		p := 0
		for xy != 0 {
			if xy&1 == 1 {
				for _, n := range nums {
					if n>>p&1 == 1 {
						x ^= n
					} else {
						y ^= n
					}
				}
				break
			}
			xy >>= 1
			p++
		}

		log.Print(nums, " -> ? [-2 8]")
		log.Printf("| %v |", []int{x, y})
	}()
}

func TestSafety(t *testing.T) {
	var wg sync.WaitGroup

	// Immutable Safety
	M := map[int]string{0: "Zero", 1: "One", 2: "Two", 3: "Three"}
	wg = sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range 100 {
				fmt.Fprintf(io.Discard, "{%d->%s}", n%4, M[n%4])
			}
		}()
	}
	wg.Wait()

	// Mutual Exclusion
	type Task struct {
		State  string
		access chan struct{} // -> sync.Mutex
	}

	p1 := Task{"New", make(chan struct{}, 1)}
	wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range 1000 {
			p1.access <- struct{}{} // -> sync.Mutex.Lock()
			p1.State = "Running"    //
			<-p1.access             // -> sync.Mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for range 1000 {
			p1.access <- struct{}{} // -> sync.Mutex.Lock()
			p1.State = "Sleeping"   //
			<-p1.access             // -> sync.Mutex.Unlock()
		}
	}()
	wg.Wait()
	log.Print(p1)

	// Monitor
	runc, slpc := make(chan *Task), make(chan *Task)
	Run := func(p *Task) { runc <- p }
	Sleep := func(p *Task) { slpc <- p }

	quitc := make(chan struct{})
	go func() {
		for range 1000 {
			switch rand.Intn(2) {
			case 0:
				Run(&p1)
			case 1:
				Sleep(&p1)
			}
		}
		quitc <- struct{}{}
	}()

	quit := false
	for !quit {
		select {
		case p := <-runc:
			p.State = "Running"
			fmt.Print(" -> ", p)
		case p := <-slpc:
			p.State = "Sleeping"
			fmt.Print(" -> ", p)
		case <-quitc:
			quit = true
		}
	}

	log.Print("\n", p1)
}

func TestNBCacheDS(t *testing.T) {
	bfr, err := httpGet("https://go.dev")
	if err != nil {
		t.Fatal(err)
	}

	if bfr, ok := bfr.(*bytes.Buffer); ok {
		node, err := html.Parse(bfr)
		if err != nil {
			t.Fatal(err)
		}

		PageCache := NewNBCacheDS(httpGet)
		for _, l := range []string{"https://pkg.go.dev", "https://github.com/golang"} {
			go func(url string) {
				PageCache.Get(url) // cache warm up!
			}(l)
		}
		time.Sleep(time.Second)

		var wg sync.WaitGroup // zero value is also fine...
		for _, lnk := range hrefXtr([]string{}, node) {
			if strings.HasPrefix(lnk, "https://") {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					ts := time.Now()
					_, err := PageCache.Get(url)
					log.Printf(" %v {%v} -> %s ", time.Since(ts), err, url)
				}(lnk)
			}
		}
		wg.Wait()
	}
}

func TestHttpGet(t *testing.T) {
	bfr, err := httpGet("https://go.dev")
	if err != nil {
		t.Fatal(err)
	}

	if bfr, ok := bfr.(*bytes.Buffer); ok {
		node, err := html.Parse(bfr)
		if err != nil {
			t.Fatal(err)
		}

		wg := sync.WaitGroup{}
		for _, lnk := range hrefXtr([]string{}, node) {
			if strings.HasPrefix(lnk, "https://") {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					ts := time.Now()
					if _, err := httpGet(url); err == nil {
						log.Printf(" [%v] -> %s", time.Since(ts), url)
					}
				}(lnk)
			}
		}
		wg.Wait()
	}
}

func TestUnsafe(t *testing.T) {
	f := float64(1.0)
	log.Printf("%#016x\n%#016x", int64(31), *(*uint64)(unsafe.Pointer(&f)))
	log.Printf("%d", *(*uint64)(unsafe.Pointer(&f)))

	i, j := int64(9), int64(10)
	log.Printf("%p %p %#x", &i, &j, uintptr(unsafe.Pointer(&i)))
	p := (*int64)(unsafe.Pointer((uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(int64(0)))))
	log.Printf("%p", p)
	*p = 17
	log.Print(" {*p = 17} -> ", j)

	s1, s2 := []string{"-", "."}, strings.Split("-*.", "*")
	// s1 == s2 *** Not-Possible -> !comparable
	log.Printf(" ? %t", reflect.DeepEqual(s1, s2))
	log.Printf(" %v %v ? %t <!>", []string{}, []string(nil), reflect.DeepEqual([]string(nil), []string{}))
	log.Printf(" %v %v ? %t <!>", map[int]int{}, map[int]int(nil), reflect.DeepEqual(map[int]int(nil), map[int]int{}))
}

func TestBST(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	Draw := func(root *TreeNode) {
		Q := list.New()
		Q.PushBack(root)
		for Q.Len() > 0 {
			for range Q.Len() {
				n := Q.Remove(Q.Front()).(*TreeNode)
				l, r := '-', '-'
				if n.Left != nil {
					l = '*'
					Q.PushBack(n.Left)
				}
				if n.Right != nil {
					r = '*'
					Q.PushBack(n.Right)
				}
				fmt.Printf("{%c %d %c}", l, n.Val, r)
			}
			fmt.Print("\n")
		}
	}

	RotateRight := func(root *TreeNode) *TreeNode {
		root, root.Left.Right, root.Left = root.Left, root, root.Left.Right
		return root
	}

	RotateLeft := func(root *TreeNode) *TreeNode {
		root, root.Right, root.Right.Left = root.Right, root.Right.Left, root
		return root
	}

	VineRight := func(root *TreeNode) *TreeNode {
		head := &TreeNode{Right: root}

		n := head
		for n.Right != nil {
			if n.Right.Left != nil {
				n.Right = RotateRight(n.Right)
			} else {
				n = n.Right
			}
		}

		return head.Right
	}

	type T = TreeNode
	x := &T{1, &T{Val: 0}, &T{Val: 2, Right: &T{Val: 3}}}
	Draw(x)
	log.Print(" --- Rotate: Right --> ")
	x = RotateRight(x)
	Draw(x)
	log.Print(" --- Rotate: Left --> ")
	x = RotateLeft(x)
	Draw(x)

	log.Print("--")
	y := &T{4, &T{2, &T{Val: 1}, &T{Val: 3}}, &T{Val: 6, Right: &T{8, &T{Val: 7}, &T{Val: 9}}}}
	Draw(VineRight(y))
}

func TestMDM(t *testing.T) {
	var M [][][2]int
	log.Print(M)

	M = make([][][2]int, 3)
	log.Print(M)

	for r := range M {
		M[r] = make([][2]int, 4)
	}
	log.Print(M)

	M[2][3][0]++
	M[2][3][1]++
	log.Print(M)

	M[1][2] = [2]int{7, 8}
	M[1][1] = [...]int{3, 4}
	log.Print(M)

	M[0][1] = M[1][2]
	M[0][1][1]++
	M[1][2][0]++
	log.Print(M)

	M[0][0][0]++
	log.Print(M)

	type T struct{ i, j, k int }
	N := make([][][3]T, 2)
	for r := range N {
		N[r] = make([][3]T, 2)
	}
	log.Print(N)

	N[1][1][2] = T{1, 1, 2}
	copy(N[0], N[1])
	N[1][1][0] = T{1, 1, 0}
	log.Print(N)
}

// 3m Longest Substring Without Repeating Characters
func Test3(t *testing.T) {
	Distance := func(s string) int {
		Mem := [256]int{}
		for i := range Mem {
			Mem[i] = -1
		}

		ls, k := 0, -1
		for i, c := range s {
			k = max(k, Mem[c])
			ls = max(i-k, ls)
			Mem[c] = i
		}
		return ls
	}
	log.Print("3 ?= ", Distance("pwwkew"))

	log.Print("3 ?= ", lengthOfLongestSubstring("abcabcbb"))
	log.Print("1 ?= ", lengthOfLongestSubstring("bbbb"))
	log.Print("3 ?= ", lengthOfLongestSubstring("pwwkew"))
}

// 20 Valid Parentheses
func Test20(t *testing.T) {
	log.Print("true ?= ", isValid("()"))
	log.Print("true ?= ", isValid("()[]{}"))
	log.Print("false ?= ", isValid("(]"))
}

// 31m Next Permutation
func Test31(t *testing.T) {
	var nums []int

	nums = []int{1, 2, 3}
	nextPermutation(nums)
	log.Print("[1 3 2] ?= ", nums)

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	log.Print("[1 2 3] ?= ", nums)
}

// 48m Rotate Image
func Test48(t *testing.T) {
	Draw := func(M [][]int) {
		for r := range M {
			fmt.Print("[")
			for c := range M[r] {
				fmt.Printf("%2d ", M[r][c])
			}
			fmt.Print("]\n")
		}
	}

	for _, M := range [][][]int{
		{{}},
		{{1}},
		{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
	} {
		log.Print("==")
		Draw(M)
		rotate(M)
		log.Print(" -> ")
		Draw(M)
	}
}

// 49m Group Anagrams
func Test49(t *testing.T) {
	log.Printf("-> %+v", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// 73m Set Matrix Zeroes
func Test73(t *testing.T) {
	Draw := func(M [][]int) {
		for r := range M {
			log.Print(M[r])
		}
	}

	for _, M := range [][][]int{
		{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		{{0, 1, 2, 0}, {3, 4, 5, 6}, {6, 7, 8, 9}},
		{{0, 1}},
	} {
		log.Print("==")
		Draw(M)
		log.Print(" -> ")
		setZeroes(M)
		Draw(M)
	}
}

// 207m Course Schedule
func Test207(t *testing.T) {
	// 0 -> 1
	log.Print("true ?= ", canFinish(2, [][]int{{1, 0}}))

	// 0 <-> 1
	log.Print("false ?= ", canFinish(2, [][]int{{1, 0}, {0, 1}}))

	//       /------->|
	// 1 -> 0 -> 3 -> 4   5 -> 6
	//      2 -->|
	log.Print("true ?= ", canFinish(7, [][]int{{0, 4}, {1, 0}, {2, 3}, {3, 4}, {0, 3}, {5, 6}}))

	log.Print("false ?= ", canFinish(7, [][]int{{0, 4}, {4, 1}, {1, 0}, {2, 3}, {3, 4}, {0, 3}, {5, 6}}))
}

// 238m Product of Array Except Self
func Test238(t *testing.T) {
	log.Print("[24 12 8 6] ?= ", productExceptSelf([]int{1, 2, 3, 4}))
	log.Print("[0 0 9 0 0] ?= ", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

// 273h Integer to English Words
func Test273(t *testing.T) {
	log.Print(" ?= ", numberToWords(123))
	log.Print(" ?= ", numberToWords(12345))
	log.Print(" ?= ", numberToWords(1234567))

	log.Print(" ?= ", numberToWords(10))
	log.Print(" ?= ", numberToWords(1000))
	log.Print(" ?= ", numberToWords(100010))
}

// 287m Find the Duplicate Number
func Test287(t *testing.T) {
	// 1 <= nums[i] <= n

	InPlace := func(nums []int) int {
		for _, n := range nums {
			if n < 0 {
				n *= -1
			}
			if nums[n-1] < 0 {
				return n
			}
			nums[n-1] *= -1
		}
		return -1
	}

	for _, f := range []func([]int) int{findDuplicate, InPlace} {
		log.Print("2 ?= ", f([]int{1, 3, 4, 2, 2}))
		log.Print("3 ?= ", f([]int{3, 1, 3, 4, 2}))
		log.Print("3 ?= ", f([]int{3, 3, 3, 3, 3}))
		log.Print("--")
	}
}

// 350 Intersection of Two Arrays II
func Test350(t *testing.T) {
	OneMap := func(nums1 []int, nums2 []int) []int {
		M := map[int]int{}
		for _, n := range nums1 {
			M[n]++
		}

		R := []int{}
		for _, n := range nums2 {
			if M[n] > 0 {
				M[n]--
				R = append(R, n)
			}
		}

		return R
	}

	for _, f := range []func([]int, []int) []int{intersect, OneMap} {
		log.Print("[2 2] ?= ", f([]int{1, 2, 2, 1}, []int{2, 2}))
		log.Print("[4 9] ?= ", f([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
		log.Print("--")
	}
}

// 394m Decode String
func Test394(t *testing.T) {
	log.Print("3[a]2[bc] ?= ", decodeString("3[a]2[bc]"))
	log.Print("3[a2[c]] ?= ", decodeString("3[a2[c]]"))
	log.Print("2[abc]3[cd]ef ?= ", decodeString("2[abc]3[cd]ef"))
}

// 697 Degree of an Array
func Test697(t *testing.T) {
	log.Print("2 ?= ", findShortestSubArray([]int{1, 2, 2, 3, 1}))
	log.Print("6 ?= ", findShortestSubArray([]int{2, 2, 3, 1, 4, 2}))
}

// 726h Number of Atoms
func Test726(t *testing.T) {
	log.Print("H2O ?= ", countOfAtoms("H2O"))
	log.Print("H2MgO2 ?= ", countOfAtoms("Mg(OH)2"))
	log.Print("K4N2O14S4 ?= ", countOfAtoms("K4(ON(SO3)2)2"))
}

// 739m Daily Temperatures
func Test739(t *testing.T) {
	Builtin := func(temperatures []int) []int {
		r := make([]int, len(temperatures))

		S := list.New()
		for i, t := range temperatures {
			for S.Len() > 0 && t > temperatures[S.Back().Value.(int)] {
				j := S.Remove(S.Back()).(int)
				r[j] = i - j
			}
			S.PushBack(i)
		}

		return r
	}

	log.Print("[1 1 4 2 1 1 0 0] ?= ", dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	log.Print("[1 1 4 2 1 1 0 0] ?= ", Builtin([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

// 763m Partition Labels
func Test763(t *testing.T) {
	log.Print("[9 7 8] ?= ", partitionLabels("ababcbacadefegdehijhklij"))
	log.Print("[10] ?= ", partitionLabels("eccbbbbdec"))
}

// 995h Minimum Number of K Consecutive Bit Flips
func Test995(t *testing.T) {
	SpaceOptimized := func(nums []int, k int) int {
		x := 0
		fflip := 0

		for i := range nums {
			if i >= k && nums[i-k] == 9 {
				fflip ^= 1
			}

			if nums[i] == fflip {
				if i+k > len(nums) {
					return -1
				}

				fflip ^= 1
				nums[i] = 9
				x++
			}
		}

		return x
	}

	for _, f := range []func([]int, int) int{minKBitFlips, SpaceOptimized} {
		log.Print("==")
		log.Print("2 ?= ", f([]int{0, 1, 0}, 1))
		log.Print("-1 ?= ", f([]int{1, 1, 0}, 2))
		log.Print("3 ?= ", f([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
	}
}

// 1051 Height Checker
func Test1051(t *testing.T) {
	// 1 <= heights[i] <= 100

	RadixSort := func(heights []int) int {
		rSort := make([]int, len(heights))
		copy(rSort, heights)

		for digitRx := 0; digitRx <= 100/10; digitRx++ {
			Bucket := [10][]int{}

			rx := 1
			for range digitRx {
				rx *= 10
			}

			for _, h := range rSort {
				Bucket[(h/rx)%10] = append(Bucket[(h/rx)%10], h)
			}

			i := 0
			for r := range Bucket {
				for _, h := range Bucket[r] {
					if h > 0 {
						rSort[i] = h
						i++
					}
				}
			}
		}

		x := 0
		for i := range rSort {
			if rSort[i] != heights[i] {
				x++
			}
		}
		return x
	}

	for _, f := range []func([]int) int{RadixSort, heightChecker} {
		log.Print("3 ?= ", f([]int{1, 1, 4, 2, 1, 3}))
		log.Print("5 ?= ", f([]int{5, 1, 2, 3, 4}))
	}
}

// 1052 Grumpy Bookstore Owner
func Test1052(t *testing.T) {
	log.Print("16 ?= ", maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	log.Print("1 ?= ", maxSatisfied([]int{1}, []int{0}, 1))
}

// 1248m Count Number of Nice Subarrays
func Test1248(t *testing.T) {
	WithQueue := func(nums []int, k int) int {
		Q := list.New()
		x := 0

		l, gap := -1, 0
		for r := range nums {
			if nums[r]&1 == 1 {
				Q.PushBack(r)
			}
			if Q.Len() > k {
				l = Q.Remove(Q.Front()).(int)
			}
			if Q.Len() == k {
				gap = Q.Front().Value.(int) - l
				x += gap
			}
		}

		return x
	}

	SlidingWindow := func(nums []int, k int) int {
		AtMost := func(k int) int {
			x := 0
			l := 0
			for r := range nums {
				k -= nums[r] & 1
				for k < 0 { // more than k odd numbers in the Window between l & r, Shrink!
					k += nums[l] & 1
					l++
				}
				if 0 <= k {
					x += r - l + 1 // Window Size
				}
			}
			return x
		}

		return AtMost(k) - AtMost(k-1)
	}

	Optimized := func(nums []int, k int) int {
		x := 0
		l := 0
		gap := 0
		for r := range nums {
			k -= nums[r] & 1
			if k == 0 { // exactly K odd numbers in the Window
				gap = 0
				for k == 0 {
					k += nums[l] & 1
					l++
					gap++
				}
			}
			x += gap
		}
		return x
	}

	for _, f := range []func([]int, int) int{numberOfSubarrays, WithQueue, SlidingWindow, Optimized} {
		log.Print("==")
		log.Print("2 ?= ", f([]int{1, 1, 2, 1, 1}, 3))
		log.Print("0 ?= ", f([]int{2, 4, 6}, 1))
		log.Print("16 ?= ", f([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
	}
}

// 1509m Minimum Difference Between Largest and Smallest Value in Three Moves
type PQ1509 struct{ sort.IntSlice }

func (o *PQ1509) Push(any) {} // not needed, only .Init() & .Pop()
func (o *PQ1509) Pop() any {
	v := o.IntSlice[o.Len()-1]
	o.IntSlice = o.IntSlice[:o.Len()-1]
	return v
}

func Test1509(t *testing.T) {
	type PQ = PQ1509

	WithHeap := func(nums []int) int {
		if len(nums) <= 4 {
			return 0
		}

		var Ns []int
		for _, n := range nums {
			Ns = append(Ns, -n)
		}
		hX := PQ{Ns}
		heap.Init(&hX) // MaxHeap

		hM := PQ{nums} // MinHeap
		heap.Init(&hM)

		Ms, Xs := []int{}, []int{}
		for range 4 {
			Xs = append(Xs, -heap.Pop(&hX).(int))
			Ms = append(Ms, heap.Pop(&hM).(int))
		}

		m := math.MaxInt
		for i := range 4 {
			m = min(Xs[i]-Ms[3-i], m)
		}
		return m
	}

	for _, f := range []func([]int) int{minDifference, WithHeap} {
		log.Print("0 ?= ", f([]int{5, 3, 2, 4}))
		log.Print("1 ?= ", f([]int{1, 5, 0, 10, 14}))
		log.Print("0 ?= ", f([]int{3, 100, 20}))
	}
}

// 1518 Water Bottles
func Test1518(t *testing.T) {
	log.Print("13 ?= ", numWaterBottles(9, 3))
	log.Print("19 ?= ", numWaterBottles(15, 4))
}

// 1579h Remove Max Number of Edges to Keep Graph Fully Traversable
func Test1579(t *testing.T) {
	log.Print("2 ?= ", maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}))
	log.Print("0 ?= ", maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}))
	log.Print("-1 ?= ", maxNumEdgesToRemove(4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}))
}

// 1598 Crawler Log Folder
func Test1598(t *testing.T) {
	Counter := func(logs []string) int {
		q := 0
		for i := range logs {
			switch logs[i] {
			case "./":
			case "../":
				if q > 0 {
					q--
				}
			default:
				q++
			}
		}
		return q
	}

	minOperations := func(logs []string) int {
		Q := list.New()
		for i := range logs {
			switch logs[i] {
			case "./":
			case "../":
				if Q.Len() > 0 {
					Q.Remove(Q.Front())
				}
			default:
				Q.PushBack(logs[i])
			}
		}
		return Q.Len()
	}

	for _, f := range []func([]string) int{minOperations, Counter} {
		log.Print("2 ?= ", f([]string{"d1/", "d2/", "../", "d21/", "./"}))
		log.Print("3 ?= ", f([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
		log.Print("--")
	}
}

// 1701m Average Waiting Time
func Test1701(t *testing.T) {
	log.Print("5 ?= ", averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}}))
	log.Print("3.25 ?= ", averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}))
}

// 1717m Maximum Score From Removing Substrings
func Test1717(t *testing.T) {
	TwoPointer := func(s string, x, y int) int {
		g := 0

		Pass := func(tkn string, v int) {
			log.Printf("-> %s %s:%d", s, tkn, v)
			w := []byte(s)
			wtr := 0
			for rdr := 0; rdr < len(s); rdr++ {
				w[wtr] = w[rdr]
				wtr++

				if wtr > 1 && w[wtr-2] == tkn[0] && w[wtr-1] == tkn[1] {
					g += v
					wtr -= 2
				}
			}
			s = string(w[:wtr])
		}

		Tx, Ty := "ab", "ba"
		if y > x {
			Tx, Ty = Ty, Tx
			x, y = y, x
		}

		Pass(Tx, x)
		Pass(Ty, y)

		return g
	}

	for _, f := range []func(s string, x, y int) int{maximumGain, TwoPointer} {
		log.Print("19 ?= ", f("cdbcbbaaabab", 4, 5))
		log.Print("20 ?= ", f("aabbaaxybbaabb", 5, 4))
		log.Print("--")
	}
}

// 1823m Find the Winner of the Circular Game
func Test1823(t *testing.T) {
	// Josephus Problem
	// f(n, k) = (f(n-1, k) + k) mod n, f(1, k) = 0
	Josephus := func(n, k int) int {
		var j func(n, k int) int
		j = func(n, k int) int {
			if n == 1 {
				return 0
			}
			return (j(n-1, k) + k) % n
		}

		return j(n, k) + 1 // [0..n-1]
	}
	log.Print("Josephus Problem (n, k): (6, 4) ?= ", Josephus(6, 4))

	Iterative := func(n, k int) int {
		j := 0
		for i := 2; i <= n; i++ {
			j = (j + k) % i
		}
		return j + 1
	}

	for _, f := range []func(int, int) int{findTheWinner, Josephus, Iterative} {
		log.Print("3 ?= ", f(5, 2))
		log.Print("1 ?= ", f(6, 5))
		log.Print("--")
	}
}

// 1190m Reverse Substrings Between Each Pair of Parentheses
func Test1190(t *testing.T) {
	Recursive := func(s string) string {
		i := 0

		var R func() string
		R = func() string {
			bfr := []byte{}

			for i < len(s) {
				switch s[i] {
				case '(':
					i++
					bfr = append(bfr, R()...)

				case ')':
					i++
					l, r := 0, len(bfr)-1
					for l < r {
						bfr[l], bfr[r] = bfr[r], bfr[l]
						l++
						r--
					}
					return string(bfr)

				default:
					bfr = append(bfr, s[i])
					i++
				}
			}

			return string(bfr)
		}

		return R()
	}

	// O(n) time complexity
	TwoPass := func(s string) string {
		Q := []int{}
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case ')':
				Q = append(Q, i)
			}
		}

		P := make([]int, len(s))
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				P[i] = Q[len(Q)-1]
				P[P[i]] = i

				Q = Q[:len(Q)-1]
			}
		}
		log.Print(P)

		w := []byte{}

		dir := 1
		for i := 0; i < len(s); {
			switch s[i] {
			case '(':
				i = P[i]
				dir *= -1
			case ')':
				i = P[i]
				dir *= -1

			default:
				w = append(w, s[i])
			}

			i += dir
		}

		return string(w)
	}

	for _, f := range []func(string) string{reverseParentheses, Recursive, TwoPass} {
		log.Print("dcba ?= ", f("(abcd)"))
		log.Print("iloveu ?= ", f("(u(love)i)"))
		log.Print("leetcode ?= ", f("(ed(et(oc))el)"))
		log.Print("--")
	}
}

// 2058m Find the Minimum and Maximum Number of Nodes Between Critical Points
func Test2058(t *testing.T) {
	Draw := func(n *ListNode) {
		for n != nil {
			fmt.Printf("{%d ", n.Val)
			n = n.Next
			if n != nil {
				fmt.Print("*}->")
			} else {
				fmt.Print("/}")
			}
		}
	}

	type L = ListNode

	for _, l := range []*L{
		{3, &L{Val: 1}},
		{5, &L{3, &L{1, &L{2, &L{5, &L{1, &L{Val: 2}}}}}}},
		{1, &L{3, &L{2, &L{2, &L{3, &L{2, &L{2, &L{2, &L{Val: 7}}}}}}}}},
	} {
		Draw(l)
		fmt.Print("\n")
		log.Print(" ?= ", nodesBetweenCriticalPoints(l))
		log.Print("--")
	}
}

// 2181m Merge Nodes in Between Zeros
func Test2181(t *testing.T) {
	Iterative := func(head *ListNode) *ListNode {
		h := &ListNode{}

		r := h
		for n := head; n.Next != nil; {
			n = n.Next

			v := &ListNode{}
			r.Next = v
			for n.Val != 0 {
				v.Val += n.Val
				n = n.Next
			}
			r = v
		}

		return h.Next
	}

	Draw := func(n *ListNode) {
		for n != nil {
			fmt.Printf("{%d ", n.Val)
			n = n.Next
			if n != nil {
				fmt.Print("*}->")
			} else {
				fmt.Print("/}")
			}
		}
	}

	type L = ListNode

	for _, f := range []func(*ListNode) *ListNode{mergeNodes, Iterative} {
		for _, l := range []*L{
			{0, &L{3, &L{1, &L{0, &L{4, &L{5, &L{2, &L{Val: 0}}}}}}}},
			{0, &L{1, &L{0, &L{3, &L{0, &L{2, &L{2, &L{Val: 0}}}}}}}},
			{0, &L{1, &L{Val: 0}}},
		} {
			Draw(l)
			fmt.Print("  =>  ")
			Draw(f(l))
			fmt.Print("\n")
		}
		log.Print("--")
	}
}

// 2192m All Ancestors of a Node in a Directed Acyclic Graph
func Test2192(t *testing.T) {
	Optimized := func(n int, edges [][]int) [][]int {
		G := make([][]int, n)
		for _, e := range edges {
			G[e[0]] = append(G[e[0]], e[1])
		}

		R := make([][]int, n)
		var Vis []bool

		var DFS func(a, v int)
		DFS = func(a, v int) {
			for _, u := range G[v] {
				if !Vis[u] {
					R[u] = append(R[u], a)
					Vis[u] = true
					DFS(a, u)
				}
			}
		}

		for v := range n {
			Vis = make([]bool, n)
			DFS(v, v)
		}
		return R
	}

	TopologicalSort := func(n int, edges [][]int) []int {
		// Kahn's
		G, D := make([][]int, n), make([]int, n)
		for _, e := range edges {
			G[e[0]] = append(G[e[0]], e[1])
			D[e[1]]++
		}

		Q := []int{} // In-degree of 0
		for v, d := range D {
			if d == 0 {
				Q = append(Q, v)
			}
		}

		log.Print("+ Degree -> ", D)
		topOrder := []int{}
		var v int
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			topOrder = append(topOrder, v)
			for _, u := range G[v] {
				D[u]--
				if D[u] == 0 {
					Q = append(Q, u)
				}
			}
		}
		log.Print("+ Degree -> ", D)
		return topOrder
	}

	log.Print(" -> ", TopologicalSort(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))

	for _, f := range []func(int, [][]int) [][]int{getAncestors, Optimized} {
		log.Print("--")
		log.Print("[[] [] [] [0 1] [0 2] [0 1 3] [0 1 2 3 4] [0 1 2 3]] ?= ", f(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))
		log.Print("[[] [0] [0 1] [0 1 2] [0 1 2 3]] ?= ", f(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}))
	}
}

// 2285m Maximum Total Importance of Roads
func Test2285(t *testing.T) {
	log.Print("43 ?= ", maximumImportance(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}))
	log.Print("20 ?= ", maximumImportance(5, [][]int{{0, 3}, {2, 4}, {1, 3}}))
}

// 2496 Maximum Value of a String in an Array
func Test2496(t *testing.T) {
	Scanner := func(strs []string) int {
		Value := func(s string) int {
			v := 0
			for i := 0; i < len(s); i++ {
				if s[i] <= '9' && '0' <= s[i] {
					v = 10*v + int(s[i]-'0')
				} else {
					return len(s)
				}
			}
			return v
		}

		x := 0
		for _, s := range strs {
			x = max(Value(s), x)
		}
		return x
	}

	for _, f := range []func([]string) int{maximumValue, Scanner} {
		log.Print("5 ?= ", f([]string{"alic3", "bob", "3", "4", "00000"}))
		log.Print("1 ?= ", f([]string{"1", "01", "001", "0001"}))
		log.Print("--")
	}
}

// 2751h Robot Collisions
func Test2751(t *testing.T) {
	log.Print("[2 17 9 15 10] ?= ", survivedRobotsHealths([]int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}, "RRRRR"))
	log.Print("[14] ?= ", survivedRobotsHealths([]int{3, 5, 2, 6}, []int{10, 10, 15, 12}, "RLRL"))
	log.Print("[] ?= ", survivedRobotsHealths([]int{1, 2, 5, 6}, []int{10, 10, 11, 11}, "RLRL"))
	log.Print("[49 11] ?= ", survivedRobotsHealths([]int{3, 40}, []int{49, 11}, "LL"))
}

// 3191m Minimum Operations to Make Binary Array Elements Equal to One I
func Test3191(t *testing.T) {
	SpaceOptimized := func(nums []int) int {
		fflip, k := 0, 3

		x := 0
		for i := range nums {
			if i-k >= 0 && nums[i-k] == 9 {
				fflip ^= 1
			}

			if nums[i] == fflip {
				if i+k > len(nums) {
					return -1
				}

				x++
				fflip ^= 1
				nums[i] = 9 // 9: flag
			}
		}
		log.Print(nums)

		return x
	}

	for _, f := range []func([]int) int{minOperations, SpaceOptimized} {
		log.Print("==")
		log.Print("3 ?= ", f([]int{0, 1, 1, 1, 0, 0}))
		log.Print("-1 ?= ", f([]int{0, 1, 1, 1}))
	}
}
