package main

import "fmt"

func main() {
	s := "this is a string"

	for idx, val := range s {
		fmt.Printf("char: %c (%08b) [%02d] \n", val, val, idx)
	}
}
