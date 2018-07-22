package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	users := []User{
		{name: "bob", age: 123},
		{name: "tom", age: 33},
	}

	for i, u := range users {
		fmt.Println("user: ", u, "index: ", i)
	}
}
