package main

import "fmt"

func main() {
	arr := [3]int{0, 1, 2}
	s1 := arr[:]
	arr[1] = 99

	fmt.Printf("%6T: %[1]v len: %[2]v\n", arr, len(arr))
	fmt.Printf("%6T: %[1]v len: %[2]v\n", s1, len(s1))

	s1[1] = 100

	// At this point, s1 points to new array
	s1 = append(s1, 3, 4)

	s1[2] = 102

	fmt.Printf("%6T: %[1]v len: %[2]v\n", arr, len(arr))
	fmt.Printf("%6T: %[1]v len: %[2]v\n", s1, len(s1))
}
