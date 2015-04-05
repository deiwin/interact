package interact_test

import (
	"io"
	"strings"

	"github.com/deiwin/interact"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"testing"
)

func TestInteract(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interact Suite")
}

var (
	actor     interact.Actor
	userInput = "\n"
	output    *gbytes.Buffer
)

var _ = BeforeEach(func() {
	output = gbytes.NewBuffer()
})

var _ = JustBeforeEach(func() {
	actor = interact.NewActor(strings.NewReader(userInput), output)
})

func NewTestBuffer(r io.Reader, w io.Writer) *TestBuffer {
	return &TestBuffer{
		r:         r,
		w:         w,
		userInput: make(chan string, 10),
	}
}

// TestBuffer is a hack that makes both the test and the test log easy to read
type TestBuffer struct {
	r         io.Reader
	w         io.Writer
	userInput chan string
}

func (b *TestBuffer) Read(p []byte) (int, error) {
	n, err := b.r.Read(p)
	if err != nil {
		return 0, err
	}
	s := strings.TrimSuffix(string(p[:n]), "\n")
	for _, line := range strings.Split(s, "\n") {
		b.userInput <- line
	}
	return n, err
}

func (b *TestBuffer) Write(p []byte) (int, error) {
	var prefix string
	select {
	case i := <-b.userInput:
		prefix = i + "\n"
	default:
	}
	return b.w.Write(append([]byte(prefix), p...))
}
