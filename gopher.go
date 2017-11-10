package gopher

import (
	"bytes"
	"container/list"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
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
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

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

// 3m Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	Mem := [256]int{}

	ls, lsCur := 0, 0

	l := 0
	for r := range s {
		if Mem[s[r]] == 0 {
			lsCur++
		}

		Mem[s[r]]++

		for Mem[s[r]] > 1 {
			Mem[s[l]]--
			if Mem[s[l]] == 0 {
				lsCur--
			}
			l++
		}

		ls = max(lsCur, ls)
	}

	return ls
}

// 20 Valid Parentheses
func isValid(s string) bool {
	Q := []byte{}
	Pair := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		if len(Q) > 0 && s[i] == Pair[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		} else {
			Q = append(Q, s[i])
		}
	}
	return len(Q) == 0
}

// 31m Next Permutation
func nextPermutation(nums []int) {
	// N0 N1 N2 .. Nr-1<Nr .. Nn
	//                  |------| descending

	for r := len(nums) - 1; r > 0; r-- {
		if nums[r-1] < nums[r] {
			j := r
			for ; j < len(nums) && nums[j] > nums[r-1]; j++ {
			}
			nums[r-1], nums[j-1] = nums[j-1], nums[r-1]

			// Reverse: descending section
			l, r := r, len(nums)-1
			for l < r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}

			return
		}
	}

	// Wrap
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
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
			if T[u][1] >= T[v][1] { // Back EDGE
				log.Print("== Back Edge (cycle): ", v, T[v], " -> ", u, T[u])
				cycle = true
				return
			}
			if !Vis[u] {
				CheckCycle(u)
			}
		}
	}

	t, T = 0, make([][2]int, numCourses)
	Vis = make([]bool, numCourses)
	for n := range numCourses {
		if !Vis[n] {
			DFS(n)
		}
	}
	log.Print("Discovery/Finishing :: ", T)
	Vis = make([]bool, numCourses)
	for _, n := range Comp {
		CheckCycle(n)
	}

	return !cycle
}

// 238m Product of Array Except Self
func productExceptSelf(nums []int) []int {
	mSum := make([]int, len(nums))
	mSum[0] = 1
	for i := 1; i < len(nums); i++ {
		mSum[i] = nums[i-1] * mSum[i-1]
	}
	log.Print(mSum)

	m := 1
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], m = mSum[i]*m, m*nums[i]
	}

	return nums
}

// 273h Integer to English Words
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	Unit := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	Teen := []string{"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	Ten := []string{"Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	var Convert func(int) string
	Convert = func(n int) string {
		if 1 <= n && n <= 9 {
			return Unit[n-1]
		}
		if 11 <= n && n <= 19 {
			return Teen[n-11]
		}
		if n%10 == 0 && n < 100 {
			return Ten[n/10-1]
		}
		if 20 < n && n < 100 {
			return Ten[n/10-1] + " " + Unit[n%10-1]
		}

		if n >= 100 {
			if n%100 > 0 {
				return Unit[n/100-1] + " Hundred " + Convert(n%100)
			}
			return Unit[n/100-1] + " Hundred"
		}
		return ""
	}

	W := []string{}
	p, P := 0, []string{"", "Thousand", "Million", "Billion"}
	for num > 0 {
		n := num % 1000
		num /= 1000
		p++

		if n == 0 {
			continue
		}

		w := Convert(n)
		if p-1 > 0 {
			w += " " + P[p-1]
		}
		W = append([]string{w}, W...)
	}

	return strings.Join(W, " and ")
}

// 287m Find the Duplicate Number
func findDuplicate(nums []int) int {
	Mem := make([]bool, len(nums))
	for _, n := range nums {
		if Mem[n] {
			return n
		}
		Mem[n] = true
	}
	return -1
}

// 350 Intersection of Two Arrays II
func intersect(nums1 []int, nums2 []int) []int {
	M1, M2 := map[int]int{}, map[int]int{}

	for _, n := range nums1 {
		M1[n]++
	}
	for _, n := range nums2 {
		M2[n]++
	}

	R := []int{}
	for n, f := range M1 {
		for range min(f, M2[n]) {
			R = append(R, n)
		}
	}
	return R
}

// 394m Decode String
func decodeString(s string) string {
	i := 0

	ParseK := func() int {
		k := 0
		for s[i] >= '0' && s[i] <= '9' {
			k = 10*k + int(s[i]-'0')
			i++
		}
		return k
	}

	var Decode func() string
	Decode = func() string {
		var dStr strings.Builder

		for i < len(s) {
			switch {
			case s[i] >= 'a' && s[i] <= 'z':
				dStr.WriteByte(s[i])
				i++
				continue
			case s[i] == ']':
				i++
				return dStr.String()
			}

			k := ParseK()
			i++ // '['
			dStr.WriteString(strings.Repeat(Decode(), k))
		}

		return dStr.String()
	}

	return Decode()
}

