package main

import "fmt"

func main() {
	type F struct {
		x, y int
		name string
	}

	var a F
	fmt.Println(a)
}
