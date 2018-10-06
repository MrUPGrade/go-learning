package main

import (
	"errors"
)

func main() {
	err := errors.New("This is a error for the typical panic mode on!")
	panic(err)
}
