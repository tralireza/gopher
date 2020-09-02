package gopher

import (
	"log"
	"math"
)

// Rotary Lock (Chapter 2)
func getMinCodeEntryTime(N int32, M int32, C []int32) int64 {
	curState := map[[2]int32]int64{}

	curState[[2]int32{0, 0}] = 0
	for _, code := range C {
		code--

		nState := map[[2]int32]int64{}
		for u, dist := range curState {
			for _, u := range [][2]int32{[2]int32{u[0], u[1]}, [2]int32{u[1], u[0]}} {
				v := [2]int32{min(code, u[1]), max(code, u[1])}
				weight := min((N+u[0]-code)%N, (N+code-u[0])%N)

				if distV, ok := nState[v]; ok {
					if dist+int64(weight) < distV {
						nState[v] = dist + int64(weight)
					}
				} else {
					nState[v] = dist + int64(weight)
				}
			}
		}

		log.Printf("-> %d(+1) : %v -> %v", code, curState, nState)

		curState = nState
	}

	t := int64(math.MaxInt64)
	for _, w := range curState {
		t = min(t, w)
	}

	return t
}
