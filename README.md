# Interact
A Golang utility belt for interacting with the user over a CLI

[![Build Status](https://travis-ci.org/deiwin/interact.svg?branch=master)](https://travis-ci.org/deiwin/interact)
[![Coverage](http://gocover.io/_badge/github.com/deiwin/interact?0)](http://gocover.io/github.com/deiwin/interact)
[![GoDoc](https://godoc.org/github.com/deiwin/interact?status.svg)](https://godoc.org/github.com/deiwin/interact)

## Example interaction

To see how the following interaction can be created see the [example test](https://github.com/deiwin/interact/blob/master/example_test.go).

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
Thanks!
```
