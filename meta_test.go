package gopher

import (
	"log"
	"testing"
)

func Test_RotaryLock2(t *testing.T) {
	for _, c := range []struct {
		rst  int64
		N, M int32
		C    []int32
	}{
		{2, 3, 3, []int32{1, 2, 3}},
		{6, 10, 4, []int32{9, 4, 4, 8}},
	} {
		log.Print("* ", c.N, c.M, c.C)
		if c.rst != getMinCodeEntryTime(c.N, c.M, c.C) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}
