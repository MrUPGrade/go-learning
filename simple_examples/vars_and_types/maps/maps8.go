package main

import "fmt"

func main() {
	m := map[string]string{
		"one":   "test1",
		"two":   "test2",
		"three": "test3",
	}

	fmt.Println(m)

	//Watch out for this
	x := m["o"]
	fmt.Println(x)
}
