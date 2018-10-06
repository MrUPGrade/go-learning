package main

import "fmt"

func main() {
	m := map[string]string{
		"one": "test1",
		"two": "test2",
	}

	fmt.Println(m)

	m["test"] = "value"
	fmt.Println(m)
}
