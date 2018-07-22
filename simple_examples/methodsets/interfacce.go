package main

import "fmt"

type Printable interface {
	PrintData()
}

type Point2D struct {
	x int
	y int
}

func (p Point2D) PrintData() {
	fmt.Println("2D data:", p.x, p.y)
}

type Point3D struct {
	x, y, z int
}

func (p Point3D) PrintData() {
	fmt.Println("3D data:", p.x, p.y, p.z)
}

func pointPrinter(points ...Printable) {
	for _, v := range points {
		v.PrintData()
	}

}

func main() {
	p2d := Point2D{1, 1}
	p3d := Point3D{1, 1, 1}

	pointPrinter(p2d, p3d)
}
