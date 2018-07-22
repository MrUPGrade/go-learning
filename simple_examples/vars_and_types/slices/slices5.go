package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3

	s2 := s1[:]
	s3 := make([]int, 3)
	copy(s3, s1)

	s1[1] = 123

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
