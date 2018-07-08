package main

import "fmt"

const name  = "Bob"

const (
	os = "Linux"
	size = 123
)

func main() {
	fmt.Println(name)
	fmt.Println(size)
	fmt.Println(os)
}
