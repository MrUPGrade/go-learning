package main

import "fmt"

type Figure interface {
	Area() int
	Scale(int)
}

type Square struct {
	a int
}

func (s Square) Area() int {
	return s.a * s.a
}

func (s *Square) Scale(scaleValue int) {
	s.a *= scaleValue
}

func main() {
	var f Figure

	s := Square{a: 5}
	s.Area()

	f = &s
	f.Scale(2)
	fmt.Println(f)
}
