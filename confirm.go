package interact

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// errNoOptionSelected is returned when the user has not selected either yes or no
	errNoOptionSelected = errors.New("Please select y/n!")
)

// ConfirmDefault specifies what an empty user input defaults to
type ConfirmDefault int

// Possible options for what an empty user input defaults to
const (
	ConfirmDefaultToYes ConfirmDefault = iota
	ConfirmDefaultToNo
	ConfirmNoDefault
)

// Confirm provides the message to the user and asks yes or no. If the user
// doesn't select either of the possible answers they will be prompted to answer
// again until they do
func (a Actor) Confirm(message string, def ConfirmDefault) (bool, error) {
	for {
		confirmed, err := a.confirmOnce(message, def)
		if err == errNoOptionSelected {
			fmt.Fprintln(a.w, err)
			continue
		}
		return confirmed, err
	}
}

func (a Actor) confirmOnce(message string, def ConfirmDefault) (bool, error) {
	var options string
	switch def {
	case ConfirmDefaultToYes:
		options = "[Y/n]"
	case ConfirmDefaultToNo:
		options = "[y/N]"
	case ConfirmNoDefault:
		options = "[y/n]"
	}
	fmt.Fprintf(a.w, "%s %s: ", message, options)

	line, err := a.rd.ReadString('\n')
	input := strings.TrimSpace(line)
	if err != nil {
		return false, err
	} else if input == "" {
		switch def {
		case ConfirmDefaultToYes:
			return true, nil
		case ConfirmDefaultToNo:
			return false, nil
		case ConfirmNoDefault:
			return false, errNoOptionSelected
		}
	}
	switch input {
	case "y":
		return true, nil
	case "n":
		return false, nil
	}
	return false, errNoOptionSelected
}
