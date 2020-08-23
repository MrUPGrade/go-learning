package main

import (
	"fmt"
)

type (
	Point2D struct {
		x int
		y int
	}

	Printable interface {
		Print() string
	}

	Readable interface {
		GetX() int
		GetY() int
	}

	Editable interface {
		SetX(int)
		SetY(int)
	}

	Combined interface {
		Printable
		Readable
		Editable
	}
)

func NewPoint2D(x int, y int) Point2D {
	return Point2D{x: x, y: y}
}

func (p Point2D) Print() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}

func (p Point2D) GetX() int {
	return p.x
}

func (p Point2D) GetY() int {
	return p.y
}

func (p *Point2D) SetX(x int) {
	p.x = x
}

func (p *Point2D) SetY(y int) {
	p.y = y
}

func main() {
	var p1Printable Printable
	var p1Readable Readable
	var p1Editable Editable
	var p2Combined Combined

	p1 := NewPoint2D(1, 2)
	p1Printable = p1
	fmt.Println(p1Printable.Print())

	p1Readable = p1
	fmt.Println("Point2d:", p1Readable.GetX(), p1Readable.GetY())

	p1Editable = &p1
	p1Editable.SetX(666)

	fmt.Println("After change", p1Printable.Print())
	fmt.Println("After change - p1", p1.Print())

	p2 := NewPoint2D(9, 9)
	fmt.Println("p2: ", p2.Print())
	p2Combined = &p2

	p2Combined.SetX(666)
	p2Combined.SetY(666)
	fmt.Println("compare:", p2Combined.Print(), p2.Print())
}
