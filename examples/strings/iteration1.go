package main

import "fmt"

func main() {
	s := "this is a string"

	for i := 0; i < len(s); i++ {
		fmt.Printf("char: %c (%08b) [%02d] \n", s[i], s[i], i)
	}
}
