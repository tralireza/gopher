package gopher

import (
	"container/heap"
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

	for i, f := range []func(int, [][]int, int) int{findTheCity, BellmanFord, SPF, FloydWarshall} {
		switch i {
		case 0:
			log.Print("3 ?= ", f(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
			log.Print("0 ?= ", f(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
		default:
			f(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4)
			f(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2)
		}
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

// 2976m Minimum Cost to Convert String I
type E2976 struct{ n, d int }
type PQ2976 []E2976

func (h PQ2976) Len() int           { return len(h) }
func (h PQ2976) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h PQ2976) Less(i, j int) bool { return h[i].d < h[j].d }
func (h *PQ2976) Push(x any)        { *h = append(*h, x.(E2976)) }
func (h *PQ2976) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func Test2976(t *testing.T) {
	type E = E2976
	type PQ = PQ2976

	WithDijkstra := func(source, target string, original, changed []byte, cost []int) int64 {
		G := [26][26]int{}
		for i := range cost {
			r, c := original[i]-'a', changed[i]-'a'
			if G[r][c] == 0 || G[r][c] < cost[i] {
				G[r][c] = cost[i]
			}
		}

		Dijkstra := func(source int, SP []int) {
			h := &PQ{}

			heap.Push(h, E{source, 0})
			SP[source] = 0

			for h.Len() > 0 {
				v := heap.Pop(h).(E).n

				for u, w := range G[v] {
					if w > 0 && w+SP[v] < SP[u] { // w = 0 => v and u are not connected/neighbors
						SP[u] = SP[v] + w
						heap.Push(h, E{u, SP[u]})
					}
				}
			}
		}

		INF := math.MaxInt>>1 - 1
		aSP := [26][26]int{} // all-Node Shortest-Path
		for r := range 26 {
			for c := range 26 {
				aSP[r][c] = INF
			}
		}

		for n := range 26 {
			Dijkstra(n, aSP[n][:])
		}

		x := int64(0)
		for i := 0; i < len(source); i++ {
			r, c := source[i]-'a', target[i]-'a'
			if aSP[r][c] == INF {
				return -1
			}
			x += int64(aSP[r][c])
		}
		return x
	}

	for _, f := range []func(string, string, []byte, []byte, []int) int64{minimumCost, WithDijkstra} {
		log.Print("28 ?= ", f("abcd", "acbe", []byte{'a', 'b', 'c', 'c', 'e', 'd'}, []byte{'b', 'c', 'b', 'e', 'b', 'e'}, []int{2, 5, 5, 1, 2, 20}))
		log.Print("12 ?= ", f("aaaa", "bbbb", []byte{'a', 'c'}, []byte{'c', 'b'}, []int{1, 2}))
		log.Print("-1 ?= ", f("abcd", "abce", []byte{'a'}, []byte{'c'}, []int{10000}))
		log.Print("--")
	}
}
