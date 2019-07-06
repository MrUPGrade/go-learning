package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	var funcs []    func()

	for _, x := range numbers {
		funcs = append(funcs, func() {
			fmt.Println(x)
		})
	}

	for _, f := range funcs {
		f()
	}

}
