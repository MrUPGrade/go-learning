package main

import "fmt"

func main() {
	a := 1

	switch a {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}
