package main

import (
	"fmt"
)

func immediateDeferLoop() {
	for x := 0; x < 4; x++ {
		defer func() {
			fmt.Println("Defer:", x)
		}()
	}
	fmt.Println("End")
}

func main() {
	immediateDeferLoop()
}
