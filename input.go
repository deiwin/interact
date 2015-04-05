package interact

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errCanceled = errors.New("Command aborted")
)

// InputCheck specifies the function signature for an input check
type InputCheck func(string) error

// GetInputAndRetry asks the user for input and performs the list of added checks
// on the provided input. If any of the checks fail to pass the error will be
// displayed to the user and they will then be asked if they want to try again.
// If the user does not want to retry the program will return an error.
func (a Actor) GetInputAndRetry(message string, checks ...InputCheck) (string, error) {
	for {
		input, err := a.GetInput(message, checks...)
		if err != nil {
			if err = a.confirmRetry(err); err != nil {
				return "", err
			}
			continue
		}
		return input, nil
	}
}

// GetInputWithDefaultAndRetry works exactly like GetInputAndRetry, but also has
// a default option which will be used instead if the user simply presses enter.
func (a Actor) GetInputWithDefaultAndRetry(message, fallback string, checks ...InputCheck) (string, error) {
	for {
		input, err := a.GetInputWithDefault(message, fallback, checks...)
		if err != nil {
			if err = a.confirmRetry(err); err != nil {
				return "", err
			}
			continue
		}
		return input, nil
	}
}

// GetInput asks the user for input and performs the list of added checks on the
// provided input. If any of the checks fail, the error will be returned.
func (a Actor) GetInput(message string, checks ...InputCheck) (string, error) {
	input, err := a.getInput(message + ": ")
	if err != nil {
		return "", err
	}
	err = runChecks(input, checks...)
	if err != nil {
		return "", err
	}
	return input, nil
}

// GetInputWithDefault works exactly like GetInput, but also has a default option
// which will be used instead if the user simply presses enter.
func (a Actor) GetInputWithDefault(message, fallback string, checks ...InputCheck) (string, error) {
	input, err := a.getInput(fmt.Sprintf("%s: (%s) ", message, fallback))
	if err != nil {
		return "", err
	}
	if input == "" {
		return fallback, nil
	}
	err = runChecks(input, checks...)
	if err != nil {
		return "", err
	}
	return input, nil
}

func (a Actor) confirmRetry(err error) error {
	retryMessage := fmt.Sprintf("%v\nDo you want to try again?", err)
	confirmed, err := a.Confirm(retryMessage, ConfirmDefaultToNo)
	if err != nil {
		return err
	} else if !confirmed {
		return errCanceled
	}
	return nil
}

func (a Actor) getInput(prompt string) (string, error) {
	fmt.Fprint(a.w, prompt)
	line, err := a.rd.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func runChecks(input string, checks ...InputCheck) error {
	for _, check := range checks {
		err := check(input)
		if err != nil {
			return err
		}
	}
	return nil
}
