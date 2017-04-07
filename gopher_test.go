package gopher

import (
	"errors"
	"fmt"
	"io"
	"log"
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

func TestSolarySystem(t *testing.T) {
	PrintPlanets()
}

func TestErrors(t *testing.T) {
	log.Print("? ", errors.New("UserError") == errors.New("UserError"))

	type UserError struct{ msg string }
	e1, e2 := UserError{"Msg"}, UserError{"Msg"}

	log.Print("? ", e1 == e2)
	log.Print("? ", &e1 == &e2)

	log.Printf(" -> %T", fmt.Errorf("%s", "UserError"))
}
