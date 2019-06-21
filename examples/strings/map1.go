package main

import "strings"

func main() {
	s := strings.Map(func(r rune) rune {
		r++
		return r
	}, "test123")

	print(s + "\n")
}
