package main

import (
	"fmt"
)

func immediateDeferLoop2() {
	for x := 0; x < 4; x++ {
		defer func(n int) {
			fmt.Println("Defer:", n)
		}(x)
	}
	fmt.Println("End")
}

func main() {
	immediateDeferLoop2()
}
