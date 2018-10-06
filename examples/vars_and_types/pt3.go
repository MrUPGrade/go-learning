package main

import "fmt"

func main() {
	msg := "This is a hello world!"
	pMsg := &msg

	*pMsg = "nope"

	fmt.Println("  msg:", msg)
	fmt.Println("*pMsg:", *pMsg)
}
