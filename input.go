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

// PromptAndRetry asks the user for input and performs the list of added checks
// on the provided input. If any of the checks fail to pass the error will be
// displayed to the user and they will then be asked if they want to try again.
// If the user does not want to retry the program will return an error.
func (a Actor) PromptAndRetry(message string, checks ...InputCheck) (string, error) {
	for {
		input, err := a.Prompt(message, checks...)
		if err != nil {
			if err = a.confirmRetry(err); err != nil {
				return "", err
			}
			continue
		}
		return input, nil
	}
}

// PromptOptionalAndRetry works exactly like GetInputAndRetry, but also has
// a default option which will be used instead if the user simply presses enter.
func (a Actor) PromptOptionalAndRetry(message, defaultOption string, checks ...InputCheck) (string, error) {
	for {
		input, err := a.PromptOptional(message, defaultOption, checks...)
		if err != nil {
			if err = a.confirmRetry(err); err != nil {
				return "", err
			}
			continue
		}
		return input, nil
	}
}

// Prompt asks the user for input and performs the list of added checks on the
// provided input. If any of the checks fail, the error will be returned.
func (a Actor) Prompt(message string, checks ...InputCheck) (string, error) {
	input, err := a.prompt(message + ": ")
	if err != nil {
		return "", err
	}
	err = runChecks(input, checks...)
	if err != nil {
		return "", err
	}
	return input, nil
}

// PromptOptional works exactly like Prompt, but also has a default option
// which will be used instead if the user simply presses enter.
func (a Actor) PromptOptional(message, defaultOption string, checks ...InputCheck) (string, error) {
	input, err := a.prompt(fmt.Sprintf("%s: (%s) ", message, defaultOption))
	if err != nil {
		return "", err
	}
	if input == "" {
		return defaultOption, nil
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

func (a Actor) prompt(message string) (string, error) {
	fmt.Fprint(a.w, message)
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
