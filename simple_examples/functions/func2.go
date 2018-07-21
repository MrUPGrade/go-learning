package main

import "fmt"

func rectangle(a, b int) (int, int) {
	area := a * b
	circumference := 2*a + 2*b

	return area, circumference
}

func main() {
	x, y := rectangle(2, 3)
	fmt.Println("Area:", x, "Circumference:", y)

	area, _ := rectangle(2, 2)
	fmt.Println("Area only:", area)
}
