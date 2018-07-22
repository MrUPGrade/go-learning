package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("go println")
	fmt.Println("println")

	time.Sleep(time.Second)
}
