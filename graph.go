package gopher

import (
	"container/heap"
	"container/list"
	"log"
	"math"
	"slices"
)

// 126h Word Ladder II
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	Mem := map[string]bool{}
	for _, w := range wordList {
		Mem[w] = true
	}

	Prev := map[string][]string{}

	Dist := map[string]int{}
	Dist[beginWord] = 0

	Q := []string{beginWord}
	var v string
	for len(Q) > 0 {
		v, Q = Q[0], Q[1:]

		if Mem[v] {
			delete(Mem, v)
		}

		for u := range Mem {
			diff := 0
			for i := range len(beginWord) {
				if v[i] != u[i] {
					diff++
				}
			}
			if diff > 1 {
				continue // v & u :: not-adjacent
			}

			if Dist[u] == 0 || Dist[u] > Dist[v]+1 {
				Dist[u] = Dist[v] + 1
				Prev[u] = append(Prev[u], v)
				Q = append(Q, u)
				continue
			}

			if Dist[u] == Dist[v]+1 {
				Prev[u] = append(Prev[u], v)
			}
		}
	}

	log.Print("Dist :: ", Dist)
	log.Print("Prev :: ", Prev)

	if Dist[endWord] == 0 {
		return [][]string{}
	}

	R := [][]string{}
	var BackTrack func(string, []string)
	BackTrack = func(v string, r []string) {
		if v == beginWord {
			t := append([]string{}, r...)
			slices.Reverse(t)
			R = append(R, t)
			return
		}
		for _, u := range Prev[v] {
			r = append(r, u)
			BackTrack(u, r)
			r = r[:len(r)-1]
		}
	}
	BackTrack(endWord, []string{endWord})
	return R
}

// 127h Word Ladder
func ladderLength(beginWord string, endWord string, wordList []string) int {
	Mem := map[string]bool{}
	for _, w := range wordList {
		Mem[w] = true
	}

	Q := []string{beginWord}
	delete(Mem, beginWord)

	t := 1           // BFS layers
	for len(Q) > 0 { // BFS
		var v string
		for range len(Q) {
			v, Q = Q[0], Q[1:]
			if v == endWord {
				return t
			}

			for l := range len(beginWord) {
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

// 695m Max Area of Island
func maxAreaOfIsland(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])

	var Island func(r, c int) int
	Island = func(r, c int) int {
		if r < 0 || r >= Rows || c < 0 || c >= Cols || grid[r][c] == 0 {
			return 0
		}

		grid[r][c] = 0
		return 1 + Island(r+1, c) + Island(r-1, c) + Island(r, c+1) + Island(r, c-1)
	}

	xArea := 0
	for r := range Rows {
		for c := range Cols {
			xArea = max(Island(r, c), xArea)
		}
	}
	return xArea
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

// 802m Find Eventual Safe States
func eventualSafeNodes(graph [][]int) []int {
	N := len(graph)
	Vis, S := make([]bool, N), make([]bool, N)

	var DFS func(int) bool
	DFS = func(v int) bool {
		if S[v] {
			return true
		}
		if Vis[v] {
			return false
		}

		Vis[v], S[v] = true, true
		for _, u := range graph[v] {
			if DFS(u) {
				return true
			}
		}

		S[v] = false
		return false
	}

	for v := range N {
		DFS(v)
	}

	R := []int{}
	for v := range N {
		if !S[v] {
			R = append(R, v)
		}
	}
	return R
}

// 827h Making a Large Island
type DJSet827 struct {
	Parent []int
	Size   []int
}

func (o DJSet827) Find(v int) int {
	if v != o.Parent[v] {
		o.Parent[v] = o.Find(o.Parent[v])
	}
	return o.Parent[v]
}
func (o DJSet827) Union(v, u int) bool {
	v, u = o.Find(v), o.Find(u)
	if u == v {
		return false
	}

	if o.Size[u] > o.Size[v] {
		v, u = u, v
	}
	o.Parent[u] = v
	o.Size[v] += o.Size[u]
	return true
}

func largestIsland(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])

	djset := DJSet827{
		Parent: make([]int, Rows*Cols),
		Size:   make([]int, Rows*Cols),
	}
	for i := range Rows * Cols {
		djset.Parent[i], djset.Size[i] = i, 1
	}

	Dir := []int{-1, 0, 1, 0, -1}
	for r := range Rows {
		for c := range Cols {
			if grid[r][c] == 1 {
				cellId := r*Cols + c
				for i := range 4 {
					r, c := r+Dir[i], c+Dir[i+1]
					if 0 <= r && r < Rows && 0 <= c && c < Cols && grid[r][c] == 1 {
						djset.Union(cellId, r*Cols+c)
					}
				}
			}
		}
	}

	log.Print(" -> ", djset)

	xArea := 0
	wCells := true // Water Cells

	for r := range Rows {
		for c := range Cols {
			if grid[r][c] == 0 {
				wCells = true

				M := map[int]struct{}{}
				for i := range 4 {
					r, c := r+Dir[i], c+Dir[i+1]
					if 0 <= r && r < Rows && 0 <= c && c < Cols && grid[r][c] == 1 {
						M[djset.Find(r*Cols+c)] = struct{}{}
					}
				}

				area := 1
				for root := range M {
					area += djset.Size[root]
				}

				xArea = max(xArea, area)
			}
		}
	}

	if !wCells {
		return Rows * Cols
	}
	return xArea
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

// 947m Most Stones Removed with Same Row or Column
func removeStones(stones [][]int) int {
	G := make([][]int, len(stones))
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] { // same Row or Column
				G[i], G[j] = append(G[i], j), append(G[j], i)
			}
		}
	}

	n := 0 // number of connected components

	Vis := make([]bool, len(G)) // number of Nodes in Graph
	for v := range len(G) {
		if !Vis[v] {
			Vis[v] = true
			n++

			Q := []int{v} // BFS
			for len(Q) > 0 {
				v, Q = Q[0], Q[1:]
				for _, u := range G[v] {
					if !Vis[u] {
						Vis[u] = true
						Q = append(Q, u)
					}
				}
			}
		}
	}

	return len(stones) - n
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

