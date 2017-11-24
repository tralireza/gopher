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
			if d <= SP[v] { // closest neighbor to source
				for u, w := range G[v] { // relaxing of all neighbors to source if possible
					if w > 0 && d+w < SP[u] {
						SP[u] = d + w
						heap.Push(h, E{u, SP[u]})
					}
				}
			}
		}
	}

	for s := range n {
		Dijkstra(s, aSP[s])
	}
	log.Print("Dijkstra: ", aSP)

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

// 595m Regions Cut By Slashes
func regionsBySlashes(grid []string) int {
	for r := range grid {
		log.Printf(`<- "%s"`, grid[r])
	}

	// Graph transformation
	// / -> - - *  | \ -> * - -
	//      - * -  |      - * -
	//      * - -  |      - - *

	G := make([][]byte, 3*len(grid))
	for r := range G {
		G[r] = make([]byte, 3*len(grid[0]))
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			for x := range 9 {
				v := grid[r][c]
				if v == '/' && x%3+x/3 == 2 || v == '\\' && x%3 == x/3 {
					G[3*r+x%3][3*c+x/3] = '*'
				} else {
					G[3*r+x%3][3*c+x/3] = ' '
				}
			}
		}
	}

	for r := range G {
		log.Printf("-> %c", G[r])
	}

	regions := 0

	// Graph connectivity (DFS/BFS)
	Rows, Cols := len(G), len(G[0])
	dirs := []int{0, 1, 0, -1, 0}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if G[r][c] == ' ' {
				Q := [][]int{{r, c}}
				G[r][c] = byte('a' + regions%26)

				var v []int
				for len(Q) > 0 {
					v, Q = Q[0], Q[1:]
					for d := range 4 {
						x, y := v[0]+dirs[d], v[1]+dirs[d+1]
						if x >= 0 && x < Rows && y >= 0 && y < Cols && G[x][y] == ' ' {
							G[x][y] = byte('a' + regions%26)
							Q = append(Q, []int{x, y})
						}
					}
				}

				regions++
			}
		}
	}

	for r := range G {
		for c := range G[r] {
			if G[r][c] == '*' {
				G[r][c] = ' '
			}
		}
		log.Printf("-- %c", G[r])
	}

	return regions
}

// 733 Flood Fill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	Rows, Cols := len(image), len(image[0])

	var Color func(r, c, baseColor int)
	Color = func(r, c, baseColor int) {
		if image[r][c] == color {
			return
		}

		image[r][c] = color

		D := []int{0, 1, 0, -1, 0}
		for d := range 4 {
			x, y := r+D[d], c+D[d+1]
			if x >= 0 && x < Rows && y >= 0 && y < Cols && image[x][y] == baseColor {
				Color(x, y, baseColor)
			}
		}
	}

	Color(sr, sc, image[sr][sc])
	return image
}

// 1192h Critical Connections in a Network
func criticalConnections(n int, connections [][]int) [][]int {
	G := make([][]int, n)
	for _, e := range connections {
		v, u := e[0], e[1]
		G[v], G[u] = append(G[v], u), append(G[u], v)
	}

	R := [][]int{}

	Index, Lowest := make([]int, n), make([]int, n)
	timer := 0

	var findBridge func(v, p int)
	findBridge = func(v, p int) {
		timer++
		Index[v], Lowest[v] = timer, timer

		for _, u := range G[v] {
			if Index[u] == 0 { // Not visited/discovered yet...
				findBridge(u, v)
				Lowest[v] = min(Lowest[v], Lowest[u])

				if Lowest[u] > Index[v] {
					R = append(R, []int{v, u}) // Bridge edge
				}

			} else if u != p {
				Lowest[v] = min(Lowest[v], Index[u])
			}
		}
	}

	findBridge(0, -1) // Node 0 as root (parent -1 -> no parent)

	return R
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

// 2976m Minimum Cost to Convert String I
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	// 1 <= Cost <= 10^6
	const INF = math.MaxInt>>1 - 1 // Overflow guard for Floyd-Warshall loop
	G := [26][26]int{}
	for r := range G {
		for c := range G[r] {
			if r != c { // cost of letter src:X -> tgt:X :: 0, ie, no conversation needed
				G[r][c] = INF
			}
		}
	}

	for i := range cost {
		r, c := original[i]-'a', changed[i]-'a'
		G[r][c] = min(G[r][c], cost[i])
	}

	// Floyd-Warshall
	for k := range 26 {
		for r := range 26 {
			for c := range 26 {
				G[r][c] = min(G[r][c], G[r][k]+G[k][c])
			}
		}
	}

	x := int64(0)
	for i := 0; i < len(source); i++ {
		r, c := source[i]-'a', target[i]-'a'
		if G[r][c] == INF {
			return -1
		}
		x += int64(G[r][c])
	}
	return x
}
