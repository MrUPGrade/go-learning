package main

import "fmt"

func main() {
	var msg string = "This is a hello world!"
	var pMsg *string = &msg

	fmt.Println("  msg:", msg)
	fmt.Println(" pMsg:", pMsg)
	fmt.Println("*pMsg:", *pMsg)
}
