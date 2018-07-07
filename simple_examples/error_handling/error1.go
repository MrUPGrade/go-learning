package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Some error")
	fmt.Println(err)
}