// 1267m Count Servers that Communicate
func countServers(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])

	xCount, yCount := make([]int, Rows), make([]int, Cols)
	for r := range Rows {
		for c := range Cols {
			if grid[r][c] == 1 {
				xCount[r]++
				yCount[c]++
			}
		}
	}

	tServers := 0
	for r := range Rows {
		for c := range Cols {
			if grid[r][c] == 1 && (xCount[r] > 1 || yCount[c] > 1) {
				tServers++
			}
		}
	}

	return tServers
}

// 1368h Minimum Cost to Make at Least One Valid Path in a Grid
func minCost(grid [][]int) int {
	const MAX = 1000_000_000
	Rows, Cols := len(grid), len(grid[0])

	Cost := make([][]int, Rows)
	for r := range Cost {
		Cost[r] = make([]int, Cols)
		for c := range Cost[r] {
			Cost[r][c] = MAX
		}
	}

	cost := 0
	Dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	Q := [][2]int{}

	var DFS func(r, c, cost int)
	DFS = func(r, c, cost int) {
		if 0 <= r && r < Rows && 0 <= c && c < Cols && Cost[r][c] == MAX {
			Cost[r][c] = cost

			Q = append(Q, [2]int{r, c})

			dir := Dirs[grid[r][c]-1]
			DFS(r+dir[0], c+dir[1], cost)
		}
	}

	DFS(0, 0, cost)

	var coord [2]int
	for len(Q) > 0 {
		cost++
		for range len(Q) {
			coord, Q = Q[0], Q[1:]

			r, c := coord[0], coord[1]
			for _, dir := range Dirs {
				DFS(r+dir[0], c+dir[1], cost)
			}
		}
	}

	return Cost[Rows-1][Cols-1]
}

// 1462m Course Schedule IV
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	FW := make([][]bool, numCourses) // Floyd-Warshall
	for r := range FW {
		FW[r] = make([]bool, numCourses)
	}

	for d := range numCourses {
		FW[d][d] = true
	}
	for _, e := range prerequisites {
		FW[e[0]][e[1]] = true
	}

	for src := range numCourses {
		for dst := range numCourses {
			for via := range numCourses {
				FW[src][dst] = FW[src][dst] || FW[src][via] && FW[via][dst]
			}
		}
	}

	R := make([]bool, 0, numCourses)
	for _, q := range queries {
		R = append(R, FW[q[0]][q[1]])
	}
	return R
}

