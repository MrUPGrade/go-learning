package main

import "fmt"

func printStuff(message string, printer func(string)) {
	printer(message)
}

func myPrinter(text string) {
	fmt.Println(text)
}

func main() {
	printStuff("This is a message", myPrinter)
}
