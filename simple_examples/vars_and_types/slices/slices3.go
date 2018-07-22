package main

import "fmt"

func main() {
	s := make([]int, 10)
	for x := 0; x < 10; x++ {
		s[x] = x
	}

	fmt.Println(s)
	fmt.Println(s[3:5])
	fmt.Println(s[:5])

	//Its not python, slices will raise errors
	//fmt.Println(s[10:20])
}
