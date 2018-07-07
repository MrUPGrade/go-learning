package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")
	fmt.Printf("User: %v", user)
}
