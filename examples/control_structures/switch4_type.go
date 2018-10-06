package main

import "fmt"

func printType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	default:
		fmt.Println("Unsupported")
	}
}

func main() {
	i := 1
	f := 1.0
	s := "Test"

	printType(i)
	printType(f)
	printType(s)
}