// 697 Degree of an Array
func findShortestSubArray(nums []int) int {
	// 1 <= nums[i] <= 49,999
	D := make([]int, 50_000)
	for _, n := range nums {
		D[n]++
	}

	lM, dX := len(D)+1, slices.Max(D)
	for n, d := range D {
		if d != dX {
			continue
		}

		l, r := 0, len(nums)-1
		for nums[l] != n {
			l++
		}
		for nums[r] != n {
			r--
		}

		lM = min(r-l+1, lM)
	}

	return lM
}

// 726h Number of Atoms
func countOfAtoms(formula string) string {
	i := 0

	Count := func() int {
		n := 0
		for i < len(formula) && formula[i] >= '0' && formula[i] <= '9' {
			n = 10*n + int(formula[i]-'0')
			i++
		}
		if n > 0 {
			return n
		}
		return 1
	}

	var W func() map[string]int
	W = func() map[string]int {
		M := map[string]int{}

		for i < len(formula) {
			switch {
			case formula[i] == ')':
				i++
				return M

			case formula[i] == '(':
				i++
				rM := W()

				n := Count() // multiplier for (...)
				for e := range rM {
					M[e] += n * rM[e]
				}

			case formula[i] >= 'A' && formula[i] <= 'Z':
				bfr := []byte{formula[i]}
				i++
				for i < len(formula) && formula[i] >= 'a' && formula[i] <= 'z' {
					bfr = append(bfr, formula[i])
					i++
				}

				M[string(bfr)] += Count()
			}
		}

		return M
	}

	M := W()

	E := []string{}
	for e := range M {
		E = append(E, e)
	}
	slices.Sort(E)

	sbr := strings.Builder{}
	for _, e := range E {
		sbr.WriteString(e)
		if M[e] > 1 {
			sbr.WriteString(strconv.Itoa(M[e]))
		}
	}
	return sbr.String()
}

// 739m Daily Temperatures
func dailyTemperatures(temperatures []int) []int {
	r := make([]int, len(temperatures))

	Stack := []int{}
	for i, t := range temperatures {
		for len(Stack) > 0 && t > temperatures[Stack[len(Stack)-1]] {
			j := Stack[len(Stack)-1]
			r[j] = i - j
			Stack = Stack[:len(Stack)-1]
		}
		Stack = append(Stack, i)
	}

	return r
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

// 995h Minimum Number of K Consecutive Bit Flips
func minKBitFlips(nums []int, k int) int {
	x := 0

	D := []int{} // Deque
	fflag := 0   // flip flag
	for i := range nums {
		if i >= k {
			fflag ^= D[0]
		}

		if nums[i] == fflag {
			if i+k > len(nums) {
				return -1
			}

			x++
			fflag ^= 1
			D = append(D, 1)
		} else {
			D = append(D, 0)
		}

		if len(D) > k {
			D = D[1:]
		}
	}

	return x
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

// 1052 Grumpy Bookstore Owner
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	uMax := 0

	winUnhappy := 0
	l := 0
	for r := range customers {
		if grumpy[r] == 1 {
			winUnhappy += customers[r]
		}
		if r-l+1 > minutes {
			if grumpy[l] == 1 {
				winUnhappy -= customers[l]
			}
			l++
		}
		uMax = max(winUnhappy, uMax)
	}

	t := 0
	for i := range customers {
		if grumpy[i] == 0 {
			t += customers[i]
		}
	}
	return t + uMax
}

// 1248m Count Number of Nice Subarrays
func numberOfSubarrays(nums []int, k int) int {
	Mem := map[int]int{} // Track the number of subarrays with sum of values
	Mem[0] = 1

	x := 0
	pSum := 0 // running Prefix Sum -> number of Odd numbers
	for i := range nums {
		pSum += nums[i] & 1
		x += Mem[pSum-k]
		Mem[pSum]++
	}
	return x
}

// 1509m Minimum Difference Between Largest and Smallest Value in Three Moves
func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	slices.Sort(nums)
	m := math.MaxInt
	for l := 0; l < 4; l++ {
		m = min(nums[len(nums)-4+l]-nums[l], m)
	}
	return m
}

