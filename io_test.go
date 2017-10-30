package gopher

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestIO(t *testing.T) {
	f, err := os.Open("fortunes.txt")
	if err != nil {
		t.Fatal(err)
	}

	Tokens := ReaderToTokens(f)

	log.Printf("%d :: %q", len(Tokens), Tokens)
	log.Print("-> ", Tokens[rand.Intn(len(Tokens))])

	for _, input := range []string{"-\n%\n%\n", `% LINE1 %
%
LINE2 %
%
% LINE3
%
%
`} {
		Tokens := ReaderToTokens(strings.NewReader(input))
		log.Printf("+++ %q", Tokens)
	}
}
