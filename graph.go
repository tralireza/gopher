package gopher

import (
	"container/heap"
	"log"
	"math"
	"slices"
)

// 127h Word Ladder
func ladderLength(beginWord string, endWord string, wordList []string) int {
	Mem := map[string]bool{}
	for _, w := range wordList {
		Mem[w] = true
	}

	L := len(beginWord)
	Q := []string{beginWord}
	delete(Mem, beginWord)

	t := 1
	for len(Q) > 0 { // BFS
		var v string
		for range len(Q) {
			v, Q = Q[0], Q[1:]
			if v == endWord {
				return t
			}

			for l := range L {
				for x := 'a'; x <= 'z'; x++ {
					u := v[:l] + string(x) + v[l+1:]
					if Mem[u] {
						Q = append(Q, u)
						delete(Mem, u)
					}
				}
			}
		}
		t++
	}

	return 0
}

// 210m Course Schedule II
func findOrder(numCourses int, prerequisites [][]int) []int {
	G := make([][]int, numCourses)
	for _, e := range prerequisites {
		G[e[1]] = append(G[e[1]], e[0])
	}

	D := make([]int, numCourses) // in-Degree of all N nodes :: [0..n) courses
	for v := range G {
		for _, u := range G[v] { // edge: v -> u
			D[u]++
		}
	}

	Color := make([]int, numCourses)
	tSort := []int{} // Topological-Sort data

	var DFS func(int) bool // true if cycle detected
	DFS = func(v int) bool {
		Color[v] = 1 // Visiting

		for _, u := range G[v] {
			switch Color[u] {
			case 0:
				if DFS(u) {
					return true // terminate early with cycles...
				}
			case 1:
				return true // cycle detected
			}
		}

		Color[v] = 2 // Done
		tSort = append(tSort, v)
		return false // no cycle
	}

	for n := range D {
		if D[n] == 0 {
			if DFS(n) {
				return []int{} // with cycles, no schedule
			}
		}
	}

	if len(tSort) != numCourses {
		return []int{}
	}
	slices.Reverse(tSort)
	return tSort
}

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

// 433m Minimum Genetic Mutation
func minMutation(startGene string, endGene string, bank []string) int {
	Mem := map[string]bool{}
	for _, m := range bank {
		Mem[m] = true
	}

	Q := []string{startGene}
	m := 0

	for len(Q) > 0 {
		var v string
		for range len(Q) {
			v, Q = Q[0], Q[1:]
			if v == endGene {
				return m
			}

			for i := range 8 {
				for _, x := range []string{"A", "C", "T", "G"} {
					u := v[:i] + x + v[i+1:]
					if Mem[u] {
						Q = append(Q, u)
						delete(Mem, u)
					}
				}
			}
		}
		m++
	}

	return -1
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

// 909m Snakes & Ladders
func snakesAndLadders(board [][]int) int {
	Rows, Cols := len(board), len(board[0])

	Cord := func(n int) (r, c int) {
		n--
		r = Rows - 1 - n/Rows
		if (Rows-r)&1 == 1 {
			c = n % Cols
			return
		}
		c = Cols - 1 - n%Cols
		return
	}

	Q := []int{1}
	r, c := Cord(1)
	board[r][c] = 0 // Done!

	throws := 0      // Dice throws...
	for len(Q) > 0 { // BFS
		var n int
		for range len(Q) {
			n, Q = Q[0], Q[1:]
			if n == Rows*Cols {
				return throws
			}

			for d := 1; d <= 6 && n+d <= Rows*Cols; d++ {
				r, c := Cord(n + d)
				if board[r][c] == -1 {
					Q = append(Q, n+d)
				} else if board[r][c] > 0 {
					Q = append(Q, board[r][c])
				}
				board[r][c] = 0 // Done!
			}
		}
		throws++
	}

	return -1
}

// 1192h Critical Connections in a Network
func criticalConnections(n int, connections [][]int) [][]int {
	G := make([][]int, n)
	for _, e := range connections {
		v, u := e[0], e[1]
		G[v], G[u] = append(G[v], u), append(G[u], v)
	}

	Index, Lowest := make([]int, n), make([]int, n)
	timer := 0

	var findBridge func(v, p int, rBridge func(v, u int))
	findBridge = func(v, p int, rBridge func(v, u int)) {
		timer++
		Index[v], Lowest[v] = timer, timer

		for _, u := range G[v] {
			if Index[u] == 0 { // Not visited/discovered yet...
				findBridge(u, v, rBridge)
				Lowest[v] = min(Lowest[v], Lowest[u])

				if Lowest[u] > Index[v] {
					rBridge(v, u) // Bridge edge {v, u}
				}

			} else if u != p {
				Lowest[v] = min(Lowest[v], Index[u])
			}
		}
	}

	R := [][]int{}
	findBridge(0, -1, func(v, u int) { R = append(R, []int{v, u}) }) // Node 0 as root (parent -1 -> no parent)

	return R
}

// 1591h Strange Printer II
func isPrintable(targetGrid [][]int) bool {
	Rows, Cols := len(targetGrid), len(targetGrid[0])
	Cords := map[int][2][2]int{}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			k := targetGrid[r][c]
			if cord, ok := Cords[k]; !ok {
				cords := [2][2]int{{r, c}, {r, c}}

				for x := r; x < Rows; x++ {
					for y := 0; y < Cols; y++ {
						if targetGrid[x][y] != k {
							if x > cords[1][0] {
								cords[1][0] = r
							}
							if y < cord[0][1] {
								cords[0][1] = y
							}
							if y > cord[1][1] {
								cords[1][1] = y
							}
						}
					}
				}
				Cords[k] = cords
			}

		}
	}

	log.Print("Cords -> ", Cords)

	G := map[int]map[int]struct{}{}

	for k, cords := range Cords {
		G[k] = map[int]struct{}{}
		for r := cords[0][0]; r <= cords[1][0]; r++ {
			for c := cords[0][1]; c <= cords[1][1]; c++ {
				if k != targetGrid[r][c] {
					G[k][targetGrid[r][c]] = struct{}{}
				}
			}
		}
	}

	log.Print("Graph -> ", G)

	tSort := []int{}
	Color := map[int]int{} // 0: White, 1: Gray, 2: Black
	var DFS func(v int) bool
	DFS = func(v int) bool {
		Color[v] = 1 // Visiting...
		for u := range G[v] {
			switch Color[u] {
			case 0:
				if DFS(u) {
					return true // terminate early
				}
			case 1:
				return true // cycle detected
			}
		}
		Color[v] = 2 // Done.
		tSort = append(tSort, v)
		return false // no cycle
	}

	for v := range G {
		if Color[v] == 0 {
			if DFS(v) {
				return false
			}
		}
	}

	slices.Reverse(tSort)
	log.Print("TopoSort :: ", tSort)
	return true
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