// 1518 Water Bottles
func numWaterBottles(numBottles int, numExchange int) int {
	x := 0

	empty := 0
	for numBottles > 0 {
		x += numBottles
		empty += numBottles
		numBottles = empty / numExchange
		empty %= numExchange
	}

	return x
}

// 1550 Three Consecutive Odds
func threeConsecutiveOdds(arr []int) bool {
	counter := 0
	for _, n := range arr {
		if n&1 == 1 {
			counter++
			if counter == 3 {
				return true
			}
		} else {
			counter = 0
		}
	}
	return false
}

// 1579h Remove Max Number of Edges to Keep Graph Fully Traversable
type DVal1579 struct{ p, r int } // p: parent, r: rank
type DJS1579 []*DVal1579

func (o *DJS1579) Connected(x, y int) bool { return o.FindSet(x) == o.FindSet(y) }
func (o *DJS1579) FindSet(x int) int {
	v := (*o)[x]
	if v.p != x {
		v.p = o.FindSet(v.p)
	}
	return v.p
}
func (o *DJS1579) Union(x, y int) {
	x, y = o.FindSet(x), o.FindSet(y)
	if x == y {
		return
	}

	X, Y := (*o)[x], (*o)[y]
	if X.r >= Y.r {
		Y.p = x
		if X.r == Y.r {
			X.r++
		}
	} else {
		X.p = y
	}
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	type DVal = DVal1579
	type DJS = DJS1579

	A, B := make(DJS, n+1), make(DJS, n+1)
	for i := range n + 1 {
		A[i], B[i] = &DVal{p: i}, &DVal{p: i}
	}

	eA, eB, eG := 0, 0, 0
	for _, e := range edges {
		if e[0] == 3 && !A.Connected(e[1], e[2]) {
			eG++
			A.Union(e[1], e[2])
			B.Union(e[1], e[2])
		}
	}
	for _, e := range edges {
		switch e[0] {
		case 1:
			if !A.Connected(e[1], e[2]) {
				eA++
				A.Union(e[1], e[2])
			}
		case 2:
			if !B.Connected(e[1], e[2]) {
				eB++
				B.Union(e[1], e[2])
			}
		}
	}

	source := 1
	for v := 2; v <= n; v++ {
		if !A.Connected(source, v) {
			return -1
		}
		if !B.Connected(source, v) {
			return -1
		}
	}

	return len(edges) - (eA + eB + eG)
}

// 1701m Average Waiting Time
func averageWaitingTime(customers [][]int) float64 {
	waiting, clock := 0, 0

	for i := range customers {
		arrive, time := customers[i][0], customers[i][1]
		if clock < arrive {
			clock = arrive
		}
		clock += time
		waiting += clock - arrive
	}

	return float64(waiting) / float64(len(customers))
}

// 1717m Maximum Score From Removing Substrings
func maximumGain(s string, x int, y int) int {
	g := 0

	Tx, Ty := "ab", "ba"
	if y > x {
		x, y = y, x
		Tx, Ty = Ty, Tx
	}
	log.Printf("Tokens :: %s:%d %s:%d", Tx, x, Ty, y)

	Q := []byte{}
	Greedy := func(tkn string, v int) {
		for i := 0; i < len(s); i++ {
			if len(Q) > 0 && tkn[0] == Q[len(Q)-1] && tkn[1] == s[i] {
				g += v
				Q = Q[:len(Q)-1]
			} else {
				Q = append(Q, s[i])
			}
		}
	}

	Greedy(Tx, x)

	s = string(Q)
	Q = []byte{}
	Greedy(Ty, y)

	return g
}

// 1823m Find the Winner of the Circular Game
func findTheWinner(n int, k int) int {
	// 1 <= k <= n
	Q := list.New()
	for p := range n {
		Q.PushBack(p + 1)
	}

	for Q.Len() > 1 {
		for range k - 1 {
			Q.PushBack(Q.Remove(Q.Front()))
		}
		Q.Remove(Q.Front())
	}

	return Q.Front().Value.(int)
}

// 1190m Reverse Substrings Between Each Pair of Parentheses
func reverseParentheses(s string) string {
	Wtr := [][]byte{{}}

	for _, c := range s {
		switch c {
		case '(':
			Wtr = append(Wtr, []byte{})

		case ')':
			bfr := Wtr[len(Wtr)-1]
			l, r := 0, len(bfr)-1
			for l < r {
				bfr[l], bfr[r] = bfr[r], bfr[l]
				l++
				r--
			}
			Wtr = Wtr[:len(Wtr)-1]
			Wtr[len(Wtr)-1] = append(Wtr[len(Wtr)-1], bfr...)

		default:
			Wtr[len(Wtr)-1] = append(Wtr[len(Wtr)-1], byte(c))
		}
	}

	return string(Wtr[0])
}

