package gopher

// 1395m Count Number of Teams
func numTeams(rating []int) int {
	x := 0

	for m := 1; m < len(rating)-1; m++ {
		l, r := 0, 0

		for i := 0; i < m; i++ {
			if rating[i] < rating[m] {
				l++
			}
		}

		for i := m + 1; i < len(rating); i++ {
			if rating[m] < rating[i] {
				r++
			}
		}

		x += l * r                               // Rating[l] < Rating[m] < Rating[r]
		x += (m - l) * (len(rating) - m - 1 - r) // Rating[l] > Raring[m] > Rating[r]
	}

	return x
}

// 1653m Minimum Deletions to Make String Balanced
func minimumDeletions(s string) int {
	A, B := make([]int, len(s)), make([]int, len(s))

	var x int
	for i := 0; i < len(s); i++ {
		B[i] = x
		if s[i] == 'b' {
			x++
		}
	}
	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		A[i] = x
		if s[i] == 'a' {
			x++
		}
	}

	dels := len(s)
	for i := 0; i < len(s); i++ {
		if A[i]+B[i] < dels {
			dels = A[i] + B[i]
		}
	}
	return dels
}
