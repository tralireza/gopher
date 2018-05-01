package gopher

import (
	"log"
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
