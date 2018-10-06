package main

import "fmt"

type MyPrinter3 func(string) ()

func printStuff3(message string, printer MyPrinter3) {
	printer(message)
}

func MyPrinterFactory(suffix string) MyPrinter3 {
	return func(msg string) {
		fmt.Println(msg + suffix)
	}

}

func main() {
	f := MyPrinterFactory(" SUFFIX")
	printStuff3("This is a message", f)
}
