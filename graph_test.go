package gopher

import (
	"log"
	"math"
	"slices"
	"testing"
)

// 1334m Find the City With the Smallest Number of Neighbors at a Threshold Distance
func Test1334(t *testing.T) {
	// 1 <= Weight_i <= 10^4

	BellmanFord := func(n int, edges [][]int, distanceThreshold int) int {
		aSP := make([][]int, n)
		for r := range aSP {
			aSP[r] = make([]int, n)
			for c := range aSP[r] {
				aSP[r][c] = math.MaxInt
			}
		}

		BellmanFord := func(s int, SP []int) {
			SP[s] = 0
			for range n - 1 {
				earlyTerminate := true

				for _, e := range edges {
					v, u, w := e[0], e[1], e[2]
					if SP[v] != math.MaxInt && SP[v]+w < SP[u] {
						SP[u] = SP[v] + w
						earlyTerminate = false
					}
					if SP[u] != math.MaxInt && SP[u]+w < SP[v] {
						SP[v] = SP[u] + w
						earlyTerminate = false
					}
				}

				if earlyTerminate {
					break
				}
			}
		}

		for s := range n {
			BellmanFord(s, aSP[s])
		}

		log.Print("Bellman-Ford: ", aSP)

		return 0
	}

	// Shortest-Path-First
	SPF := func(n int, edges [][]int, distanceThreshold int) int {
		G := make([][]int, n)
		for r := range G {
			G[r] = make([]int, n)
		}

		for _, e := range edges {
			v, u, w := e[0], e[1], e[2]
			G[v][u] = w
			G[u][v] = w
		}

		aSP := make([][]int, n)
		for r := range aSP {
			aSP[r] = make([]int, n)
			for c := range aSP[r] {
				aSP[r][c] = math.MaxInt
			}
		}

		SPF := func(s int, SP []int) {
			UCount := make([]int, n) // Node/Vertex update counter -> cycle detection

			SP[s] = 0
			Q := []int{s}
			var v int
			for len(Q) > 0 {
				v, Q = Q[0], Q[1:]
				for u, w := range G[v] {
					if w > 0 && SP[v]+w < SP[u] {
						SP[u] = SP[v] + w
						Q = append(Q, u)

						UCount[u]++
						if UCount[u] > n {
							log.Print("Negative Weight cycle detected on Vertex: ", u)
						}
					}
				}
			}
		}

		for s := range n {
			SPF(s, aSP[s])
		}
		log.Print("SPF: ", aSP)

		return 0
	}

	FloydWarshall := func(n int, edges [][]int, distanceThreshold int) int {
		// assumption: MinInt/2 - 1 < Weight_i < MaxInt/2 - 1

		aSP := make([][]int, n)
		for r := range aSP {
			aSP[r] = make([]int, n)
			for c := range aSP[r] {
				aSP[r][c] = math.MaxInt>>1 - 1 // preventing overflow in summation of min() in main loop
			}
		}
		for i := range aSP {
			aSP[i][i] = 0
		}

		for _, e := range edges {
			v, u, w := e[0], e[1], e[2]
			aSP[v][u] = w
			aSP[u][v] = w
		}

		for k := range n {
			for r := range n {
				for c := range n {
					aSP[r][c] = min(aSP[r][c], aSP[r][k]+aSP[k][c])
				}
			}
		}

		log.Print("Floyd-Warshall: ", aSP)

		return 0
	}

	for _, f := range []func(int, [][]int, int) int{findTheCity, BellmanFord, SPF, FloydWarshall} {
		log.Print("3 ?= ", f(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
		log.Print("0 ?= ", f(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
		log.Print("--")
	}
}

// 2392h Build a Matrix With Conditions
func Test2392(t *testing.T) {
	// 1 <= Node(labels) <= k

	TopoOrder := func(k int, edges [][]int) []int {
		G := make([][]int, k+1)
		for _, e := range edges {
			G[e[0]] = append(G[e[0]], e[1])
		}

		Color := make([]int, k+1) // Visited:Color of a Node: 0: white, 1: gray, 2: black
		Color[0] = -1             // ignore Node at label 0

		tSort := []int{}
		var DFS func(int) bool
		DFS = func(v int) bool {
			Color[v] = 1 // Gray as in Visiting...
			for _, u := range G[v] {
				switch Color[u] {
				case 0:
					if DFS(u) {
						return true
					}
				case 1:
					return true // cycle detected
				}
			}
			Color[v] = 2 // Black as in Done
			tSort = append(tSort, v)
			return false
		}

		for n := range Color {
			if Color[n] == 0 {
				if DFS(n) {
					return []int{}
				}
			}
		}

		slices.Reverse(tSort)
		return tSort
	}
	log.Print("Topological Order (DFS): ", TopoOrder(3, [][]int{{1, 2}, {3, 2}}))
	log.Print("Topological Order (DFS) [cycle]: ", TopoOrder(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))

	log.Print("[[3 0 0] [0 0 1] [0 2 0]] ?= ", buildMatrix(3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}}))
	log.Print("[] ?= ", buildMatrix(3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}))
}
