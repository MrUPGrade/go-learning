package main

import "fmt"

type Point2 struct {
	x int
	y int
}

func (p Point2) Print() {
	fmt.Println("x: ", p.x, "y: ", p.y)
}

func (p Point2) Add(x, y int) {
	p.x += x
	p.y += y
}

func (p *Point2) PAdd(x, y int) {
	p.x += x
	p.y += y
}

func main() {
	point1 := Point2{1, 3}
	point2 := Point2{1, 3}

	point1.Add(2, 2)
	point2.PAdd(2, 2)

	point1.Print()
	point2.Print()
}
