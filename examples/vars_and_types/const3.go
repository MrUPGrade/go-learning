package main

import "fmt"

const n1 float32 = 3.12

func main() {
	var f1 float32
	var f2 float64

	f1 = n1
	// You can`t do that, because const is typed
	//f2 = n1

	fmt.Println(f1)
	fmt.Println(f2)
}
