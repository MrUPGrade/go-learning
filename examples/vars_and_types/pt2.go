package main

import "fmt"

func main() {
	msg := "This is a hello world!"
	pMsg := &msg

	fmt.Println("  msg:", msg)
	fmt.Println(" pMsg:", pMsg)
	fmt.Println("*pMsg:", *pMsg)
}
