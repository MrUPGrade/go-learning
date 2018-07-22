package main

import "fmt"

func main() {
	type Human struct {
		name string
		age  int
	}

	h1 := Human{"Stefan", 123}

	fmt.Println(h1)
	fmt.Println(h1.name, h1.age)
}
