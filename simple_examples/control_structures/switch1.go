package main

import "fmt"

func main() {
	a := 2

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}
