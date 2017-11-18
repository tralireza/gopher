package gopher

import (
	"log"
	"strings"
)

// 71m Simplify Path
func simplifyPath(path string) string {
	P := strings.Split(path, "/")

	log.Printf(" -> %q", P)

	Q := []string{}
	for _, p := range P {
		switch p {
		case "":
		case ".":
		case "..":
			if len(Q) > 0 {
				Q = Q[:len(Q)-1]
			}
		default:
			Q = append(Q, p)
		}
	}

	log.Printf(" -> %q", Q)

	return "/" + strings.Join(Q, "/")
}
