package gopher

import (
	"container/heap"
	"container/list"
	"log"
	"math"
	"reflect"
	"slices"
	"testing"
)

// 126h Word Ladder II
func Test126(t *testing.T) {
	log.Printf(`["hit" "hot" "dot" "dog" "cog"] ["hit" "hot" "lot" "log" "cog"] ?= %q`, findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	log.Print("[] ?= ", findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	log.Print(" -> ", findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))

	/*
		log.Print("--")
		R := findLadders("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"})
		log.Print(" -> ", len(R))
	*/
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

// 695m Max Area of Island
func Test695(t *testing.T) {
	BFS := func(grid [][]int) int {
		xArea := 0
		for r := range grid {
			for c := range grid[r] {
				if grid[r][c] != 0 {
					q := list.List{}
					q.PushBack([2]int{r, c})

					area := 0
					for q.Len() > 0 {
						coord := q.Remove(q.Front()).([2]int)
						r, c := coord[0], coord[1]
						grid[r][c] = 0
						area++

						Dir := []int{-1, 0, 1, 0, -1}
						for d := range 4 {
							r, c := r+Dir[d], c+Dir[d+1]
							if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[r]) && grid[r][c] != 0 {
								q.PushBack([2]int{r, c})
							}
						}
					}

					xArea = max(area, xArea)
				}
			}
		}

		return xArea
	}

	for _, f := range []func([][]int) int{maxAreaOfIsland, BFS} {
		log.Print("6 ?= ", f([][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
		log.Print("0 ?= ", f([][]int{{0, 0, 0, 0, 0, 0, 0, 0}}))
		log.Print("--")
	}
}

// 733 Flood Fill
func Test733(t *testing.T) {
	log.Print(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	log.Print(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0))
}

func Test753(t *testing.T) {
	for _, c := range []struct {
		rst  string
		n, k int
	}{
		{"10", 1, 2},
		{"01100", 2, 2},

		{"0011101000", 3, 2},
		{"00222121112201202101102001000", 3, 3},
	} {
		log.Print("* ", c.n, c.k)
		if c.rst != crackSafe(c.n, c.k) {
			t.FailNow()
		}
	}
}

// 802m Find Eventual Safe States
func Test802(t *testing.T) {
	Kahn := func(graph [][]int) []int {
		N := len(graph)

		In := make([]int, N)

		G := make([][]int, N)
		for v := range N {
			for _, u := range graph[v] {
				G[u] = append(G[u], v)
				In[v]++
			}
		}

		log.Print(" ", graph, " -> ", G)
		log.Print(" -> ", In)

		S := make([]bool, N)

		Q := []int{}
		for v := range N {
			if In[v] == 0 {
				Q = append(Q, v)
			}
		}

		var v int
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			S[v] = true

			for _, u := range G[v] {
				In[u]--
				if In[u] == 0 {
					Q = append(Q, u)
				}
			}
		}

		R := []int{}
		for v := range N {
			if S[v] {
				R = append(R, v)
			}
		}
		return R
	}

	for _, f := range []func([][]int) []int{eventualSafeNodes, Kahn} {
		log.Print("[2 4 5 6] ?= ", f([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
		log.Print("[4] ?= ", f([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
		log.Print("--")
	}
}

// 827h Making a Large Island
func Test827(t *testing.T) {
	log.Print("3 ?= ", largestIsland([][]int{{1, 0}, {0, 1}}))
	log.Print("4 ?= ", largestIsland([][]int{{1, 1}, {1, 0}}))
	log.Print("4 ?= ", largestIsland([][]int{{1, 1}, {1, 1}}))
}

// 909m Snakes & Ladders
func Test909(t *testing.T) {
	// 'Boustrophedon' style

	for _, c := range []struct {
		rst   int
		board [][]int
	}{
		{4, [][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}},
		{1, [][]int{{-1, -1}, {-1, 3}}},

		{4, [][]int{
			{-1, -1, -1, 46, 47, -1, -1, -1},
			{51, -1, -1, 63, -1, 31, 21, -1},
			{-1, -1, 26, -1, -1, 38, -1, -1},
			{-1, -1, 11, -1, 14, 23, 56, 57},
			{11, -1, -1, -1, 49, 36, -1, 48},
			{-1, -1, -1, 33, 56, -1, 57, 21},
			{-1, -1, -1, -1, -1, -1, 2, -1},
			{-1, -1, -1, 8, 3, -1, 6, 56}},
		},
	} {
		if c.rst != snakesAndLadders(c.board) {
			t.FailNow()
		}
	}
}

// 947m Most Stones Removed with Same Row or Column
func Test947(t *testing.T) {
	log.Print("5 ?= ", removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
	log.Print("3 ?= ", removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	log.Print("0 ?= ", removeStones([][]int{{0, 0}}))
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

// 1267m Count Servers that Communicate
func Test1267(t *testing.T) {
	log.Print("0 ?= ", countServers([][]int{{1, 0}, {0, 1}}))
	log.Print("3 ?= ", countServers([][]int{{1, 0}, {1, 1}}))
	log.Print("4 ?= ", countServers([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
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

// 1368h Minimum Cost to Make at Least One Valid Path in a Grid
func Test1368(t *testing.T) {
	log.Print("3 ?= ", minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}))
	log.Print("0 ?= ", minCost([][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}))
	log.Print("1 ?= ", minCost([][]int{{1, 2}, {4, 3}}))
}

// 1462m Course Schedule IV
func Test1462(t *testing.T) {
	log.Print("[false true] ?= ", checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}))
	log.Print("[false false] ?= ", checkIfPrerequisite(2, [][]int{}, [][]int{{0, 1}, {1, 0}}))
	log.Print("[true true] ?= ", checkIfPrerequisite(3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}))
}

// 1514m Path with Maximum Probability
func Test1514(t *testing.T) {
	BellmanFord := func(n int, edges [][]int, succProb []float64, start_node, end_node int) float64 {
		Dist := make([]float64, n)

		Dist[start_node] = 1
		for range n - 1 {
			for i, e := range edges { // relaxing all edges...
				v, u := e[0], e[1]
				if Dist[v]*succProb[i] > Dist[u] {
					Dist[u] = Dist[v] * succProb[i]
				}
				if Dist[u]*succProb[i] > Dist[v] {
					Dist[v] = Dist[u] * succProb[i]
				}
			}
		}
		return Dist[end_node]
	}

	SPF := func(n int, edges [][]int, succProb []float64, start_node, end_node int) float64 {
		// Shortest-Path-First

		type E struct {
			n int     // Node
			w float64 // Weight ie, Probability
		}

		G := make([][]E, n)
		for i, e := range edges {
			u, v, w := e[0], e[1], succProb[i]
			G[v], G[u] = append(G[v], E{u, w}), append(G[u], E{v, w})
		}

		Dist := make([]float64, n)
		Dist[start_node] = 1

		Q := []int{start_node}
		var v int
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			for _, u := range G[v] {
				if Dist[v]*u.w > Dist[u.n] { // relaxing neighbor edges...
					Dist[u.n] = Dist[v] * u.w
					Q = append(Q, u.n)
				}
			}
		}

		return Dist[end_node]
	}

	for _, f := range []func(int, [][]int, []float64, int, int) float64{maxProbability, BellmanFord, SPF} {
		log.Print("0.25 ?= ", f(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{.5, .5, .2}, 0, 2))
		log.Print("0.3 ?= ", f(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{.5, .5, .3}, 0, 2))
		log.Print("0 ?= ", f(3, [][]int{{0, 1}}, []float64{.5}, 0, 2))
		log.Print("--")
	}
}

// 1591h Strange Printer II
func Test1591(t *testing.T) {
	log.Print("true ?= ", isPrintable([][]int{{1, 1, 1, 1}, {1, 2, 2, 1}, {1, 2, 2, 1}, {1, 1, 1, 1}}))
	log.Print("true ?= ", isPrintable([][]int{{1, 1, 1, 1}, {1, 1, 3, 3}, {1, 1, 3, 4}, {5, 5, 1, 4}}))
	log.Print("false ?= ", isPrintable([][]int{{1, 2, 1}, {2, 1, 2}, {1, 2, 1}}))
}

// 1765m Map of Highest Peak
func Test1767(t *testing.T) {
	log.Print("[[1 0] [2 1]] ?= ", highestPeak([][]int{{0, 1}, {0, 0}}))
	log.Print("[[1 1 0] [0 1 1] [1 2 2]] ?= ", highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}))
}

// 1905m Count Sub Islands
func Test1905(t *testing.T) {
	log.Print("3 ?= ", countSubIslands([][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}, [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}))
	log.Print("2 ?= ", countSubIslands([][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}, [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}))
}

// 1976m Number of Ways to Arrive at Destination
func Test1976(t *testing.T) {
	type E = E1976
	type PQ = PQ1976

	Dijkstra := func(n int, roads [][]int) int {
		const M = 1000_000_007

		G := make([][]E, n)
		for _, e := range roads {
			v, u, w := e[0], e[1], e[2]
			G[v], G[u] = append(G[v], E{Node: u, Dist: w}), append(G[u], E{Node: v, Dist: w})
		}

		Dist := make([]int, n)

		for v := range n {
			Dist[v] = math.MaxInt
		}

		Prev := make([][]int, n)

		Count := make([]int, n)
		Count[0] = 1

		h := PQ{}

		heap.Push(&h, E{Node: 0, Dist: 0})
		Dist[0] = 0

		for h.Len() > 0 {
			log.Print("-> PQ: ", h)
			v := heap.Pop(&h).(E)

			if Dist[v.Node] != v.Dist {
				log.Print(" + already visited :: ", v, " distance from start(v): ", Dist[v.Node])
				Prev[v.Node] = Prev[v.Node][1:]
				continue
			}

			for _, u := range G[v.Node] {
				if Dist[v.Node]+u.Dist < Dist[u.Node] {
					Dist[u.Node] = Dist[v.Node] + u.Dist
					heap.Push(&h, E{Node: u.Node, Dist: Dist[u.Node]})

					Count[u.Node] = Count[v.Node]
					Prev[u.Node] = append(Prev[u.Node], v.Node)

				} else if Dist[v.Node]+u.Dist == Dist[u.Node] {
					Count[u.Node] += Count[v.Node]
					Count[u.Node] %= M
					Prev[u.Node] = append(Prev[u.Node], v.Node)
				}
			}
		}

		log.Print("-> Dist: ", Dist)
		log.Print("-> Prev: ", Prev)

		return Count[n-1]
	}

	for _, c := range []struct {
		rst, n int
		roads  [][]int
	}{
		{4, 7, [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}},
		{1, 2, [][]int{{1, 0, 10}}},
		{2, 5, [][]int{{3, 0, 5}, {0, 1, 1}, {1, 2, 4}, {0, 4, 3}, {3, 2, 5}, {3, 4, 1}, {1, 3, 1}}},
	} {
		rst, n, roads := c.rst, c.n, c.roads
		for _, f := range []func(int, [][]int) int{countPaths, Dijkstra} {
			if rst != f(n, roads) {
				t.FailNow()
			}
		}
		log.Printf(":: %d <- %v", rst, roads)
	}
}

func Test2115(t *testing.T) {
	for _, c := range []struct {
		rst         []string
		recipes     []string
		ingredients [][]string
		supplies    []string
	}{
		{[]string{"bread"}, []string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}},
	} {
		rst, recipes, ingredients, supplies := c.rst, c.recipes, c.ingredients, c.supplies
		if !reflect.DeepEqual(rst, findAllRecipes(recipes, ingredients, supplies)) {
			t.FailNow()
		}
		log.Printf(":: %v <-", rst)
	}
}

// 2127h Maximum Employees to Be Invited to a Meeting
func Test2127(t *testing.T) {
	log.Print("3 ?= ", maximumInvitations([]int{2, 2, 1, 2}))
	log.Print("3 ?= ", maximumInvitations([]int{1, 2, 0}))
	log.Print("4 ?= ", maximumInvitations([]int{3, 0, 1, 4, 1}))
}

func Test2359(t *testing.T) {
	for _, c := range []struct {
		rst          int
		edges        []int
		node1, node2 int
	}{
		{2, []int{2, 2, 3, -1}, 0, 1},
		{2, []int{1, 2, -1}, 0, 2},

		{0, []int{1, 0}, 1, 0},
		{4, []int{4, 3, 0, 5, 3, -1}, 4, 0},
	} {
		log.Print("* ", c.edges)
		if c.rst != closestMeetingNode(c.edges, c.node1, c.node2) {
			t.FailNow()
		}
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

// 2467m Most Profitable Path in a Tree
func Test2467(t *testing.T) {
	log.Print("6 ?= ", mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
	log.Print("-7280 ?= ", mostProfitablePath([][]int{{0, 1}}, 1, []int{-7280, 2350}))

	log.Print("-3360 ?= ", mostProfitablePath([][]int{{0, 1}, {0, 2}}, 2, []int{-3360, -5394, -1146}))
}

func Test2503(t *testing.T) {
	for _, c := range []struct {
		rst     []int
		grid    [][]int
		queries []int
	}{
		{[]int{5, 8, 1}, [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, []int{5, 6, 2}},
		{[]int{0}, [][]int{{5, 2, 1}, {1, 1, 2}}, []int{3}},
	} {
		if !reflect.DeepEqual(c.rst, maxPoints_GridQueries(c.grid, c.queries)) {
			t.FailNow()
		}
		log.Printf(":: %v   <-   G: %v   Q: %v", c.rst, c.grid, c.queries)
	}
}

// 2608m Shortest Cycle in a Graph
func Test2608(t *testing.T) {
	for _, f := range []func(int, [][]int) int{findShortestCycle} {
		log.Print("3 ?= ", f(7, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}))
		log.Print("-1 ?= ", f(4, [][]int{{0, 1}, {0, 2}}))

		log.Print("4 ?= ", f(6, [][]int{{4, 1}, {3, 2}, {5, 0}, {3, 0}, {4, 0}, {2, 1}, {5, 1}}))
		log.Print("--")
	}
}

// 2658m Maximum Number of Fish in a Grid
func Test2658(t *testing.T) {
	BFS := func(grid [][]int) int {
		xFish := 0

		for r := range grid {
			for c := range grid[r] {
				if grid[r][c] != 0 {
					fish := 0
					q := list.List{}

					q.PushBack([2]int{r, c})
					for q.Len() > 0 {
						coord := q.Remove(q.Front()).([2]int)
						r, c := coord[0], coord[1]

						fish += grid[r][c]
						grid[r][c] = 0

						Dir := []int{-1, 0, 1, 0, -1}
						for i := range 4 {
							r, c := r+Dir[i], c+Dir[i+1]
							if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[r]) && grid[r][c] != 0 {
								q.PushBack([2]int{r, c})
							}
						}
					}

					xFish = max(fish, xFish)
				}
			}
		}

		return xFish
	}

	for _, f := range []func([][]int) int{findMaxFish, BFS} {
		log.Print("7 ?= ", f([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}))
		log.Print("1 ?= ", f([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}))

		log.Print("24 ?= ", f([][]int{{4, 5, 5}, {0, 10, 0}}))
		log.Print("--")
	}
}

// 2685m Count the Number of Complete Components
func Test2685(t *testing.T) {
	BFS := func(n int, edges [][]int) int {
		G := make([][]int, n)
		for _, e := range edges {
			G[e[0]], G[e[1]] = append(G[e[0]], e[1]), append(G[e[1]], e[0])
		}

		Vis := make([]bool, n)

		cliques := 0
		Components := [][]int{}

		for v := range n {
			if !Vis[v] {
				Q := []int{v}
				Vis[v] = true

				cmp := []int{}
				for len(Q) > 0 {
					v, Q = Q[0], Q[1:]
					cmp = append(cmp, v)

					for _, u := range G[v] {
						if !Vis[u] {
							Vis[u] = true
							Q = append(Q, u)
						}
					}
				}

				Components = append(Components, cmp)
			}
		}

		log.Print("-> Components: ", Components)
		for _, cmp := range Components {
			edges := 0
			for _, v := range cmp {
				edges += len(G[v])
			}

			vertices := len(cmp)
			if vertices*(vertices-1) == edges {
				cliques++
			}
		}

		return cliques
	}

	UnionFind := func(n int, edges [][]int) int {
		djset := make([][3]int, n)
		for i := range djset {
			djset[i][0], djset[i][1], djset[i][2] = i, 0, 1 // 0: Leader, 1: #edges, 2: #vertices
		}

		var Find func(int) int
		Find = func(x int) int {
			if djset[x][0] != x {
				djset[x][0] = Find(djset[x][0])
			}
			return djset[x][0]
		}

		Union := func(x, y int) {
			x, y = Find(x), Find(y)
			if x != y {
				if djset[x][1] < djset[y][1] {
					x, y = y, x
				}
				djset[y][0] = x

				djset[x][1] += djset[y][1]
				djset[x][2] += djset[y][2]
			}
			djset[x][1]++
		}

		for _, e := range edges {
			Union(e[0], e[1])
		}

		log.Print("-> DJSet: ", djset)

		cliques := 0
		for v := range n {
			if djset[v][0] == v && djset[v][1] == djset[v][2]*(djset[v][2]-1)/2 {
				cliques++
			}
		}
		return cliques
	}

	for _, c := range []struct {
		rst, n int
		edges  [][]int
	}{
		{3, 6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}},
		{1, 6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}},
		{2, 5, [][]int{{1, 2}, {3, 4}, {1, 4}, {2, 3}, {1, 3}, {2, 4}}},
	} {
		rst, n, edges := c.rst, c.n, c.edges
		for _, f := range []func(int, [][]int) int{countCompleteComponents, BFS, UnionFind} {
			if rst != f(n, edges) {
				t.FailNow()
			}
		}
		log.Printf(":: %d <- %v", rst, edges)
		log.Print("---")
	}
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

func Test3108(t *testing.T) {
	for _, c := range []struct {
		rst   []int
		n     int
		edges [][]int
		query [][]int
	}{
		{[]int{1, -1}, 5, [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}, [][]int{{0, 3}, {3, 4}}},
		{[]int{0}, 3, [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}, [][]int{{1, 2}}},
		{[]int{0}, 7, [][]int{{3, 0, 2}, {5, 4, 12}, {6, 3, 7}, {4, 2, 2}, {6, 2, 2}}, [][]int{{6, 0}}},
	} {
		rst, n, edges, query := c.rst, c.n, c.edges, c.query
		if !reflect.DeepEqual(rst, minimumCostWalk(n, edges, query)) {
			t.FailNow()
		}
		log.Printf(":: %v <- %v", rst, edges)
	}
}

// 3203h Find Minimum Distance After Merging Two Trees
func Test3203(t *testing.T) {
	log.Print("3 ?= ", minimumDiameterAfterMerge([][]int{{0, 1}, {0, 2}, {0, 3}}, [][]int{{0, 1}}))
	log.Print("5 ?= ", minimumDiameterAfterMerge([][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}))
}

// 3286m Find a Safe Walk Through a Grid
func Test3286(t *testing.T) {
	log.Print("true ?= ", findSafeWalk([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}, 1))
	log.Print("false ?= ", findSafeWalk([][]int{{0, 1, 1, 0, 0, 0}, {1, 0, 1, 0, 0, 0}, {0, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}}, 3))
	log.Print("true ?= ", findSafeWalk([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 5))
}

func Test3341(t *testing.T) {
	for _, c := range []struct {
		rst      int
		moveTime [][]int
	}{
		{6, [][]int{{0, 4}, {4, 4}}},
		{3, [][]int{{0, 0, 0}, {0, 0, 0}}},
		{3, [][]int{{0, 1}, {1, 2}}},

		{60, [][]int{{15, 58}, {67, 4}}},
	} {
		log.Print("** ", c.moveTime)
		if c.rst != minTimeToReach(c.moveTime) {
			t.FailNow()
		}
	}
}

func Test3342(t *testing.T) {
	for _, c := range []struct {
		rst      int
		moveTime [][]int
	}{
		{7, [][]int{{0, 4}, {4, 4}}},
		{6, [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}}},
		{4, [][]int{{0, 1}, {1, 2}}},
	} {
		log.Print("** ", c.moveTime)
		if c.rst != minTimeToReachII(c.moveTime) {
			t.FailNow()
		}
	}
}

func Test3372(t *testing.T) {
	for _, c := range []struct {
		rst            []int
		edges1, edges2 [][]int
		k              int
	}{
		{
			[]int{9, 7, 9, 8, 8},
			[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
			[][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}},
			2,
		},
		{
			[]int{6, 3, 3, 3, 3},
			[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			[][]int{{0, 1}, {1, 2}, {2, 3}},
			1,
		},
	} {
		log.Print("* ", c.k)
		if !reflect.DeepEqual(c.rst, maxTargetNodes(c.edges1, c.edges2, c.k)) {
			t.FailNow()
		}
	}
}

func Test3373(t *testing.T) {
	for _, c := range []struct {
		rst            []int
		edges1, edges2 [][]int
	}{
		{
			[]int{8, 7, 7, 8, 8},
			[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
			[][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}},
		},
		{
			[]int{3, 6, 6, 6, 6},
			[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			[][]int{{0, 1}, {1, 2}, {2, 3}},
		},
	} {
		log.Print("*")
		if !reflect.DeepEqual(c.rst, maxTargetNodesII(c.edges1, c.edges2)) {
			t.FailNow()
		}
	}
}
