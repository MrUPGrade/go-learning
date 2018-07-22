package main

import "fmt"

func main() {
	m := map[string]string{
		"one":   "test1",
		"two":   "test2",
		"three": "test3",
	}

	fmt.Println(m)
	fmt.Println(m["one"])


	if value, exist := m["two"]; exist {
		fmt.Println("Two: ", value)
	}

	if value, exist := m["Four"]; exist {
		fmt.Println("Four: ", value)
	}
}
