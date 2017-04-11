package gopher

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
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

func Fib(n int) int {
	Mem := map[int]int{0: 0, 1: 1}
	var fib func(int) int
	fib = func(n int) int {
		if v, ok := Mem[n]; ok {
			return v
		}
		Mem[n] = fib(n-1) + fib(n-2)
		log.Printf(" -> fib(%d) %d", n, Mem[n])
		return Mem[n]
	}
	return fib(n)
}
