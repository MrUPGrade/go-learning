package main

import (
	"fmt"
)

func deferLoop() {
	for x := 0; x < 4; x++ {
		defer fmt.Println("Defer number", x)
	}
}

func main() {
	deferLoop()
}
