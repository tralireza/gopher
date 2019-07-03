package gopher

// 1358m Number of Substrings Containing All Three Characters
func numberOfSubstrings(s string) int {
	count := 0

	M := map[byte]int{}
	l := 0

	for r := 0; r < len(s); r++ {
		M[s[r]]++

		for len(M) == 3 {
			count += len(s) - r

			M[s[l]]--
			if M[s[l]] == 0 {
				delete(M, s[l])
			}

			l++
		}
	}

	return count
}

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
