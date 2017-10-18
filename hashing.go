package gopher

// 438m Find All Anagrams in a String
func findAnagrams(s string, p string) []int {
	fP, fS := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		fP[p[i]-'a']++
	}

	R := []int{}
	for i := 0; i < len(s); i++ {
		fS[s[i]-'a']++
		if i >= len(p) {
			fS[s[i-len(p)]-'a']--
		}
		if fS == fP {
			R = append(R, i-len(p)+1)
		}
	}
	return R
}
