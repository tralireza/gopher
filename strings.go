package gopher

import (
	"log"
	"strconv"
	"strings"
)

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
