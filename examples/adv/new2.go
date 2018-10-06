package main

import "fmt"

func main() {
	type User struct {
		name  string
		age   int
		score float32
	}

	p1 := new([]User)
	p2 := make([]User, 0)

	fmt.Println(p1)
	for _, x := range *p1 {
		fmt.Println(x)
	}

	fmt.Println(p2)
	for _, x := range p2 {
		fmt.Println(x)
	}
}
