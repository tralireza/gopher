package gopher

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"text/tabwriter"
	"time"

	"golang.org/x/net/html"
)

func init() {
	log.SetFlags(0)
}

type ByteCounter int

func (o ByteCounter) String() string { return "@" + strconv.Itoa(int(o)) }

func (o *ByteCounter) Write(p []byte) (int, error) {
	*o += ByteCounter(len(p))
	return len(p), nil
}

type Planet struct {
	Name     string
	Order    int
	Moons    []string
	Distance int
	Mass     float64
	Radius   float64
}

var (
	SolarSystem = []*Planet{
		{"Mercury", 1, nil, 58, .055, .3830},
		{"Venus", 2, nil, 108, .815, .9499},
		{"Earth", 3, []string{"Moon"}, 150, 1, 1},
		{"Mars", 4, []string{"Phobos", "Deimos"}, 228, .107, .532},
		{"Jupiter", 5, []string{"Io", "Europa", "Ganymede", "Callisto"}, 778, 318, 10.8},
		{"Saturn", 6, []string{"Mimas", "Enceladus", "Tethys", "Dione", "Rhea", "Titan"}, 1433, 95, 8.9},
		{"Uranus", 7, []string{"Miranda", "Ariel", "Umbriel", "Titania", "Oberon"}, 2870, 14.5, 3.96},
		{"Neptune", 8, []string{"Triton"}, 4500, 17.1, 3.86},
	}
)

type ByMass []*Planet

func (o ByMass) Len() int           { return len(o) }
func (o ByMass) Less(i, j int) bool { return o[i].Mass < o[j].Mass }
func (o ByMass) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

func PrintPlanets() {
	w := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	const format = "%v\t%v\t%v\t%v\t%v\n"
	fmt.Fprintf(w, format, "Planet", "Position", "Distance", "Mass", "Radius")
	fmt.Fprintf(w, format, "======", "--------", "--------", "----", "------")
	sort.Sort(sort.Reverse(ByMass(SolarSystem)))
	for _, p := range SolarSystem {
		fmt.Fprintf(w, format, p.Name, p.Order, p.Distance, p.Mass, p.Radius)
	}
	w.Flush()
}

func SqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		log.Printf(" -> %T %[1]v", x)
		return fmt.Sprintf("%d", x)
	case float32, float64:
		return fmt.Sprintf("%g", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return fmt.Sprintf("'%s'", x)
	default:
		return fmt.Sprintf("'%v'", x)
	}
}

// Non-Blocking Cache with Duplicate Suppression
func NewNBCacheDS(f func(string) (interface{}, error)) *fCache {
	return &fCache{f: f, cache: map[string]*fResult{}}
}

type fResult struct {
	bdy   interface{}
	err   error
	ready chan struct{}
}
type fCache struct {
	f     func(string) (interface{}, error)
	cache map[string]*fResult
	mtx   sync.Mutex
}

func (p *fCache) Get(url string) (interface{}, error) {
	p.mtx.Lock()
	if r, ok := p.cache[url]; ok {
		p.mtx.Unlock()
		<-r.ready // wait for data to be ready
		return r.bdy, r.err
	}

	r := fResult{ready: make(chan struct{})}
	p.cache[url] = &r
	p.mtx.Unlock()
	r.bdy, r.err = p.f(url)
	close(r.ready) // data is ready...
	return r.bdy, r.err
}

func httpGet(url string) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	rq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rsp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	bfr := bytes.Buffer{}
	if _, err := io.Copy(&bfr, rsp.Body); err != nil {
		return nil, err
	}
	return &bfr, nil
}

func hrefXtr(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = hrefXtr(links, c)
	}
	return links
}

func Fib(n int) int {
	rCalls, Mem := 0, map[int]int{0: 0, 1: 1}
	var fib func(int) int
	fib = func(n int) int {
		rCalls++
		if v, ok := Mem[n]; ok {
			return v
		}
		Mem[n] = fib(n-1) + fib(n-2)
		log.Printf(" -> fib(%d) %d [%d]", n, Mem[n], rCalls)
		return Mem[n]
	}
	return fib(n)
}

