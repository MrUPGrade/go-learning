package main

import (
	"fmt"
	"os"
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
	os.Exit(-1)
	fmt.Println("end program")
}
