package gopher

import (
	"log"
	"slices"
	"testing"
)

func Test38(t *testing.T) {
	for _, c := range []struct {
		rst string
		n   int
	}{
		{"1211", 4},
		{"1", 1},
	} {
		if c.rst != countAndSay(c.n) {
			t.FailNow()
		}
		log.Printf(":: %s   <- %d", c.rst, c.n)
	}
}

// 412 Fizz Buzz
func Test412(t *testing.T) {
	log.Print(" ?= ", fizzBuzz(3))
	log.Print(" ?= ", fizzBuzz(5))
	log.Print(" ?= ", fizzBuzz(15))
}

func Test466(t *testing.T) {
	for _, c := range []struct {
		rst int
		s1  string
		n1  int
		s2  string
		n2  int
	}{
		{2, "acb", 4, "ab", 2},
		{1, "acb", 1, "acb", 1},

		{12, "aaa", 20, "aaaaa", 1},
	} {
		log.Print("* ", c.s1, c.n1, c.s2, c.n2)
		if c.rst != getMaxRepetitions(c.s1, c.n1, c.s2, c.n2) {
			t.FailNow()
		}
	}
}

// 520 Detect Capital
func Test520(t *testing.T) {
	log.Print("true ?= ", detectCapitalUse("USA"))
	log.Print("false ?= ", detectCapitalUse("FlaG"))
}

// 557 Reverse Words in a String III
func Test557(t *testing.T) {
	log.Printf(`"s'teL ekat edoCteeL tsetnoc" ?= %q`, reverseWords("Let's take LeetCode contest"))
	log.Printf(`"rM gniD" ?= %q`, reverseWords("Mr Ding"))
}

// 696 Count Binary Substrings
func Test696(t *testing.T) {
	log.Print("6 ?= ", countBinarySubstrings("00110011"))
	log.Print("4 ?= ", countBinarySubstrings("10101"))
}

