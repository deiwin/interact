package interact_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/deiwin/interact"
)

// First let's declare some simple input validators
var (
	checkNotEmpty = func(input string) error {
		// note that the inputs provided to these checks are already trimmed
		if input == "" {
			return errors.New("Input should not be empty!")
		}
		return nil
	}
	checkIsAPositiveNumber = func(input string) error {
		if n, err := strconv.Atoi(input); err != nil {
			return err
		} else if n < 0 {
			return errors.New("The number can not be negative!")
		}
		return nil
	}
)

func Example() {
	var userInput bytes.Buffer
	var b = &TestBuffer{
		r:         &userInput,
		w:         os.Stdout,
		userInput: make(chan string, 10),
	}
	actor := interact.NewActor(b, b)

	// A simple prompt for non-empty input
	userInput.WriteString("hello\n") // To keep the test simple we have to provide user inputs up front
	if result, err := actor.GetInput("Please enter something that's not empty", checkNotEmpty); err != nil {
		log.Fatal(err)
	} else if result != "hello" {
		log.Fatalf("Expected 'hello', got '%s'", result)
	}

	// A more complex example with the user retrying
	userInput.WriteString("-2\ny\n5\n")
	if result, err := actor.GetInputAndRetry("Please enter a positive number", checkNotEmpty, checkIsAPositiveNumber); err != nil {
		log.Fatal(err)
	} else if result != "5" {
		log.Fatalf("Expected '5', got '%s'", result)
	}

	// An example with the user retrying and then opting to use the default value
	userInput.WriteString("-2\ny\n\n")
	if result, err := actor.GetInputWithDefaultAndRetry("Please enter another positive number", "7", checkNotEmpty, checkIsAPositiveNumber); err != nil {
		log.Fatal(err)
	} else if result != "7" {
		log.Fatalf("Expected '7', got '%s'", result)
	}

	// This will force the last user input to be printed as well
	fmt.Fprint(b, "")

	// Output:
	// Please enter something that's not empty: hello
	// Please enter a positive number: -2
	// The number can not be negative!
	// Do you want to try again? [y/N]: y
	// Please enter a positive number: 5
	// Please enter another positive number: (7) -2
	// The number can not be negative!
	// Do you want to try again? [y/N]: y
	// Please enter another positive number: (7)
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
