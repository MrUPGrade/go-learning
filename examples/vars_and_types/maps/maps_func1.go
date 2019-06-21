package main

import "fmt"

func modifyMap1(m map[string]string) map[string]string {
	m["test"] = "value"
	return m
}

func main() {
	m1 := map[string]string {
		"outer": "value",
	}

	m2 := modifyMap1(m1)

	fmt.Println(m1)
	fmt.Println(m2)
}
