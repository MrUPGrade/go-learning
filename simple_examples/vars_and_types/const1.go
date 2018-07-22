package main

import "fmt"

func main() {
	const name = "Bob"

	const (
		os   = "Linux"
		size = 123
	)

	fmt.Println(name)
	fmt.Println(size)
	fmt.Println(os)
}
