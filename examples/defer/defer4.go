package main

import (
	"fmt"
)

func main() {
	fmt.Println("start program")
	if true {
		fmt.Println("start if")
		defer func() {
			fmt.Println("defer func")
		}()
		fmt.Println("end if")
	}
	fmt.Println("end program")
}
