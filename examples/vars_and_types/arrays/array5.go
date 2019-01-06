package main

import "fmt"

func main() {
	const (
		PL = iota
		EN
		GB
	)

	a1 := [4]string{PL: "PL", EN: "EN", GB: "GB", 3: "unknown"}
	for idx, item := range a1 {
		fmt.Printf("country: %10v code: %d\n", item, idx)
	}
}
