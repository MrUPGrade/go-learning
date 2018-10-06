package main

import "fmt"

func main() {
	s1 := make([]int, 0, 3)

	for x := 1; x < 32; x++ {
		s1 = append(s1, x)
		fmt.Println("item:", x, "Size:", len(s1), "cap:", cap(s1))
	}
}
