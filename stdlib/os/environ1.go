package main

import (
	"fmt"
	"os"
)

func main() {
	environ := os.Environ()

	l := len(environ)

	for x := 0; x < l; x++ {
		fmt.Printf("%d %s\n", x, environ[x])
	}
}
