package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["first"] = "Jan"
	m["last"] = "Kowalski"

	fmt.Println(m)
}
