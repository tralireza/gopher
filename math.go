package gopher

// 319m Bulb Switcher
func bulbSwitch(n int) int {
	if n <= 1 {
		return n
	}

	q := 2 // Square Root
	for q*q <= n {
		q++
	}
	return q - 1
}
