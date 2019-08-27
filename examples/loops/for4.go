package main

import "fmt"

func main() {
	// For each
	for idx := range []int{10, 20, 30, 40} {
		fmt.Printf("index: %v\n", idx)
	}
}