// 1514m Path with Maximum Probability
type E1514 struct {
	n int
	w float64
}
type PQ1514 []E1514

func (h PQ1514) Len() int           { return len(h) }
func (h PQ1514) Less(i, j int) bool { return h[i].w > h[j].w } // MaxHeap
func (h PQ1514) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ1514) Push(x any)        { *h = append(*h, x.(E1514)) }
func (h *PQ1514) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// Dijkstra

	type E = E1514
	type PQ = PQ1514

	G := make([][]E, n)
	for i, e := range edges {
		v, u := e[0], e[1]
		G[v], G[u] = append(G[v], E{u, succProb[i]}), append(G[u], E{v, succProb[i]})
	}

	Dist := make([]float64, n)
	Dist[start_node] = 1 // ie, Probability

	h := PQ{}
	heap.Push(&h, E{start_node, 1})
	for h.Len() > 0 {
		log.Print(" (PQ) -> ", h)

		v := heap.Pop(&h).(E)
		for _, u := range G[v.n] {
			if Dist[u.n] < Dist[v.n]*u.w {
				Dist[u.n] = Dist[v.n] * u.w
				heap.Push(&h, E{u.n, Dist[u.n]})
			}
		}
	}

	return Dist[end_node]
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

// 1765m Map of Highest Peak
func highestPeak(isWater [][]int) [][]int {
	Rows, Cols := len(isWater), len(isWater[0])
	M, Vis := make([][]int, Rows), make([][]bool, Rows)
	for r := range M {
		M[r] = make([]int, Cols)
		Vis[r] = make([]bool, Cols)
	}

	Q := [][3]int{}
	for r := range Rows {
		for c := range Cols {
			if isWater[r][c] == 1 {
				Q = append(Q, [3]int{r, c, 0})
				Vis[r][c] = true
			}
		}
	}

	Dirs := []int{-1, 0, 1, 0, -1}

	var v [3]int
	for len(Q) > 0 {
		v, Q = Q[0], Q[1:]
		for i := range 4 {
			r, c, h := v[0]+Dirs[i], v[1]+Dirs[i+1], v[2]
			if 0 <= r && r < Rows && 0 <= c && c < Cols && !Vis[r][c] {
				Vis[r][c] = true
				M[r][c] = h + 1
				Q = append(Q, [3]int{r, c, h + 1})
			}
		}
	}

	return M
}

// 1905m Count Sub Islands
func countSubIslands(grid1, grid2 [][]int) int {
	Rows, Cols := len(grid1), len(grid1[0])

	Dir := []int{0, 1, 0, -1, 0}

	var Islands func(r, c int) bool
	Islands = func(r, c int) bool {
		grid2[r][c] = -1

		isSub := true // Sub-Island flag...
		for i := range 4 {
			x, y := r+Dir[i], c+Dir[i+1]
			if x >= 0 && x < Rows && y >= 0 && y < Cols && grid2[x][y] == 1 {
				if grid1[x][y] != 1 {
					isSub = false
				}
				if !Islands(x, y) {
					isSub = false
				}
			}
		}
		return isSub
	}

	t := 0
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if grid1[r][c] == 1 && grid2[r][c] == 1 {
				if Islands(r, c) {
					t++
				}
			}
		}
	}
	return t
}

// 1976m Number of Ways to Arrive at Destination
type E1976 struct{ Node, Dist, hSeq int } // Node, Distance, HeapSeq(i)
type PQ1976 []E1976

func (h PQ1976) Len() int           { return len(h) }
func (h PQ1976) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h PQ1976) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].hSeq, h[j].hSeq = i, j
}
func (h *PQ1976) Push(x any) {
	e := x.(E1976)
	e.hSeq = h.Len()
	*h = append(*h, e)
}
func (h *PQ1976) Pop() any {
	e := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	e.hSeq = -1
	return e
}

