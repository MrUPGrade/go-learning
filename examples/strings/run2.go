package main

import "fmt"

func main() {
	s := "Sh1t ju$t g0t rƸa| ⫗"
	for _, r := range s {
		fmt.Printf("char: %c rune: %b \n", r, r)
	}
}
