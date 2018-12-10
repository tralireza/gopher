// Array
package gopher

// 485 Max Consecutive Ones
func findMaxConsecutiveOnes(nums []int) int {
	tCur, tMax := 0, 0
	for _, n := range nums {
		if tCur+n > tCur {
			tCur++
		} else {
			tCur = 0
		}

		tMax = max(tCur, tMax)
	}

	return tMax
}

// 1437 Check If All 1's Are at Least Length K Places Away
func kLengthApart(nums []int, k int) bool {
	dist := k
	for _, n := range nums {
		switch n {
		case 1:
			if dist < k {
				return false
			}
			dist = 0
		case 0:
			dist++
		}
	}

	return true
}