func countPaths(n int, roads [][]int) int {
	type E = E1976
	type PQ = PQ1976

	M := 1000_000_007

	G := make([][]E, n)
	for _, e := range roads {
		v, u, w := e[0], e[1], e[2]
		G[v], G[u] = append(G[v], E{u, w, -1}), append(G[u], E{v, w, -1})
	}

	Count := make([]int, n)
	Count[0] = 1

	Dist := make([]int, n)
	for i := range Dist {
		Dist[i] = math.MaxInt
	}

	Prev := make([][]int, n)

	h := PQ{}

	heap.Push(&h, E{Node: 0, Dist: 0})
	Dist[0] = 0

	for h.Len() > 0 {
		log.Print("-> PQ: ", h)
		v := heap.Pop(&h).(E)
		if Dist[v.Node] < v.Dist {
			continue
		}

		for _, u := range G[v.Node] {
			if Dist[u.Node] > Dist[v.Node]+u.Dist {
				Dist[u.Node] = Dist[v.Node] + u.Dist
				u.Dist = Dist[u.Node]
				heap.Push(&h, u)
				Count[u.Node] = Count[v.Node]

			} else if Dist[u.Node] == Dist[v.Node]+u.Dist {
				Count[u.Node] += Count[v.Node]
			}
		}
	}

	log.Print("-> Dist: ", Dist)
	log.Print("-> Prev: ", Prev)
	log.Print("-> Count: ", Count)

	return Count[n-1] % M
}

// 2115m Find All Possible Recipes from Given Supplies
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	S := map[string]struct{}{}
	for _, s := range supplies {
		S[s] = struct{}{}
	}

	M := map[string]int{}
	for i, r := range recipes {
		M[r] = i
	}

	G := make([][]string, len(recipes))
	D := make([]int, len(recipes))

	for i, ings := range ingredients {
		G[i] = append(G[i], ings...)
		for _, ing := range ings {
			if _, ok := S[ing]; !ok {
				D[i]++
			}
		}
	}

	Q := []int{}
	for i, d := range D {
		if d == 0 {
			Q = append(Q, i)
		}
	}

	R := []string{}
	var u int
	for len(Q) > 0 {
		u, Q = Q[0], Q[1:]
		R = append(R, recipes[u])

		for _, ing := range G[u] {
			D[M[ing]]--
			if D[M[ing]] == 0 {
				Q = append(Q, M[ing])
			}
		}
	}

	return R
}

// 2127h Maximum Employees to Be Invited to a Meeting
func maximumInvitations(favorite []int) int {
	N := len(favorite)

	Din := make([]int, N) // InDegree
	for _, f := range favorite {
		Din[f]++
	}

	log.Print(" -> ", Din)

	dq := list.List{}
	for p, fcount := range Din {
		if fcount == 0 {
			dq.PushBack(p)
		}
	}

	Depth := make([]int, N)
	for p := range Depth {
		Depth[p] = 1
	}

	for dq.Len() > 0 {
		p := dq.Remove(dq.Front()).(int)
		f := favorite[p]

		Depth[f] = max(Depth[p]+1, Depth[f])
		Din[f]--
		if Din[f] == 0 {
			dq.PushBack(f)
		}
	}

	log.Print(" -> ", Depth)

	longestCycle, twoCycles := 0, 0
	for p := range N {
		if Din[p] == 0 {
			continue
		}

		lCycle, cur := 0, p
		for Din[cur] != 0 {
			Din[cur] = 0
			lCycle++
			cur = favorite[cur]
		}

		switch lCycle {
		case 2:
			twoCycles += Depth[p] + Depth[favorite[p]]
		default:
			longestCycle = max(longestCycle, lCycle)
		}
	}

	return max(longestCycle, twoCycles)
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

// 2467m Most Profitable Path in a Tree
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	N := len(edges) + 1

	G := make([][]int, N)
	for _, e := range edges {
		u, v := e[0], e[1]
		G[u], G[v] = append(G[u], v), append(G[v], u)
	}

	BDist := make([]int, N) // nodes distance to Bob
	for i := range BDist {
		BDist[i] = N
	}
	BDist[bob] = 0

	var Search func(u, p, t int) int
	Search = func(u, p, t int) int {
		xCur, xNext := 0, math.MinInt

		for _, v := range G[u] {
			if v != p {
				xNext = max(Search(v, u, t+1), xNext)
				BDist[u] = min(BDist[v]+1, BDist[u])
			}
		}

		if BDist[u] > t {
			xCur += amount[u]
		} else if BDist[u] == t {
			xCur += amount[u] / 2
		}

		if xNext != math.MinInt {
			return xCur + xNext
		}
		return xCur
	}

	return Search(0, 0, 0)
}

