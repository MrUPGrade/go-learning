package main

import "fmt"

func main() {
	m := map[string]string{}

	fmt.Println(m)

	m["first"] = "Jan"
	m["last"] = "Kowalski"

	fmt.Println(m)
}
