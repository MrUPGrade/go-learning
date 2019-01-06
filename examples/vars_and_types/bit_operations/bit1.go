package main

import "fmt"

func main() {
	a := 1
	b := 8
	fmt.Printf("%.5b %.5b %.5b", a, b, a | b)
}
