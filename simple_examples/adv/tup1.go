package main

import "fmt"

func ret3() (int, int, int) {
	return 1, 2, 3
}
func main() {
	fmt.Println(ret3())
}
