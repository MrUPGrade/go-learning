package main

import "fmt"

func panicFunc() {
	fmt.Println("PANIC!!!")
	panic("Damn!")
	fmt.Println("After PANIC")
}

func flow() {
	fmt.Println("Flow starts")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()

	panicFunc()
	fmt.Println("Flow ends")
}

func main() {
	fmt.Println("Program starts")
	flow()
	fmt.Println("Program ends")

}
