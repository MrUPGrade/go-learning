package main

import "fmt"

func main() {
	fmt.Println("Program starts")

	c := make(chan int)
	go func() {
		panic("Something happened")
		c <- 1
	}()
	<-c
	fmt.Println("Program ends")
}
