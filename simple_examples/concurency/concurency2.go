package main

import (
	"fmt"
)

func main() {
	isDone := make(chan bool)

	go func() {
		fmt.Println("go println")
		isDone <- true
	}()

	fmt.Println("println")
	<-isDone
}
