package gopher

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"reflect"
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

// 1550 Three Consecutive Odds
func Test1550(t *testing.T) {
	log.Println("false ?= ", threeConsecutiveOdds([]int{2, 6, 4, 1}))
	log.Println("true ?= ", threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
}

// 1579h Remove Max Number of Edges to Keep Graph Fully Traversable
func Test1579(t *testing.T) {
	log.Print("2 ?= ", maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}))
	log.Print("0 ?= ", maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}))
	log.Print("-1 ?= ", maxNumEdgesToRemove(4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}))
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
