package main

import "fmt"

func main() {
	v := true

	if msg := "true"; v {
		fmt.Println(msg)
	}

	//Not in this scope
	//fmt.Println(msg)
}
