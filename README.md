# Interact
A Golang utility belt for interacting with the user over a CLI

[![Build Status](https://travis-ci.org/deiwin/interact.svg?branch=master)](https://travis-ci.org/deiwin/interact)
[![Coverage](http://gocover.io/_badge/github.com/deiwin/interact?0)](http://gocover.io/github.com/deiwin/interact)
[![GoDoc](https://godoc.org/github.com/deiwin/interact?status.svg)](https://godoc.org/github.com/deiwin/interact)

## Example interaction

Code like this:
```go
actor := interact.NewActor(os.Stdin, os.Stdout)

message := "Please enter something that's not empty"
notEmpty, err := actor.Prompt(message, checkNotEmpty)
if err != nil {
  log.Fatal(err)
}
message = "Please enter a positive number"
n1, err := actor.PromptAndRetry(message, checkNotEmpty, checkIsAPositiveNumber)
if err != nil {
  log.Fatal(err)
}
message = "Please enter another positive number"
n2, err := actor.PromptOptionalAndRetry(message, "7", checkNotEmpty, checkIsAPositiveNumber)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Thanks! (%s, %s, %s)\n", notEmpty, n1, n2)
```

Can create an interaction like this:

![asciicast](https://cloud.githubusercontent.com/assets/2261897/7066876/6194ec42-decf-11e4-823a-019f921f52a1.gif)

For a more comprehensive example see the [example test](https://github.com/deiwin/interact/blob/master/example_test.go).
