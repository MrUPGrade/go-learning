package main

import "fmt"

func main() {

	var i int

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}
