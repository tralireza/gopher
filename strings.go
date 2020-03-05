package gopher

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// 38 Count and Say
func countAndSay(n int) string {
	s := "1"
	for range n - 1 {
		Enc := [][2]int{}

		count, prv := 0, byte('^')
		s += "$"
		for i := range s {
			if s[i] == prv {
				count++
			} else {
				Enc = append(Enc, [2]int{count, int(prv)})
				prv = s[i]
				count = 1
			}
		}

		t := ""
		for _, e := range Enc[1:] {
			t += fmt.Sprintf("%d%c", e[0], byte(e[1]))
		}

		s = t
	}

	return s
}

// 412 Fizz Buzz
func fizzBuzz(n int) []string {
	S := []string{}

	for i := 1; i <= n; i++ {
		var s string
		switch {
		case i%15 == 0:
			s = "FizzBuzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		default:
			s = strconv.Itoa(i)
		}
		S = append(S, s)
	}

	return S
}

// 466h Count The Repetitions
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	R, M := make([]int, len(s2)+1), make([]int, len(s2)+1)

	r, m := 0, 0
	for i := 0; i < n1; i++ {
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[m] {
				m++
			}
			if m == len(s2) {
				m = 0
				r++
			}
		}

		R[i], M[i] = r, m
		log.Print("-> ", i, R, M)

		for k := 0; k < i; k++ {
			if M[k] == m {
				prv := R[k]
				pattern := (R[i] - R[k]) * ((n1 - 1 - k) / (i - k))
				rest := R[k+(n1-1-k)%(i-k)] - R[k]

				return (prv + pattern + rest) / n2
			}
		}
	}

	return R[n1-1] / n2
}

// 520 Detect Capital
func detectCapitalUse(word string) bool {
	return word[1:] == strings.ToLower(word[1:]) || word == strings.ToUpper(word)
}

// 557 Reverse Words in a String III
func reverseWords(s string) string {
	S := []string{}
	for _, sstr := range strings.Fields(s) {
		rstr := make([]byte, len(sstr))
		for l, r := 0, len(sstr)-1; l <= r; l, r = l+1, r-1 {
			rstr[l], rstr[r] = sstr[r], sstr[l]
		}
		S = append(S, string(rstr))
	}
	log.Print(" -> ", strings.Join(S, " "))

	bfr := []byte{}

	s = s + " "
	l, r := 0, 0
	for r < len(s) {
		if s[r] == ' ' {
			for x := 0; x < r-l; x++ {
				bfr = append(bfr, s[r-x-1])
			}
			bfr = append(bfr, ' ')
			l = r + 1
		}
		r++
	}
	bfr = bfr[:len(bfr)-1]

	return string(bfr)
}

// 696 Count Binary Substrings
func countBinarySubstrings(s string) int {
	count, prv, cur := 0, 0, 1
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			count += min(prv, cur)
			prv, cur = cur, 1
		} else {
			cur++
		}
	}

	return count + min(prv, cur)
}

// 806 Number of Lines To Write String
func numberOfLines(widths []int, s string) []int {
	lw, l := 0, 1
	for i := 0; i < len(s); i++ {
		w := widths[s[i]-'a']
		switch lw+w > 100 {
		case true:
			l++
			lw = w
		default:
			lw += w
		}
	}

	return []int{l, lw}
}

// 824 Goat Latin
func toGoatLatin(sentence string) string {
	W := []string{}
	ending := []byte{'a'}
	for _, w := range strings.Fields(sentence) {
		switch w[0] {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
			W = append(W, w+"ma"+string(ending))
		default:
			W = append(W, w[1:]+string(w[0])+"ma"+string(ending))
		}
		ending = append(ending, 'a')
	}

	log.Print(" -> ", strings.Join(W, " "))

	sentence += " "
	bfr, w, a := []byte{}, []byte{}, []byte{'a'}
	for i := 0; i < len(sentence); i++ {
		switch sentence[i] {
		case ' ':
			switch w[0] {
			case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
			default:
				chr := w[0]
				copy(w, w[1:])
				w[len(w)-1] = chr
			}
			bfr = append(bfr, w...)

			bfr = append(bfr, []byte("ma")...)
			bfr = append(bfr, a...)
			bfr = append(bfr, ' ')

			w = []byte{}
			a = append(a, 'a')

		default:
			w = append(w, sentence[i])
		}
	}

	return string(bfr[:len(bfr)-1])
}

// 917 Reverse Only Letters
func reverseOnlyLetters(s string) string {
	rs := []rune(s)

	l, r := 0, len(s)-1
	for l < r {
		switch {
		case !unicode.IsLetter(rs[l]):
			l++
		case !unicode.IsLetter(rs[r]):
			r--
		default:
			rs[l], rs[r] = rs[r], rs[l]
			l++
			r--
		}
	}

	return string(rs)
}

