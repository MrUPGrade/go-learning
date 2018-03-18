package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, value := range os.Args {
		fmt.Printf("idx: %v, arg value: %v\n", idx, value)
	}
}
