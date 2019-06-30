package gopher

// 3208m Alternating Groups II
func numberOfAlternatingGroups(colors []int, k int) int {
	groups, wSize := 0, 1

	prvColor := colors[0]
	for i := 1; i < len(colors)+k-1; i++ {
		curColor := colors[i%len(colors)]

		switch curColor {
		case prvColor:
			wSize = 1
		default:
			wSize++
			if wSize >= k {
				groups++
			}
			prvColor = curColor
		}

	}

	return groups
}
