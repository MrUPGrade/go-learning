package main

import "fmt"

func main() {
	a := 2

	switch a {
	case 1, 2, 3, 4:
		fmt.Println("one")
	case 10:
		fmt.Println("ten")
	default:
		fmt.Println("default")
	}
}
