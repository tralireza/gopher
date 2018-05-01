package gopher

// 319m Bulb Switcher
func bulbSwitch(n int) int {
	q := 0 // Square Root
	for q*q <= n {
		q++
	}
	return q - 1
}
