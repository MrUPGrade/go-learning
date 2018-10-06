package main

import "fmt"

func main() {
	s := []int{}
	fmt.Println(s)

	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
