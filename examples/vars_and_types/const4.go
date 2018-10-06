package main

import "fmt"

const (
	count1 = iota
	count2 = iota
	count3 = iota
)

func main() {
	fmt.Println(count1, count2, count3)
}
