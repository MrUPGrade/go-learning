package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T: %[1]v\n", arr)
}
