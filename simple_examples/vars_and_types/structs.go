package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p := Person{
		First: "Jan",
		Last:  "Kowalski",
		Age:   123,
	}

	pp := &p

	fmt.Println(p)
	fmt.Println(pp)
}
