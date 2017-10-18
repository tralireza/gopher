package gopher

// 438m Find All Anagrams in a String
func findAnagrams(s string, p string) []int {
	R := []int{}
	if len(s) < len(p) {
		return R
	}

	fP, fS := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		fP[p[i]-'a']++
		fS[s[i]-'a']++
	}

	if fS == fP {
		R = append(R, 0)
	}
	for i := len(p); i < len(s); i++ {
		fS[s[i-len(p)]-'a']--
		fS[s[i]-'a']++
		if fS == fP {
			R = append(R, i-len(p)+1)
		}
	}

	return R
}
