package main

import "fmt"

func main() {
	var a [3]int
	a[1] = 123

	// Do not build :)
	fmt.Println(a[4])
}
