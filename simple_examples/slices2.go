package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[3] = 12

	//Can`t extend slice, need append
	//s[11] = 3
	fmt.Println(s)
}
