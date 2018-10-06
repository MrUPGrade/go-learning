package main

import "fmt"

func main() {
	type Person struct {
		First string
		Last  string
		Age   int
	}

	p := Person{
		First: "Jan",
		Last:  "Kowalski",
		Age:   123,
	}

	pp := &p

	fmt.Println(p)
	fmt.Println(pp)
	fmt.Println(*pp)
}