// 48m Rotate Image
func rotate(matrix [][]int) {
	N := len(matrix)

	// Transpose: M -> Mt
	for r := 0; r < N; r++ {
		for c := r + 1; c < N; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	// column exchange
	for r := 0; r < N; r++ {
		for c := 0; c < N/2; c++ {
			matrix[r][c], matrix[r][N-c-1] = matrix[r][N-c-1], matrix[r][c]
		}
	}
}

// 49m Group Anagrams
func groupAnagrams(strs []string) [][]string {
	aGrp := map[[26]int][]string{}
	for _, s := range strs {
		w := [26]int{}
		for i := 0; i < len(s); i++ {
			w[s[i]-'a']++
		}
		log.Print(s, " -> ", w)

		aGrp[w] = append(aGrp[w], s)
	}

	r := [][]string{}
	for _, ag := range aGrp {
		r = append(r, ag)
	}
	return r
}

// 73m Set Matrix Zeroes
func setZeroes(matrix [][]int) {
	M, N := len(matrix), len(matrix[0])

	r0, c0 := false, false
	for r := 0; r < M && !c0; r++ {
		c0 = c0 || matrix[r][0] == 0
	}
	for c := 0; c < N && !r0; c++ {
		r0 = r0 || matrix[0][c] == 0
	}

	for r := 1; r < M; r++ {
		for c := 1; c < N; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c], matrix[r][0] = 0, 0
			}
		}
	}

	for r := 1; r < M; r++ {
		for c := 1; c < N; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if r0 {
		for c := 0; c < N; c++ {
			matrix[0][c] = 0
		}
	}
	if c0 {
		for r := 0; r < M; r++ {
			matrix[r][0] = 0
		}
	}
}

// 207m Course Schedule
func canFinish(numCourses int, prerequisites [][]int) bool {
	Graph := make([][]int, numCourses)
	for _, e := range prerequisites {
		v, u := e[0], e[1]
		Graph[v] = append(Graph[v], u)
	}
	log.Print("Graph :: lsAdj -> ", Graph)

	Vis := make([]bool, numCourses)

	Comp := []int{} // Connected Components
	for n := range numCourses {
		if Vis[n] || len(Graph[n]) == 0 {
			continue
		}
		Comp = append(Comp, n)

		v, Q := n, []int{n}
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			Vis[v] = true
			for _, u := range Graph[v] {
				if !Vis[u] {
					Q = append(Q, u)
				}
			}
		}
	}
	log.Print("Components :: ", Comp)

	var t int
	var T [][2]int

	var DFS func(int)
	DFS = func(v int) {
		t++
		T[v][0] = t
		Vis[v] = true
		for _, u := range Graph[v] {
			if !Vis[u] {
				DFS(u)
			}
		}
		t++
		T[v][1] = t
	}

	cycle := false

	var CheckCycle func(int)
	CheckCycle = func(v int) {
		if cycle {
			return // done!
		}
		Vis[v] = true
		for _, u := range Graph[v] {
			if T[u][1] >= T[v][1] { // Cross or Back EDGE
				log.Print("--- Cross/Back Edge (cycle): ", v, T[v], " -> ", u, T[u])
				cycle = true
				return
			}
			if !Vis[u] {
				CheckCycle(u)
			}
		}
	}

	for _, n := range Comp {
		t, T = 0, make([][2]int, numCourses)
		Vis = make([]bool, numCourses)
		DFS(n)
		log.Print(n, " -> ", T)

		Vis = make([]bool, numCourses)
		CheckCycle(n)
	}

	return !cycle
}

// 763m Partition Labels
func partitionLabels(s string) []int {
	lP := []int{}

	lMap := [26]int{}
	for i, r := range s {
		lMap[byte(r)-'a'] = i
	}
	log.Print(len(s), " :: ", lMap)

	pSize := 0
	var iMax int
	for i := 0; i < len(s); i++ {
		pSize++
		iMax = max(iMax, lMap[s[i]-'a'])
		if i == iMax {
			lP = append(lP, pSize)
			pSize = 0
		}
	}

	return lP
}

// 1051 Height Checker
func heightChecker(heights []int) int {
	// 1 <= heights[i] <= 100
	fH := make([]int, 100+1)
	for _, h := range heights {
		fH[h]++
	}

	hSort := make([]int, len(heights))
	i := 0
	for h, f := range fH {
		for range f {
			hSort[i] = h
			i++
		}
	}

	x := 0
	for i := range hSort {
		if hSort[i] != heights[i] {
			x++
		}
	}
	return x
}
