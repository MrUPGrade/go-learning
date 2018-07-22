package main

import "fmt"

func main() {
	var arr2d [3][3]int

	fmt.Println(arr2d)

	for _, x := range arr2d {
		fmt.Println(x)
	}
}
