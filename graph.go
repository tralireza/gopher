package gopher

import (
	"container/heap"
	"log"
	"math"
)

// 1334m Find the City With the Smallest Number of Neighbors at a Threshold Distance
type E1334 struct{ n, distance int }
type PQ1334 []E1334

func (h PQ1334) Len() int               { return len(h) }
func (h PQ1334) Less(i int, j int) bool { return h[i].distance < h[j].distance }
func (h PQ1334) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ1334) Push(x any)            { *h = append(*h, x.(E1334)) }
func (h *PQ1334) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	G := make([][]int, n)
	for r := range G {
		G[r] = make([]int, n)
	}

	for _, e := range edges {
		v, u, w := e[0], e[1], e[2]
		G[v][u] = w
		G[u][v] = w
	}

	log.Print("(weighted) Graph: ", G)

	// (all) Shortest-Path between Cities
	aSP := make([][]int, n)
	for r := range aSP {
		aSP[r] = make([]int, n)
		for c := range aSP[r] {
			aSP[r][c] = math.MaxInt
		}
	}

	type E = E1334
	type PQ = PQ1334

	Dijkstra := func(s int, SP []int) {
		h := &PQ{}
		heap.Push(h, E{s, 0})
		SP[s] = 0

		for h.Len() > 0 {
			e := heap.Pop(h).(E)
			v, d := e.n, e.distance
			if d > SP[v] {
				continue
			}

			// closest neighbor to source
			for u, w := range G[v] { // relaxing of all neighbors to source if possible
				if w > 0 && d+w < SP[u] {
					SP[u] = d + w
					heap.Push(h, E{u, SP[u]})
				}
			}
		}
	}

	for s := range n {
		Dijkstra(s, aSP[s])
	}
	log.Print(aSP)

	city, reachables := -1, n
	for v := range n {
		t := 0
		for u, distance := range aSP[v] {
			if u != v && distance <= distanceThreshold {
				t++
			}
		}
		if t <= reachables {
			city, reachables = v, t
		}
	}
	return city
}

// 2392h Build a Matrix With Conditions
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	TopoSort := func(edges [][]int) []int {
		G := make([][]int, k+1) // Graph (DAG)
		D := make([]int, k+1)   // InDegree
		for _, e := range edges {
			G[e[0]] = append(G[e[0]], e[1])
			D[e[1]]++
		}

		// exclusive Node 0
		log.Print("Graph :: ", G)
		log.Print("InDegree :: ", D)

		// Kahn's algorithm
		Q := []int{}
		for v, degree := range D[1:] {
			if degree == 0 {
				Q = append(Q, v+1)
			}
		}
		tSort := []int{}
		var v int
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			tSort = append(tSort, v)
			for _, u := range G[v] {
				D[u]--
				if D[u] == 0 {
					Q = append(Q, u)
				}
			}
		}
		return tSort
	}

	rTopo := TopoSort(rowConditions)
	if len(rTopo) < k { // cycle :: G !DAG
		return nil
	}
	log.Print("row Topological order -> ", rTopo)

	cTopo := TopoSort(colConditions)
	if len(cTopo) < k { // cycle :: G !DAG
		return nil
	}
	log.Print("column Topological order -> ", cTopo)

	M := make([][]int, k)
	for r := 0; r < k; r++ {
		M[r] = make([]int, k)
		for c := 0; c < k; c++ {
			if rTopo[r] == cTopo[c] {
				M[r][c] = rTopo[r]
			}
		}
	}
	return M
}
