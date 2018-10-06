package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please pass in a file input")
		os.Exit(1)
	}

	dir, _ := os.Getwd()
	fmt.Println(">", dir)

	file, err := os.OpenFile(args[1], os.O_RDONLY, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("->", scanner.Text())
	}
}
