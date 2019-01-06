package main

import "fmt"

func main() {
	arr := [...]int{10: 1}
	fmt.Printf("%T: %[1]v", arr)
}
