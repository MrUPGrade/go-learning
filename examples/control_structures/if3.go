package main

import (
	"fmt"
)

func main() {
	var v string

	fmt.Print("Please input text: ")

	_, err := fmt.Scan(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println("Value: ", v)

	if len(v) < 4 {
		fmt.Println("Small")
	} else {
		fmt.Println("Big")
	}
}
