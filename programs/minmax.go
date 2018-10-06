package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide some numbers")
		os.Exit(1)
	}

	var numbers []float64

	for i := 1; i < len(args); i ++ {
		num, err := strconv.ParseFloat(args[i], 64)

		if err != nil {
			errorMsg := "Argument '" + args[i] + "' is not a number"
			io.WriteString(os.Stderr, errorMsg)
			continue
		}

		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		io.WriteString(os.Stderr, "No numbers were provided")
		os.Exit(1)
	}

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	fmt.Println("Min:", min, "Max:", max)
}
