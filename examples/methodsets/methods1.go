package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) Print() {
	fmt.Println("x: ", p.x, "y: ", p.y)
}

func main() {
	point := Point{1, 3}
	point.Print()
}