// 2503h Maximum Number of Points From Grid Queries
type PQ2503 [][3]int

func (h PQ2503) Len() int           { return len(h) }
func (h PQ2503) Less(i, j int) bool { return h[i][0] <= h[j][0] }
func (h PQ2503) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ2503) Push(x any) { *h = append(*h, x.([3]int)) }
func (h *PQ2503) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func maxPoints_GridQueries(grid [][]int, queries []int) []int {
	Rows, Cols := len(grid), len(grid[0])
	R := make([]int, len(queries))

	sQry := [][2]int{}
	for i, qry := range queries {
		sQry = append(sQry, [2]int{qry, i})
	}
	slices.SortFunc(sQry, func(x, y [2]int) int { return x[0] - y[0] })

	log.Print("-> ", sQry)

	Vis := make([][]bool, Rows)
	for r := range Rows {
		Vis[r] = make([]bool, Cols)
	}

	pq := PQ2503{}
	Vis[0][0] = true
	heap.Push(&pq, [3]int{grid[0][0], 0, 0})

	Dir := []int{1, 0, -1, 0, 1}
	points := 0
	for _, qry := range sQry {
		for pq.Len() > 0 && pq[0][0] < qry[0] {
			points++

			log.Print("-> ", qry, points, pq)

			e := heap.Pop(&pq).([3]int)
			r, c := e[1], e[2]

			for d := range 4 {
				r, c := r+Dir[d], c+Dir[d+1]
				if r >= 0 && Rows > r && c >= 0 && Cols > c && !Vis[r][c] {
					Vis[r][c] = true
					heap.Push(&pq, [3]int{grid[r][c], r, c})
				}
			}
		}

		R[qry[1]] = points
	}

	return R
}

// 2608m Shortest Cycle in a Graph
func findShortestCycle(n int, edges [][]int) int {
	G := make([][]int, n)
	for _, e := range edges {
		G[e[0]] = append(G[e[0]], e[1])
		G[e[1]] = append(G[e[1]], e[0])
	}

	log.Print(" -> Graph :: ", G)

	mCycle := math.MaxInt
	for src := range G {
		Dist := make([]int, n)
		for i := range Dist {
			Dist[i] = math.MaxInt
		}

		Q := list.List{}
		Q.PushBack(src)
		Dist[src] = 0

		for Q.Len() > 0 {
			v := Q.Remove(Q.Front()).(int)
			for _, u := range G[v] {
				switch Dist[u] {
				case math.MaxInt:
					Dist[u] = Dist[v] + 1
					Q.PushBack(u)
				default:
					if Dist[v] <= Dist[u] {
						mCycle = min(Dist[u]+Dist[v]+1, mCycle)
					}
				}
			}
		}
	}

	if mCycle == math.MaxInt {
		return -1
	}
	return mCycle
}

// 2658m Maximum Number of Fish in a Grid
func findMaxFish(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])

	var Catch func(r, c int) int
	Catch = func(r, c int) int {
		if r < 0 || c < 0 || Rows <= r || Cols <= c || grid[r][c] == 0 {
			return 0
		}

		fish := grid[r][c]
		grid[r][c] = 0
		return fish + Catch(r-1, c) + Catch(r+1, c) + Catch(r, c-1) + Catch(r, c+1)
	}

	xFish := 0
	for r := range Rows {
		for c := range Cols {
			xFish = max(Catch(r, c), xFish)
		}
	}
	return xFish
}

// 2685m Count the Number of Complete Components
func countCompleteComponents(n int, edges [][]int) int {
	cliques := 0

	Vis := make([]bool, n)
	G := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		G[v], G[u] = append(G[v], u), append(G[u], v)
	}

	log.Print("-> Graph: ", G)

	var DFS func(int) (int, int)
	DFS = func(v int) (vertices int, edges int) {
		vertices++

		for _, u := range G[v] {
			edges++
			if !Vis[u] {
				Vis[u] = true

				vs, es := DFS(u)

				vertices += vs
				edges += es
			}
		}

		return vertices, edges
	}

	components := 0
	for v := range n {
		if !Vis[v] {
			Vis[v] = true

			components++

			vertices, edges := DFS(v)
			if vertices*(vertices-1) == edges {
				cliques++
			}
		}
	}

	log.Print("-> Components: ", components)

	return cliques
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

