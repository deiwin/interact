# Interact
A Golang utility belt for interacting with the user over a CLI

[![Build Status](https://travis-ci.org/deiwin/interact.svg?branch=master)](https://travis-ci.org/deiwin/interact)
[![Coverage](http://gocover.io/_badge/github.com/deiwin/interact?0)](http://gocover.io/github.com/deiwin/interact)
[![GoDoc](https://godoc.org/github.com/deiwin/interact?status.svg)](https://godoc.org/github.com/deiwin/interact)

## Example interaction

Code like this:
```
actor := interact.NewActor(os.Stdin, os.Stdout)

message := "Please enter something that's not empty"
notEmpty, err := actor.GetInput(message, checkNotEmpty)
if err != nil {
  log.Fatal(err)
}
message = "Please enter a positive number"
n1, err := actor.GetInputAndRetry(message, checkNotEmpty, checkIsAPositiveNumber)
if err != nil {
  log.Fatal(err)
}
message = "Please enter another positive number"
n2, err := actor.GetInputWithDefaultAndRetry(message, "7", checkNotEmpty, checkIsAPositiveNumber)
if err != nil {
  log.Fatal(err)
}
log.Printf("Thanks! (%s, %s, %s)\n", notEmpty, n1, n2)
```

Can create an interaction like this:
```
Please enter something that's not empty: hello
Please enter a positive number: -2
The number can not be negative!
Do you want to try again? [y/N]: y
Please enter a positive number: 5
Please enter another positive number: (7) -2
The number can not be negative!
Do you want to try again? [y/N]: y
Please enter another positive number: (7)
Thanks! (hello, 5, 7)
```

For a more comprehensive example see the [example test](https://github.com/deiwin/interact/blob/master/example_test.go).
