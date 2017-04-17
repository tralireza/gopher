package gopher

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"reflect"
	"sync"
	"testing"
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
