package main

import (
	"strings"
	"fmt"
)

func main() {
	input := "this is a log string with lots of spaces"
	fields := strings.Fields(input)

	fmt.Println("len:", len(fields))
	for idx, val := range fields {
		fmt.Println("idx: ", idx, "val: ", val)
	}
}
