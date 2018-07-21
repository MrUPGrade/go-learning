package main

import "fmt"

type MyPrinter func(string) ()

func printStuff2(message string, printer MyPrinter) {
	printer(message)
}

func myPrinter2(text string) {
	fmt.Println(text)
}

func main() {
	printStuff2("This is a message", myPrinter2)
}
