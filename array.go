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

// 1752 Check If Array Is Sorted and Rotated
func check(nums []int) bool {
	inversions := 0
	if nums[0] < nums[len(nums)-1] {
		inversions++
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			inversions++
		}
	}
	return inversions <= 1
}
