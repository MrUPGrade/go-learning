package main

import "fmt"

func main() {
	arr := [3]int{0, 1, 2}
	s1 := arr[:]
	arr[1] = 99

	fmt.Printf("%T: %[1]v\n", arr)
	fmt.Printf("%T: %[1]v\n", s1)

	s1[1] = 100
	s1 = append(s1, 3, 4)
	s1[2] = 102

	fmt.Printf("%T: %[1]v\n", arr)
	fmt.Printf("%T: %[1]v\n", s1)
}