// 806 Number of Lines To Write String
func Test806(t *testing.T) {
	log.Print("[3 60] ?= ", numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	log.Print("[2 4] ?= ", numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

// 824 Goat Latin
func Test824(t *testing.T) {
	// 1 <= S.Length

	if toGoatLatin("I speak Goat Latin") != "Imaa peaksmaaa oatGmaaaa atinLmaaaaa" {
		t.Fail()
	}
	if toGoatLatin("The quick brown fox jumped over the lazy dog") != "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa" {
		t.Fail()
	}
}

// 917 Reverse Only Letters
func Test917(t *testing.T) {
	for _, c := range []struct {
		rst, s string
	}{
		{"j-Ih-gfE-dCba", "a-bC-dEf-ghIj"},
		{"Qedo1ct-eeLg=ntse-T!", "Test1ng-Leet=code-Q!"},
	} {
		log.Print("* ", c.s)
		if c.rst != reverseOnlyLetters(c.s) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1163(t *testing.T) {
	for _, c := range []struct {
		rst, s string
	}{
		{"zmwrjvfamgpoowncslddrkjhchqswkamnsitrcmnhn", "vmjtxddvzmwrjvfamgpoowncslddrkjhchqswkamnsitrcmnhn"},
		{"zab", "zaazaabcdezaazab"},
		{"zrziy", "zrziy"}, // 10/36
		{"aaa", "aaa"},

		{"bab", "abab"},
		{"tcode", "leetcode"},
	} {
		log.Print("* ", c.s)
		if c.rst != lastSubstring(c.s) {
			t.FailNow()
		}
	}

	RadixSort := func(nums []int) {
		log.Print("-- RadixSort: ", nums)

		N := slices.Max(nums)
		for r := 1; r <= N; r *= 10 {
			E, B := []int{}, [10][]int{}
			for _, n := range nums {
				if n < 0 {
					E = append(E, n)
				} else {
					B[n/r%10] = append(B[n/r%10], n)
				}
			}

			copy(nums, E)
			offset := len(E)
			for d := range 10 {
				copy(nums[offset:], B[d])
				offset += len(B[d])
			}
		}

		log.Print(":: RadixSort: ", nums)
	}
	RadixSort([]int{325, 7, 457, 657, -1, 1000, 839, 0, 436, 7, 10, 720, 355, -1})

	CountingSort := func(nums []int) {
		log.Print("-- CountingSort: ", nums)

		xVal, mVal := slices.Max(nums), slices.Min(nums)
		F := make([]int, xVal-mVal+1)
		for _, n := range nums {
			F[n-mVal]++
		}

		i := 0
		for n, f := range F {
			for range f {
				nums[i] = n + mVal
				i++
			}
		}

		log.Print(":: CountingSort: ", nums)
	}
	CountingSort([]int{9, 7, 8, 17, 11, 3, 9, 3, 3, 16, 3, 9, 11, 19})
}

const Input1163 = "jyqxwwxglawjvneegoxztrcyjqlduczzhgdlesnaeyialxfhtcgwkxjcdsllpqwurenryothdqzdbjmppjyvwzxobkvlrxjytmpklararqdqjjnblxaliqhjvtbzysfkbhroccnlwnslpsvkarenxfezocpdocgamvufzcfjkxijwybwgbfmnnwuuunsoupaxbylxggremxxakntirsqjwkyxkldqokrlwevrvoovoekhesvxmbnycclrdhrzzbovalhtnzdhfuyatdgeyazstiovogkiuuvsjvvofvrfwyoxydkgkvhporcxccrlcecgqakknogwyemwcfmokuflsskyevbdkmmumftzcpdonagopprxcmwwuarqxbxglrnprstubwfjmxpwdsribxcglhhzthhajimjawanewsqmwifzndqwojclkdilkisapeegpeixshskpfdnsbmfjiojelllsvuquupkwvnkgfdwreabvhyswnsnsdofccebjqmawlkqbzcrxqcvargeqvruhgypqcfbltnhswzjbjayqglgsyttnvpxrjbbotzcmoscbykzxoqoqkooycfiviewtmpyzzpicglhsydafzdzresxjeqhahsukeprzooumbltzxhmqktoypcjenuqqlkpwtvyscfcxcodnokzxpcjlimqmeltiipawblteiyaftlvefhrglstuwupkfvjzhrlvejljfahcenhnsqmmcfpnbtwrkukzncabvgyvvfqhsairahkulbejckkoapagatvkhceqswlpzijcwddrooijdcircayscwmordpckluyryrguednmhzleeklgggqujqeobgesjdbpuueenraljjecjxssdosskkbhrnykrfvumazfcjalcttxewlxiwtsojrmeakgzkwympgkdrshbiaamlwwwvacewcjgaruzmcpblpgqdyykxjyybhwwgowlcsliiitgffqdfprvrrf"

func Benchmark1163_Trie(b *testing.B) {
	for range b.N {
		lastSubstring_Trie(Input1163)
	}
}

func Benchmark1163_SuffixArray(b *testing.B) {
	for range b.N {
		lastSubstring_SuffixArray(Input1163)
	}
}

func Test1668(t *testing.T) {
	// 1 <= L(s), L(w) <= 100

	for _, c := range []struct {
		rst            int
		sequence, word string
	}{
		{2, "ababc", "ab"},
		{1, "ababc", "ba"},
		{0, "ababc", "ac"},

		{5, "aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"}, // 205/212
	} {
		if c.rst != maxRepeating(c.sequence, c.word) {
			t.Error()
		}
	}
}

// 1813m Sentence Similarity III
func Test1813(t *testing.T) {
	log.Print("true ?= ", areSentencesSimilar("Hello Jane", "Hello my name is Jane"))
	log.Print("false ?= ", areSentencesSimilar("of", "A lot of words"))
	log.Print("true ?= ", areSentencesSimilar("Eating right now", "Eating"))
}

// 2185 Counting Words With a Given Prefix
func Test2185(t *testing.T) {
	log.Print("2 ?= ", prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
	log.Print("0 ?= ", prefixCount([]string{"leetcode", "win", "loops", "success"}, "code"))
}

// 2379 Minimum Recolors to Get K Consecutive Black Blocks
func Test2379(t *testing.T) {
	log.Print("3 ?= ", minimumRecolors("WBBWWBBWBW", 7))
	log.Print("0 ?= ", minimumRecolors("WBWBBBW", 2))
}

// 2405m Optimal Partition of String
func Test2405(t *testing.T) {
	log.Print("4 ?= ", partitionString("abacaba"))
	log.Print("6 ?= ", partitionString("ssssss"))
}

// 3042 Count Prefix and Suffix Pairs I
func Test3042(t *testing.T) {
	log.Print("4 ?= ", countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}))
	log.Print("2 ?= ", countPrefixSuffixPairs([]string{"pa", "papa", "ma", "mama"}))
	log.Print("0 ?= ", countPrefixSuffixPairs([]string{"abab", "ab"}))
}

func Test3330(t *testing.T) {
	for _, c := range []struct {
		rst  int
		word string
	}{
		{5, "abbcccc"},
		{1, "abcd"},
		{4, "aaaa"},
	} {
		log.Print("* ", c.word)
		if c.rst != possibleStringCount(c.word) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test3442(t *testing.T) {
	for _, c := range []struct {
		rst int
		s   string
	}{
		{3, "aaaaabbc"},
		{1, "abcabcab"},
	} {
		if c.rst != maxDifference(c.s) {
			t.FailNow()
		}
	}
}
