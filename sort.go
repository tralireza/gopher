package gopher

import "slices"

// 2418 Sort the People
func sortPeople(names []string, heights []int) []string {
	type P struct {
		name   string
		height int
	}

	D := []P{}
	for i := range names {
		D = append(D, P{name: names[i], height: heights[i]})
	}
	slices.SortFunc(D, func(x, y P) int { return y.height - x.height })

	for i := range D {
		names[i] = D[i].name
	}
	return names
}
