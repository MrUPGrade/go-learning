package main

import "fmt"

type Human struct {
	name string
	age  int
}

func main() {
	h1 := Human{"Stefan", 123}

	fmt.Println(h1)
	fmt.Println(h1.name, h1.age)
}
