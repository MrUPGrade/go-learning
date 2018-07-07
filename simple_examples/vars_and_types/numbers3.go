package main

import "fmt"

func main() {
	i := 2.3 // float64 !
	var j float64 = 4

	i += 3.2

	fmt.Printf("i: %v, j: %v, multi: %v\n", i, j, i*j)
}
