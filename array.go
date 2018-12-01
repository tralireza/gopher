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
