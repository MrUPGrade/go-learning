package main

import "fmt"

func greet(name, greeting string) {
	fmt.Printf("Hi %v! %v\n", name, greeting)
}

func main() {
	greet("Bob", "How are you doing?")
}
