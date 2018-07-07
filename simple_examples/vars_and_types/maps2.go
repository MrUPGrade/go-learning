package main

import "fmt"

func main() {
	var m map[string]string

	fmt.Println(m)

	// Maps have to be initialized!
	m["test"] = "value"
	fmt.Println(m)
}
