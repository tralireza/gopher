package gopher

import "log"

func init() {
	log.SetFlags(0)
}

type ByteCounter int

func (o *ByteCounter) Write(p []byte) (int, error) {
	*o += ByteCounter(len(p))
	return len(p), nil
}
