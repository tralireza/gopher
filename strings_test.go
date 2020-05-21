package gopher

import (
	"log"
	"reflect"
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

func Test748(t *testing.T) {
	for _, c := range []struct {
		rst, licensePlate string
		words             []string
	}{
		{"steps", "1s3 PSt", []string{"step", "steps", "stripe", "stepple"}},
		{"pest", "1s3 456", []string{"looks", "pest", "stew", "show"}},
	} {
		log.Printf("* %q %q", c.licensePlate, c.words)
		if c.rst != shortestCompletingWord(c.licensePlate, c.words) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test771(t *testing.T) {
	for _, c := range []struct {
		rst            int
		jewels, stones string
	}{
		{3, "aA", "aAAbbbb"},
		{0, "z", "ZZ"},
	} {
		log.Printf("* %q %q", c.jewels, c.stones)
		if c.rst != numJewelsInStones(c.jewels, c.stones) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test804(t *testing.T) {
	for _, c := range []struct {
		rst   int
		words []string
	}{
		{2, []string{"gin", "zen", "gig", "msg"}},
		{1, []string{"a"}},
	} {
		log.Printf("* %q", c.words)
		if c.rst != uniqueMorseRepresentations(c.words) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

// 806 Number of Lines To Write String
func Test806(t *testing.T) {
	log.Print("[3 60] ?= ", numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	log.Print("[2 4] ?= ", numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func Test819(t *testing.T) {
	for _, c := range []struct {
		rst, paragraph string
		banned         []string
	}{
		{"ball", "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}},
		{"a", "a.", []string{}},
	} {
		log.Printf("* %q %q", c.paragraph, c.banned)
		if c.rst != mostCommonWord(c.paragraph, c.banned) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
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

func Test859(t *testing.T) {
	for _, c := range []struct {
		rst     bool
		s, goal string
	}{
		{true, "ab", "ba"},
		{false, "ab", "ab"},
		{true, "aa", "aa"},
	} {
		log.Printf("* %q %q", c.s, c.goal)
		if c.rst != buddyStrings(c.s, c.goal) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
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

func Test927(t *testing.T) {
	// 1 <= N, T <= 1000

	for _, c := range []struct {
		rst         bool
		name, typed string
	}{
		{true, "alex", "aaleex"},
		{false, "saeed", "ssaaedd"},

		{false, "alex", "aaleexa"},
		{false, "alexd", "ale"},
	} {
		log.Printf("* %q %q", c.name, c.typed)
		if c.rst != isLongPressedName(c.name, c.typed) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test929(t *testing.T) {
	for _, c := range []struct {
		rst    int
		emails []string
	}{
		{2, []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}},
		{3, []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}},
	} {
		log.Print("* ", c.emails)
		if c.rst != numUniqueEmails(c.emails) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test953(t *testing.T) {
	for _, c := range []struct {
		rst   bool
		words []string
		order string
	}{
		{true, []string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"},
		{false, []string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"},
		{false, []string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"},
	} {
		log.Printf("* %q %q", c.words, c.order)
		if c.rst != isAlienSorted(c.words, c.order) {
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

func Test1189(t *testing.T) {
	for _, c := range []struct {
		rst  int
		text string
	}{
		{1, "nlaebolko"},
		{2, "loonbalxballpoon"},
	} {
		log.Printf("* %q", c.text)
		if c.rst != maxNumberOfBalloons(c.text) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1446(t *testing.T) {
	for _, c := range []struct {
		rst int
		s   string
	}{
		{5, "abbcccddddeeeeedcba"},
	} {
		log.Printf("* %q", c.s)
		if c.rst != maxPower(c.s) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1470(t *testing.T) {
	for _, c := range []struct {
		rst, nums []int
		n         int
	}{
		{[]int{2, 3, 5, 4, 1, 7}, []int{2, 5, 1, 3, 4, 7}, 3},
		{[]int{1, 4, 2, 3, 3, 2, 4, 1}, []int{1, 2, 3, 4, 4, 3, 2, 1}, 4},
	} {
		log.Print("* ", c.nums, c.n)
		if !reflect.DeepEqual(c.rst, shuffle(c.nums, c.n)) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1556(t *testing.T) {
	for _, c := range []struct {
		rst string
		n   int
	}{
		{"987", 987},
		{"1.234", 1234},

		{"0", 0},
		{"51.040", 51040},
	} {
		log.Print("* ", c.n)
		if c.rst != thousandSeparator(c.n) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1662(t *testing.T) {
	for _, c := range []struct {
		rst          bool
		word1, word2 []string
	}{
		{true, []string{"ab", "c"}, []string{"a", "bc"}},
		{false, []string{"a", "cb"}, []string{"ab", "c"}},
		{true, []string{"abc", "d", "defg"}, []string{"abcddefg"}},
	} {
		log.Printf("* %q %q", c.word1, c.word2)
		if c.rst != arrayStringsAreEqual(c.word1, c.word2) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
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

func Test1725(t *testing.T) {
	for _, c := range []struct {
		rst        int
		rectangles [][]int
	}{
		{3, [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}},
		{3, [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}},
	} {
		log.Print("* ", c.rectangles)
		if c.rst != countGoodRectangle(c.rectangles) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

// 1813m Sentence Similarity III
func Test1813(t *testing.T) {
	log.Print("true ?= ", areSentencesSimilar("Hello Jane", "Hello my name is Jane"))
	log.Print("false ?= ", areSentencesSimilar("of", "A lot of words"))
	log.Print("true ?= ", areSentencesSimilar("Eating right now", "Eating"))
}

func Test1832(t *testing.T) {
	for _, c := range []struct {
		rst      bool
		sentence string
	}{
		{true, "thequickbrownfoxjumpsoverthelazydog"},
		{false, "pangram"},
	} {
		log.Printf("* %q", c.sentence)
		if c.rst != checkIfPangram(c.sentence) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test1880(t *testing.T) {
	for _, c := range []struct {
		rst                               bool
		firstWord, secondWord, targetWord string
	}{
		{true, "acb", "cba", "cdb"},
		{false, "aaa", "a", "aab"},
		{true, "aaa", "a", "aaaa"},
	} {
		log.Printf("* %q %q %q", c.firstWord, c.secondWord, c.targetWord)
		if c.rst != isSumEqual(c.firstWord, c.secondWord, c.targetWord) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

// 2185 Counting Words With a Given Prefix
func Test2185(t *testing.T) {
	log.Print("2 ?= ", prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
	log.Print("0 ?= ", prefixCount([]string{"leetcode", "win", "loops", "success"}, "code"))
}

func Test2196(t *testing.T) {
	for _, c := range []struct {
		rst []string
		s   string
	}{
		{[]string{"K1", "K2", "L1", "L2"}, "K1:L2"},
		{[]string{"A1", "B1", "C1", "D1", "E1", "F1"}, "A1:F1"},
	} {
		log.Print("* ", c.s)
		if !reflect.DeepEqual(c.rst, cellsInRange(c.s)) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test2269(t *testing.T) {
	for _, c := range []struct {
		rst, num, k int
	}{
		{2, 240, 2},
		{2, 430043, 2},
	} {
		log.Print("* ", c.num, c.k)
		if c.rst != divisorSubstrings(c.num, c.k) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
}

func Test2315(t *testing.T) {
	for _, c := range []struct {
		rst int
		s   string
	}{
		{0, "iamprogrammer"},
		{5, "yo|uar|e**|b|e***au|tifu|l"},
	} {
		log.Printf("* %q", c.s)
		if c.rst != countAsterisks(c.s) {
			t.FailNow()
		}
		log.Print(":: ", c.rst)
	}
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
