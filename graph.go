package gopher

import "log"

// 2392h Build a Matrix With Conditions
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	TopoSort := func(edges [][]int) []int {
		G := make([][]int, k+1) // Graph (DAG)
		D := make([]int, k+1)   // InDegree
		for _, e := range edges {
			G[e[0]] = append(G[e[0]], e[1])
			D[e[1]]++
		}

		log.Print("Graph :: ", G[1:])
		log.Print("InDegree :: ", D[1:])

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
	if len(rTopo) < k { // Cycle :: G !DAG
		return nil
	}
	log.Print("row Topological order -> ", rTopo)

	cTopo := TopoSort(colConditions)
	if len(cTopo) < k { // Cycle :: G !DAG
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
