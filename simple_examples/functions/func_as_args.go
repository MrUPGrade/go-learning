package main

import "fmt"

func myPrint(a func(string) ()) {
	a("test")
}

func main() {
	myPrint(func(b string) {
		fmt.Println(b)
	})


	go func (t string) {
		fmt.Println(t)
	}("Test 2")
}
