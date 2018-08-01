package gopher

import "log"

// 476 Number Complement
func findComplement(num int) int {
	bits := 0
	for x := num; x > 0; x >>= 1 {
		bits++
	}
	return (1<<bits - 1) ^ num
}

// 3315m Construct the Minimum Bitwise Array II
func minBitwiseArray(nums []int) []int {
	R := []int{}

	for _, p := range nums {
		r := -1
		for i := 31; i >= 0; i-- {
			if p&(1<<i) != 0 {
				n := p & ^(1 << i)

				log.Print(" -> ", p, i, n)

				if n|(n+1) == p {
					r = n
					break
				}
			}
		}
		R = append(R, r)
	}

	return R
}
