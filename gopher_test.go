package gopher

import (
	"bytes"
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
	log.Print(" -> ", j)
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
