package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		io.WriteString(os.Stderr, "Sorry, but you have to provide an argument")
	} else {
		for x := 1; x < len(args); x++ {
			line := args[x] + "\n"
			io.WriteString(os.Stdout, line)
		}
	}
}
