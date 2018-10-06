package main

import "fmt"

func rectangle2(a, b int) (area int, circumference int) {
	area = a * b
	circumference = 2*a + 2*b

	return
}

func main() {
	x, y := rectangle2(2, 3)
	fmt.Println("Area:", x, "Circumference:", y)

	area, _ := rectangle2(2, 2)
	fmt.Println("Area only:", area)
}
