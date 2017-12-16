package gopher

// 476 Number Complement
func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	bits := 0
	for x := num; x > 0; x >>= 1 {
		bits++
	}

	return (1<<bits - 1) ^ num
}
