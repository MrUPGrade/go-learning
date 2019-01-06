package main

import "fmt"

func main() {
	type MyInt int

	var m1 MyInt
	var m2 MyInt

	m1, m2 = 1, 2

	fmt.Println(m1, m2)
}