// 3108h Minimum Cost Walk in Weighted Graph
func minimumCostWalk(n int, edges [][]int, query [][]int) []int {
	R := []int{}
	return R
}

// 3203h Find Minimum Distance After Merging Two Trees
func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	lAdj := func(edges [][]int) [][]int {
		G := make([][]int, len(edges)+1)
		for _, e := range edges {
			G[e[0]] = append(G[e[0]], e[1])
			G[e[1]] = append(G[e[1]], e[0])
		}
		return G
	}

	fDiameter := func(G [][]int) int {
		nSearch := func(n int) (int, int) {
			Q := []int{n}
			Vis := make([]bool, len(G))
			Vis[n] = true

			xdis := -1
			for len(Q) > 0 {
				for range len(Q) {
					n, Q = Q[0], Q[1:]
					for _, u := range G[n] {
						if !Vis[u] {
							Vis[u] = true
							Q = append(Q, u)
						}
					}
				}
				xdis++
			}

			return n, xdis
		}

		fNode, _ := nSearch(0)
		_, diameter := nSearch(fNode)
		return diameter
	}

	d1, d2 := fDiameter(lAdj(edges1)), fDiameter(lAdj(edges2))
	return slices.Max([]int{d1, d2, (d1+1)/2 + (d2+1)/2 + 1})
}

// 3286m Find a Safe Walk Through a Grid
func findSafeWalk(grid [][]int, health int) bool {
	Rows, Cols := len(grid), len(grid[0])

	health -= grid[0][0]
	if health == 0 {
		return false
	}

	Vis := make([][]int, Rows)
	for r := range Vis {
		Vis[r] = make([]int, Cols)
	}

	Q := [][]int{{0, 0, health}}
	Vis[0][0] = health

	Dir := []int{0, 1, 0, -1, 0}

	var v []int
	for len(Q) > 0 {
		v, Q = Q[0], Q[1:]
		r, c, health := v[0], v[1], v[2]
		if r == Rows-1 && c == Cols-1 {
			return true
		}

		for dir := range 4 {
			r, c := r+Dir[dir], c+Dir[dir+1]
			if 0 <= r && r < Rows && 0 <= c && c < Cols {
				health := health - grid[r][c]
				if health > Vis[r][c] {
					Vis[r][c] = health
					Q = append(Q, []int{r, c, health})
				}
			}
		}
	}

	return false
}

// 3341m Find Minimum Time To Reach Last Room I
type PQ3341 [][3]int

func (h PQ3341) Len() int           { return len(h) }
func (h PQ3341) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h PQ3341) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ3341) Push(x any) { *h = append(*h, x.([3]int)) }
func (h *PQ3341) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func minTimeToReach(moveTime [][]int) int {
	Rows, Cols := len(moveTime), len(moveTime[0])

	Grid := make([][]int, Rows) // Distances from {0,0}
	for r := range Grid {
		Grid[r] = make([]int, Cols)
	}

	for r := range Grid {
		for c := range Grid[r] {
			Grid[r][c] = math.MaxInt
		}
	}

	pq := PQ3341{}
	heap.Push(&pq, [3]int{0, 0, 0})
	Grid[0][0] = 0

	Dirs := []int{-1, 0, 1, 0, -1}
	for pq.Len() > 0 {
		v := heap.Pop(&pq).([3]int)
		time, r, c := v[0], v[1], v[2]

		log.Print("-> ", v, pq, Grid)

		if r == Rows-1 && c == Cols-1 {
			return time
		}

		for d := range 4 {
			r, c := r+Dirs[d], c+Dirs[d+1]
			if 0 <= r && r < Rows && 0 <= c && c < Cols && max(time, moveTime[r][c]) < Grid[r][c] {
				heap.Push(&pq, [3]int{max(time, moveTime[r][c]) + 1, r, c})
				Grid[r][c] = max(time, moveTime[r][c])
			}
		}
	}

	return -1
}

// 3342m Find Minimum Time To Reach Last Room II
type PQ3342 [][4]int

