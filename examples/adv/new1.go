package main

import "fmt"

func main() {
	type User struct {
		name  string
		age   int
		score float32
	}

	p := new(User)
	p.age = 3
	fmt.Println(*p)

}
