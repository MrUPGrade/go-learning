package main

import "fmt"

func main() {
	for i, r := range "Witaj, 世界" {
		fmt.Printf("%2d\t%q\t%d\n", i, r, r)
	}
}
