package main

import "fmt"

const nameLen = 5

func main() {
	fmt.Print("Input your name: ")
	var name string
	_, err := fmt.Scan(&name)

	if err != nil {
		panic(err)
	}

	switch {
	case len(name) < nameLen:
		fmt.Println(name, "is a short name")
	case len(name) == nameLen:
		fmt.Println(name, "is 5 letters")
	case len(name) >= nameLen:
		fmt.Println(name, "is a long name")
	}
}
