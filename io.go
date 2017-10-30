package gopher

import (
	"bufio"
	"io"
)

func ReaderToTokens(rdr io.Reader) []string {
	Tokens := []string{}

	scr := bufio.NewScanner(rdr)
	tkn := ""
	for scr.Scan() {
		line := scr.Text()
		if line == "%" {
			Tokens = append(Tokens, tkn)
			tkn = ""
		} else {
			if tkn != "" {
				tkn += "\n"
			}
			tkn += line
		}
	}

	return Tokens
}
