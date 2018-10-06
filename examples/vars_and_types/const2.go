package main

import "fmt"

func main() {
	const (
		num2 = 3.12
	)

	var f1 float32
	var f2 float64

	// You can`t do that with typed values
	f1 = num2
	f2 = num2

	fmt.Println(f1)
	fmt.Println(f2)
}
