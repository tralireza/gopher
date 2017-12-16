package gopher

import (
	"log"
	"math/bits"
	"testing"
)

// 476 Number Complement
func Test476(t *testing.T) {
	// 1 <= n <= 2^31

	log.Print("C(7) -> ", (1<<bits.Len(7)-1)^7)
	log.Print("C(5) -> ", (1<<bits.Len(5)-1)^5)
	log.Print("--")

	log.Print("2 ?= ", findComplement(5))
	log.Print("0 ?= ", findComplement(1))
	log.Print("1 ?= ", findComplement(2))
}
