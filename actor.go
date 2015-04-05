// Package interact is a utility belt for interacting with the user over a CLI
package interact

import (
	"bufio"
	"io"
)

// An Actor provides methods to interact with the user
type Actor struct {
	rd *bufio.Reader
	w  io.Writer
}

// NewActor creates a new Actor instance with the specified io.Reader
func NewActor(rd io.Reader, w io.Writer) Actor {
	return Actor{bufio.NewReader(rd), w}
}
