package interact_test

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

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
	var b = NewTestBuffer(&userInput, os.Stdout)
	// Normally you would initiate the actor with os.Stdin and os.Stdout, but to
	// make this example test work nice we need to use something different
	actor := interact.NewActor(b, b)

	// A simple prompt for non-empty input
	userInput.WriteString("hello\n") // To keep the test simple we have to provide user inputs up front
	if result, err := actor.Prompt("Please enter something that's not empty", checkNotEmpty); err != nil {
		log.Fatal(err)
	} else if result != "hello" {
		log.Fatalf("Expected 'hello', got '%s'", result)
	}

	// A more complex example with the user retrying
	userInput.WriteString("-2\ny\n5\n")
	if result, err := actor.PromptAndRetry("Please enter a positive number", checkNotEmpty, checkIsAPositiveNumber); err != nil {
		log.Fatal(err)
	} else if result != "5" {
		log.Fatalf("Expected '5', got '%s'", result)
	}

	// An example with the user retrying and then opting to use the default value
	userInput.WriteString("-2\ny\n\n")
	if result, err := actor.PromptOptionalAndRetry("Please enter another positive number", "7", checkNotEmpty, checkIsAPositiveNumber); err != nil {
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
