package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["first"] = "Jan"
	m["last"] = "Kowalski"

	fmt.Println(m)
}
