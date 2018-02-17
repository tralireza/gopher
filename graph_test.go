package gopher

import (
	"container/heap"
	"log"
	"math"
	"slices"
	"testing"
)

// 126h Word Ladder II
func Test126(t *testing.T) {
	log.Printf(`["hit" "hot" "dot" "dog" "cog"] ["hit" "hot" "lot" "log" "cog"] ?= %q`, findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	log.Print("[] ?= ", findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))

	log.Print("--")
	log.Print(" ->", findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))

	R := findLadders("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"})
	log.Print(" -> ", R)
	log.Print(" -> ", len(R))
}

// 127h Word Ladder
func Test127(t *testing.T) {
	log.Print("5 ?= ", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	log.Print("0 ?= ", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

// 210m Course Schedule II
func Test210(t *testing.T) {
	// Prerequisite,  {a_i, b_i}  ::  b_i -> a_i

	log.Print("[0 1] ?= ", findOrder(2, [][]int{{1, 0}}))
	log.Print("[0 2 1 3] ?= ", findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	log.Print("[0] ?= ", findOrder(1, [][]int{}))
}

// 433m Minimum Genetic Mutation
func Test443(t *testing.T) {
	log.Print("1 ?= ", minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	log.Print("2 ?= ", minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}

// 595m Regions Cut By Slashes
func Test595(t *testing.T) {
	WithDJS := func(grid []string) int {
		//           .-.-.  (/: bottom-left: v "Connected-To" top-right: u)
		// '/''\' -> |/|\|  (\: bottom-right: v "Connected-To" top-left: u)
		//           .-.-.
		//
		// 2x2 grid -> 3x3 nodes Graph

		type P struct{ x, y int }
		DJS := map[P]P{}     // Node{x,y} -> leader/root
		Ranks := map[P]int{} // DJS ranks

		var FindSet func(a P) P
		FindSet = func(a P) P {
			r := DJS[a]
			if r != a {
				DJS[a] = FindSet(r)
			}
			return DJS[a]
		}

		Union := func(a, b P) {
			a, b = FindSet(a), FindSet(b)
			if a == b {
				return
			}
			ra, rb := Ranks[a], Ranks[b]
			if ra >= rb {
				DJS[b] = a
				if ra == rb {
					Ranks[a]++
				}
			} else {
				DJS[a] = b
			}
		}

		Rows, Cols := len(grid), len(grid[0])

		// init DJS
		for r := 0; r <= Rows; r++ {
			for c := 0; c <= Cols; c++ {
				DJS[P{r, c}] = P{r, c}
			}
		}

		// Union/Join boundary points/vertices
		Bdr := P{0, 0}
		for r := 0; r <= Rows; r++ {
			Union(P{r, 0}, Bdr)
			Union(P{r, Cols}, Bdr)
		}
		for c := 1; c < Cols; c++ {
			Union(P{0, c}, Bdr)
			Union(P{Rows, c}, Bdr)
		}
		regions := 1 // all boundary points (connected) form a loop/region

		for r := 0; r < Rows; r++ {
			for c := 0; c < Cols; c++ {
				switch grid[r][c] {
				case '/':
					a, b := P{r + 1, c}, P{r, c + 1}
					if FindSet(a) == FindSet(b) {
						regions++
					}
					Union(a, b) // Union bottom-left and top-right
				case '\\':
					a, b := P{r, c}, P{r + 1, c + 1}
					if FindSet(a) == FindSet(b) {
						regions++
					}
					Union(a, b) // Union top-left and bottom-right
				}
			}
		}

		log.Print("DJS: ", DJS)
		log.Print("Ranks: ", Ranks)

		return regions
	}

	for _, f := range []func([]string) int{regionsBySlashes, WithDJS} {
		log.Print("2 ?= ", f([]string{" /", "/ "}))
		log.Print("1 ?= ", f([]string{"  ", "/ "}))
		log.Print("5 ?= ", f([]string{"/\\", "\\/"}))

		log.Print("14 ?= ", f([]string{
			"//\\\\////",
			"//\\\\/\\//",
			"\\/ /\\\\/\\",
			"///\\\\\\\\ ",
			"//  / \\\\",
			"\\/\\/ //\\",
			" // \\ \\\\",
			"/\\\\/\\\\\\/",
		}))
		log.Print("---")
	}
}

// 733 Flood Fill
func Test733(t *testing.T) {
	log.Print(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	log.Print(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0))
}

// 909m Snakes & Ladders
func Test909(t *testing.T) {
	// 'Boustrophedon' style

	log.Print("4 ?= ", snakesAndLadders([][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}))
	log.Print("1 ?= ", snakesAndLadders([][]int{{-1, -1}, {-1, 3}}))

	log.Print("4 ?= ", snakesAndLadders([][]int{{-1, -1, -1, 46, 47, -1, -1, -1}, {51, -1, -1, 63, -1, 31, 21, -1}, {-1, -1, 26, -1, -1, 38, -1, -1}, {-1, -1, 11, -1, 14, 23, 56, 57}, {11, -1, -1, -1, 49, 36, -1, 48}, {-1, -1, -1, 33, 56, -1, 57, 21}, {-1, -1, -1, -1, -1, -1, 2, -1}, {-1, -1, -1, 8, 3, -1, 6, 56}}))
}

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

// 1192h Critical Connections in a Network
func Test1192(t *testing.T) {
	Tarjan := func(n int, edges [][]int) [][]int {
		R := [][]int{}

		G := make([][]int, n)
		for _, e := range edges {
			v, u := e[0], e[1]
			G[v] = append(G[v], u)
			G[u] = append(G[u], v)
		}

		log.Print("Graph :: ", G)

		Index := make([]int, n)
		Lowest := make([]int, n)
		OnStack := make([]bool, n)

		Q := []int{}
		index := 1 // discovery time

		var Tarjan func(int)
		Tarjan = func(v int) {
			Index[v], Lowest[v], OnStack[v] = index, index, true
			index++
			Q = append(Q, v)

			for _, u := range G[v] {
				if Index[u] == 0 {
					Tarjan(u)
					Lowest[v] = min(Lowest[v], Lowest[u])
				} else if OnStack[u] {
					Lowest[v] = min(Lowest[v], Index[u])
				}
			}

			if Index[v] == Lowest[v] {
				Scc := []int{} // Strongly-Connected-Components
				var u int
				for {
					u, Q = Q[len(Q)-1], Q[:len(Q)-1]
					OnStack[u] = false
					Scc = append(Scc, u)
					if u == v {
						break
					}
				}
				R = append(R, Scc)
			}
		}

		for v := range n { // Nodes [0...n-1]
			if Index[v] == 0 {
				Tarjan(v)
			}
		}

		return R
	}

	log.Print("Tarjan :: ", Tarjan(4, [][]int{{0, 1}, {2, 3}}))

	log.Print("[[1 3]] ?= ", criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
	log.Print("[[0 1]] ?= ", criticalConnections(2, [][]int{{0, 1}}))
}

// 1591h Strange Printer II
func Test1591(t *testing.T) {
	log.Print("true ?= ", isPrintable([][]int{{1, 1, 1, 1}, {1, 2, 2, 1}, {1, 2, 2, 1}, {1, 1, 1, 1}}))
	log.Print("true ?= ", isPrintable([][]int{{1, 1, 1, 1}, {1, 1, 3, 3}, {1, 1, 3, 4}, {5, 5, 1, 4}}))
	log.Print("false ?= ", isPrintable([][]int{{1, 2, 1}, {2, 1, 2}, {1, 2, 1}}))
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
		Color[0] = -1             // ignore (no) Node at label 0

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
