package main

import "fmt"

func printer(msg ...string) {
	for _, x := range msg {
		fmt.Println(x)
	}
}

func main() {
	printer("one", "two", "three")
}