func (h PQ3342) Len() int           { return len(h) }
func (h PQ3342) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h PQ3342) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ3342) Push(x any) { *h = append(*h, x.([4]int)) }
func (h *PQ3342) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func minTimeToReachII(moveTime [][]int) int {
	Rows, Cols := len(moveTime), len(moveTime[0])
	Grid := make([][]int, Rows)
	for r := range Grid {
		Grid[r] = make([]int, Cols)
	}

	for r := range Grid {
		for c := range Grid[r] {
			Grid[r][c] = math.MaxInt
		}
	}

	pq := PQ3342{}
	heap.Push(&pq, [4]int{0, 0, 0, 0})
	Grid[0][0] = 0

	Dirs := []int{-1, 0, 1, 0, -1}

	for pq.Len() > 0 {
		v := heap.Pop(&pq).([4]int)
		time, r, c, double := v[0], v[1], v[2], v[3]

		log.Print("-> ", time, r, c, double, pq, Grid)

		if r == Rows-1 && c == Cols-1 {
			return time
		}

		delta := 1
		if double == 1 {
			delta++
		}

		for d := range 4 {
			r, c := r+Dirs[d], c+Dirs[d+1]
			if 0 <= r && r < Rows && 0 <= c && c < Cols && max(time, moveTime[r][c])+delta < Grid[r][c] {
				Grid[r][c] = max(time, moveTime[r][c]) + delta
				heap.Push(&pq, [4]int{Grid[r][c], r, c, 1 ^ double})
			}
		}
	}

	return -1
}

// 3372m Maximize the Number of Target Nodes After Connecting Trees I
func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	MkTree := func(edges [][]int) [][]int {
		T := make([][]int, len(edges)+1)
		for _, edge := range edges {
			v, u := edge[0], edge[1]
			T[v], T[u] = append(T[v], u), append(T[u], v)
		}
		return T
	}

	T1, T2 := MkTree(edges1), MkTree(edges2)

	log.Print("-> ", T1)
	log.Print("-> ", T2)

	var Search func(v, p, k int, Tree [][]int) int
	Search = func(v, p, k int, Tree [][]int) int {
		if k < 0 {
			return 0
		}

		count := 1
		for _, u := range Tree[v] {
			if u != p {
				count += Search(u, v, k-1, Tree)
			}
		}
		return count
	}

	K := []int{k - 1, k}
	var Trg []int

	xTrg := 0
	for _, T := range [][][]int{T2, T1} {
		k, K = K[0], K[1:]

		Trg = []int{}
		for src := range T {
			count := Search(src, math.MaxInt, k, T)
			Trg = append(Trg, count+xTrg)
		}
		xTrg = slices.Max(Trg)

		log.Printf("-> k: %d => %v", k, Trg)
	}

	return Trg
}

// 3373h Maximize the Number of Target Nodes After Connecting Trees II
func maxTargetNodesII(edges1 [][]int, edges2 [][]int) []int {
	MkTree := func(edges [][]int) [][]int {
		T := make([][]int, len(edges)+1)
		for _, edge := range edges {
			v, u := edge[0], edge[1]
			T[v], T[u] = append(T[v], u), append(T[u], v)
		}
		return T
	}

	T1, T2 := MkTree(edges1), MkTree(edges2)

	log.Print("-> ", T1)
	log.Print("-> ", T2)

	Search := func(src, parity int, Tree [][]int) int {
		Q := list.New()
		Q.PushBack([3]int{src, math.MaxInt, parity})

		count := 0

		for Q.Len() > 0 {
			q := Q.Remove(Q.Front()).([3]int)
			v, p, parity := q[0], q[1], q[2]

			if parity&1 == 0 {
				count++
			}

			for _, u := range Tree[v] {
				if u != p {
					Q.PushBack([3]int{u, v, parity ^ 1})
				}
			}
		}

		return count
	}

	Evens := []int{}
	for src := range len(T1) {
		Evens = append(Evens, Search(src, 0, T1))
	}
	log.Print("-> ", Evens)

	Odds := []int{}
	for src := range len(T2) {
		Odds = append(Odds, Search(src, 1, T2))
	}
	log.Print("-> ", Odds)
	xOdd := slices.Max(Odds)

	R := []int{}
	for _, even := range Evens {
		R = append(R, even+xOdd)
	}

	return R
}