// 1163h Last Substring in Lexicographical Order
// 1 <= N <= 4*10^5
func lastSubstring_SuffixArray(s string) string {
	N := len(s)

	RadixSort := func(L [][3]int, halfK int) {
		vMax := slices.MaxFunc(L, func(a, b [3]int) int { return a[halfK] - b[halfK] })
		N := max(1, vMax[halfK])
		for r := 1; r <= N; r *= 10 {
			E, B := [][3]int{}, [10][][3]int{}
			for _, l := range L {
				if l[halfK] < 0 {
					E = append(E, l)
				} else {
					B[l[halfK]/r%10] = append(B[l[halfK]/r%10], l)
				}
			}

			copy(L, E)
			offset := len(E)
			for d := range 10 {
				copy(L[offset:], B[d])
				offset += len(B[d])
			}
		}
	}

	P := []int{}
	for i := 0; i < len(s); i++ {
		P = append(P, int(s[i]-'a'))
	}

	L := make([][3]int, N)

	k := 1
	for (k >> 1) < N {
		log.Print("-> ", k, P)

		for i := 0; i < N; i++ {
			L[i][0] = P[i]
			if i+k < N {
				L[i][1] = P[i+k]
			} else {
				L[i][1] = -1
			}
			L[i][2] = i
		}
		log.Print("-> ", k, L)

		RadixSort(L, 1)
		RadixSort(L, 0)
		log.Print("-> (R) ", k, L)

		slices.SortFunc(L, func(a, b [3]int) int {
			if a[0] == b[0] {
				return a[1] - b[1]
			}
			return a[0] - b[0]
		})
		log.Print("-> (S) ", k, L)

		for i := 0; i < N; i++ {
			if i > 0 && L[i][0] == L[i-1][0] && L[i][1] == L[i-1][1] {
				P[L[i][2]] = P[L[i-1][2]]
			} else {
				P[L[i][2]] = i
			}
		}
		log.Print("-> ", k, P)

		k <<= 1
	}

	for i := range P {
		if P[i]+1 == N {
			return s[i:]
		}
	}
	return ""
}

func lastSubstring_Trie(s string) string {
	type Trie struct {
		Children map[rune]*Trie
	}

	trie := &Trie{map[rune]*Trie{}}

	Insert := func(s string) {
		n := trie
		for _, chr := range s {
			c := n.Children[chr]
			if c == nil {
				c = &Trie{map[rune]*Trie{}}
				n.Children[chr] = c
			}
			n = c
		}
	}

	for i := 0; i < len(s)-1; i++ {
		Insert(s[i:])
	}

	GetLargest := func() string {
		bfr := bytes.Buffer{}

		Chrs := []rune("abcdefghijklmnopqrstuvwxyz")
		n := trie
	LOOP:
		for {
			for i := 25; i >= 0; i-- {
				c := n.Children[Chrs[i]]
				if c != nil {
					bfr.WriteRune(Chrs[i])
					n = c
					continue LOOP
				}
			}
			break
		}

		return bfr.String()
	}

	return GetLargest()
}

func lastSubstring(s string) string {
	{
		tStart := time.Now()
		log.Printf(":: Suffix Array -> %q [@ %v]", lastSubstring_SuffixArray(s), time.Since(tStart))
	}

	{
		tStart := time.Now()
		log.Printf(":: Trie -> %q [@ %v]", lastSubstring_Trie(s), time.Since(tStart))
	}

	n := len(s)

	i, j := 0, 1
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}

		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j += k + 1
		}
	}

	log.Print(":: ", s[i:])
	return s[i:]
}

// 1668 Maximum Repeating Substring
func maxRepeating(sequence string, word string) int {
	xRepeat := 0
	for i := 0; i <= len(sequence)-len(word); i++ {
		cur := 0
		start := i
		for i < len(sequence) && sequence[i] == word[(i-start)%len(word)] {
			i++
			if (i-start)%len(word) == 0 {
				cur++
			}
		}

		xRepeat = max(cur, xRepeat)
		i = start
	}

	repeats := 0
	for p := word; strings.Contains(sequence, p); p += word {
		repeats++
	}
	log.Printf(":: %d ~ %d", repeats, xRepeat)

	return xRepeat
}

// 1813m Sentence Similarity III
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	Source, Pattern := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	if len(Source) < len(Pattern) {
		Source, Pattern = Pattern, Source
	}

	i, j := 0, len(Source)-1
	l, r := 0, len(Pattern)-1

	for l <= r {
		if Source[i] == Pattern[l] {
			i++
			l++
		} else if Source[j] == Pattern[r] {
			j--
			r--
		} else {
			return false
		}
	}

	return true
}

// 2185 Counting Words With a Given Prefix
func prefixCount(words []string, pref string) int {
	count := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			count++
		}
	}

	return count
}

// 2379 Minimum Recolors to Get K Consecutive Black Blocks
func minimumRecolors(blocks string, k int) int {
	recolors, cur := math.MaxInt, 0

	l := 0
	for r := 0; r < len(blocks); r++ {
		switch blocks[r] {
		case 'W':
			cur++
		}

		if r-l+1 >= k {
			recolors = min(recolors, cur)

			switch blocks[l] {
			case 'W':
				cur--
			}
			l++
		}
	}

	return recolors
}

// 2405m Optimal Partition of String
func partitionString(s string) int {
	t := 1

	mask := 0
	for i := 0; i < len(s); i++ {
		if mask&(1<<(s[i]-'a')) != 0 {
			t++
			mask = 0
		}
		mask |= 1 << (s[i] - 'a')
	}

	return t
}

// 3042 Count Prefix and Suffix Pairs I
func countPrefixSuffixPairs(words []string) int {
	pairs := 0

	for i, n := range words { //Needle
		for _, h := range words[i+1:] { // Haystack
			if strings.HasPrefix(h, n) && strings.HasSuffix(h, n) {
				pairs++
			}
		}
	}

	return pairs
}

// 3330 Find the Original Typed String I
func possibleStringCount(word string) int {
	count := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			count++
		}
	}

	return count
}

// 3442 Maximum Difference Between Even and Odd Frequency I
func maxDifference(s string) int {
	F := [26]int{}
	for i := 0; i < len(s); i++ {
		F[s[i]-'a']++
	}

	oMax, eMin := 0, len(s)
	for _, f := range F {
		if f > 0 {
			switch f & 1 {
			case 1:
				oMax = max(f, oMax)
			case 0:
				eMin = min(f, eMin)
			}
		}
	}

	return oMax - eMin
}
