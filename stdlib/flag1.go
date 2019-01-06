package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "", "enter your name")

func main() {
	flag.Parse()
	if *name != "" {
		fmt.Println("Your name is:", *name)
	} else {
		fmt.Println("Name not provided")
	}
}
