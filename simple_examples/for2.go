package main

import "fmt"

func main() {
	// For each
	for idx, x := range []int{10, 20, 30, 40} {
		fmt.Printf("idx: %v, value: %v\n", idx, x)
	}
}
