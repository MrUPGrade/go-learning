package main

import "fmt"

func main() {
	m1 := new(map[string]int)
	m2 := &map[string]int{}

	// This still has to be done
	*m1 = make(map[string]int)
	(*m1)["a"] = 1

	(*m2)["a"] = 1

	fmt.Printf("%T %v %v \n", m1, m1, *m1)
	fmt.Printf("%T %v %v \n", m2, m2, *m2)
}
