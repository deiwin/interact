package interact_test

import (
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
