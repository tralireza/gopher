package gopher

import (
	"log"
	"testing"
)

// 476 Number Complement
func Test476(t *testing.T) {
	// 1 <= n <= 2^31

	log.Print("2 ?= ", findComplement(5))
	log.Print("0 ?= ", findComplement(1))
	log.Print("1 ?= ", findComplement(2))
}