// 2058m Find the Minimum and Maximum Number of Nodes Between Critical Points
func nodesBetweenCriticalPoints(head *ListNode) []int {
	// 1 <= Nodes <= 10^5
	dX, dM := -1, 100_000

	first, prv, cur := -1, -1, 0

	var p *ListNode
	for n := head; n != nil; n = n.Next {
		if p != nil && n.Next != nil {
			if p.Val < n.Val && n.Next.Val < n.Val || // local Maxima
				p.Val > n.Val && n.Next.Val > n.Val { // local Minima
				if first == -1 && prv == -1 {
					first, prv = cur, cur
				} else {
					dX = cur - first
					dM = min(cur-prv, dM)
					prv = cur
				}
			}
		}

		cur++
		p = n
	}

	if dM == 100_000 {
		return []int{-1, -1}
	}
	return []int{dM, dX}
}

// 2181m Merge Nodes in Between Zeros
func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Val == 0 && head.Next == nil {
		return nil
	}

	mVal := 0
	n := head.Next
	for n.Val != 0 {
		mVal += n.Val
		n = n.Next
	}
	return &ListNode{mVal, mergeNodes(n)}
}

// 2192m All Ancestors of a Node in a Directed Acyclic Graph
func getAncestors(n int, edges [][]int) [][]int {
	G := make([][]int, n)
	for _, e := range edges {
		G[e[1]] = append(G[e[1]], e[0]) // Inverse/Transpose Graph
	}

	var Vis []bool
	var r []int

	var DFS func(int)
	DFS = func(v int) {
		for _, u := range G[v] {
			if !Vis[u] {
				r = append(r, u)
				Vis[u] = true
				DFS(u)
			}
		}
	}

	R := [][]int{}
	for v := range n {
		Vis, r = make([]bool, n), []int{}
		Vis[v] = true
		DFS(v)
		slices.Sort(r)
		R = append(R, r)
	}
	return R
}

// 2285m Maximum Total Importance of Roads
func maximumImportance(n int, roads [][]int) int64 {
	D := make([]int, n)
	for _, e := range roads {
		D[e[0]]++
		D[e[1]]++
	}

	slices.SortFunc(D, func(x, y int) int { return y - x })
	log.Print(D)

	x := int64(0)
	for _, d := range D {
		x += int64(d * n)
		n--
	}
	return x
}

// 2496 Maximum Value of a String in an Array
func maximumValue(strs []string) int {
	x := 0
	for _, s := range strs {
		if v, err := strconv.Atoi(s); err != nil {
			x = max(len(s), x)
		} else {
			x = max(v, x)
		}
	}
	return x
}

// 2751h Robot Collisions
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	type Robot struct {
		p, h, i int
		dir     byte
	}

	Robs := []Robot{}
	for i := range directions {
		Robs = append(Robs, Robot{positions[i], healths[i], i, directions[i]})
	}
	slices.SortFunc(Robs, func(x, y Robot) int { return x.p - y.p })

	Q := []Robot{}
	for i := range Robs {
		if Robs[i].dir == 'R' {
			Q = append(Q, Robs[i])
		} else {
			for len(Q) > 0 && Q[len(Q)-1].dir == 'R' && Q[len(Q)-1].h < Robs[i].h {
				Robs[i].h--
				Q = Q[:len(Q)-1]
			}

			if len(Q) > 0 && Q[len(Q)-1].dir == 'R' {
				if Q[len(Q)-1].h == Robs[i].h {
					Q = Q[:len(Q)-1]
				} else {
					Q[len(Q)-1].h--
				}
			} else {
				Q = append(Q, Robs[i])
			}
		}
	}

	slices.SortFunc(Q, func(x, y Robot) int { return y.i - x.i })
	R := []int{}
	for i := len(Q) - 1; i >= 0; i-- {
		R = append(R, Q[i].h)
	}
	return R
}

// 3191m Minimum Operations to Make Binary Array Elements Equal to One I
func minOperations(nums []int) int {
	fflip, Q, k := 0, []int{}, 3
	x := 0
	for i := range nums {
		if i >= k {
			fflip ^= Q[0]
		}

		if nums[i] == fflip {
			if i+k > len(nums) {
				return -1
			}
			x++
			fflip ^= 1
			Q = append(Q, 1) // flip
		} else {
			Q = append(Q, 0) // no-flip
		}

		if len(Q) > k {
			Q = Q[1:]
		}
	}
	return x
}
