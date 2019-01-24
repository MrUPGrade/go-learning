package main

import "fmt"

func main() {
	a, b := 1, "test"
	myPrint := func(format string, a ...interface{}) {
		fmt.Printf(format, a...)
	}

	myPrint("%v %v\n", b, a)
	myPrint("fuck\n")
}
