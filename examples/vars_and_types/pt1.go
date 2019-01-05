package main

import "fmt"

func main() {
	var msg = "This is a hello world!"
	var pMsg = &msg

	fmt.Println("  msg:", msg)
	fmt.Println(" pMsg:", pMsg)
	fmt.Println("*pMsg:", *pMsg)
}
