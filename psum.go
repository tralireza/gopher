// gopher :: Prefix Sum
package gopher

import (
	"log"
	"slices"
)

// 1013 Partition Array Into Three Parts With Equal Sum
func canThreePartsEqualSum(arr []int) bool {
	aSum := 0
	for _, n := range arr {
		aSum += n
	}

	if aSum%3 != 0 {
		return false
	}
	t := aSum / 3

	counter, curSum := 0, 0
	for _, n := range arr {
		curSum += n
		if curSum == t {
			counter++
			curSum = 0
		}
	}
	return counter >= 3
}

// 1310m XOR Queries of a Subarray
func xorQueries(arr []int, queries [][]int) []int {
	pSum := make([]int, len(arr)+1)
	for i := range arr {
		pSum[i+1] = pSum[i] ^ arr[i]
	}

	R := []int{}
	for _, query := range queries {
		R = append(R, pSum[query[1]+1]^pSum[query[0]])
	}
	return R
}

// 1422 Maximum Score After Splitting a String
func maxScore(s string) int {
	rone := 0
	for r := len(s) - 1; r >= 0; r-- {
		if s[r] == '1' {
			rone++
		}
	}

	xScore := 0

	lzero := int((s[0] - '0') ^ 1)
	for i := 1; i < len(s); i++ {
		xScore = max(xScore, lzero+rone)
		switch s[i] {
		case '0':
			lzero++
		case '1':
			rone--
		}
	}

	return xScore
}

// 1769m Minimum Number of Operations to Move All Balls to Each Box
func minMoveOperations(boxes string) []int {
	N := len(boxes)

	rBs, rOps := make([]int, N+1), make([]int, N+1)
	for i := N - 1; i >= 0; i-- {
		rBs[i] = rBs[i+1] + int(boxes[i]-'0')
		rOps[i] = rBs[i] + rOps[i+1]
	}

	lBs, lOps := make([]int, N+1), make([]int, N+1)
	for i := 1; i <= N; i++ {
		lBs[i] = lBs[i-1] + int(boxes[i-1]-'0')
		lOps[i] = lBs[i] + lOps[i-1]
	}

	log.Print(" >>> ", lBs, lOps)
	log.Print(" <<< ", rBs, rOps)

	R := []int{}
	for i := 0; i < N; i++ {
		R = append(R, lOps[i]+rOps[i+1])
	}

	log.Print(" -> ", R)

	return R
}

// 1991 Find the Middle Index in Array
func findMiddleIndex(nums []int) int {
	rSum := 0
	for _, n := range nums {
		rSum += n
	}

	lSum := 0
	for i, n := range nums {
		rSum -= n
		if rSum == lSum {
			return i
		}
		lSum += n
	}
	return -1
}

// 2134m Minimum Swaps to Group All 1's Together II
func minSwaps(nums []int) int {
	ones := 0
	for _, n := range nums {
		ones += n
	}

	circular := make([]int, len(nums)*2)
	copy(circular, nums)
	copy(circular[len(nums):], nums)

	log.Print(nums, " -> ", circular)

	// Prefix Sum for zeros
	pSum := make([]int, 2*len(nums)+1)
	for i := range circular {
		pSum[i+1] = pSum[i]
		if circular[i] == 0 {
			pSum[i+1]++
		}
	}

	ops := len(nums) - ones
	for r := ones - 1; r < len(circular); r++ {
		ops = min(pSum[r+1]-pSum[r-ones+1], ops)
	}
	return ops
}

// 2381m Shifting Letters II
func shiftingLetters(s string, shifts [][]int) string {
	S := [][2]int{}
	for _, v := range shifts {
		x := v[2]
		if x == 0 {
			x = -1
		}
		S = append(S, [2]int{v[0], x})
		S = append(S, [2]int{v[1] + 1, -x})
	}

	slices.SortFunc(S, func(x, y [2]int) int { return x[0] - y[0] })
	log.Print(" -> ", S)

	bfr := make([]byte, 0, len(s))
	p, r := 0, 0

	for i := 0; i < len(s); i++ {
		for p < len(S) && S[p][0] <= i {
			r += S[p][1]
			p++
		}

		r = (r%26 + 26) % 26
		bfr = append(bfr, 'a'+(s[i]-'a'+byte(r)+26)%26)
	}

	return string(bfr)
}

// 2559m Count Vowel Strings in Ranges
func vowelStrings(words []string, queries [][]int) []int {
	M := [26]bool{}
	for i := 0; i < 5; i++ {
		M["aeiou"[i]-'a'] = true
	}

	pSum := make([]int, len(words)+1)
	for i, w := range words {
		pSum[i+1] = pSum[i]
		if M[w[0]-'a'] && M[w[len(w)-1]-'a'] {
			pSum[i+1]++
		}
	}

	log.Print(" -> ", pSum)

	R := []int{}
	for _, Q := range queries {
		R = append(R, pSum[Q[1]+1]-pSum[Q[0]])
	}
	return R
}

// 2574 Left and Right Sum Difference
func leftRightDifference(nums []int) []int {
	lSum, rSum := 0, 0
	for _, n := range nums {
		rSum += n
	}

	R := []int{}
	for _, n := range nums {
		rSum -= n
		r := rSum - lSum
		if r < 0 {
			r *= -1
		}
		R = append(R, r)
		lSum += n
	}
	return R
}

// 2670 Find the Distinct Difference Array
func distinctDifferenceArray(nums []int) []int {
	rM := map[int]int{}
	for _, n := range nums {
		rM[n]++
	}

	lM := map[int]struct{}{}
	R := []int{}
	for _, n := range nums {
		lM[n] = struct{}{}

		if rM[n] > 0 {
			rM[n]--
			if rM[n] == 0 {
				delete(rM, n)
			}
		}

		R = append(R, len(lM)-len(rM))
	}
	return R
}

// 3152m Special Array II
func isArraySpecial(nums []int, queries [][]int) []bool {
	pSum := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i]&1 != nums[i-1]&1 {
			pSum[i] = pSum[i-1] + 1
		}
	}

	R := []bool{}
	for _, v := range queries {
		start, end := v[0], v[1]
		v := false
		if pSum[end]-pSum[start] == end-start {
			v = true
		}
		R = append(R, v)
	}
	return R
}

// 3179m Find the N-th Value After K Seconds
func valueAfterKSeconds(n int, k int) int {
	pSum := make([]int, n)
	for i := range n {
		pSum[i] = 1
	}

	for range k {
		for i := range pSum[:n-1] {
			pSum[i+1] += pSum[i]
			pSum[i+1] %= 1e9 + 7
		}
	}
	return pSum[n-1]
}
