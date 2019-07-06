package main

import "fmt"

func main() {
	fmt.Println("Program start")

	alwaysRecover := func() {
		recVal := recover()
		if recVal != nil {
			fmt.Printf("Recovered from error: '%v' of type: '%T'\n", recVal, recVal)
		}
	}

	funcWithPanic := func() {
		fmt.Println("Inner func start")
		defer alwaysRecover()
		// Comment out panic to see how recover will react
		panic("test123")
		fmt.Println("Inner func start")
	}

	funcWithPanic()
	fmt.Println("Program ends")
}
