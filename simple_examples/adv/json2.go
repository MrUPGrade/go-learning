package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type t1 struct {
		Name  string `json:"JSONName"`
		value string `json:"JSONValue"`
		Age   int32
	}

	v1 := t1{
		Name:  "some name",
		value: "some value",
		Age:   32,
	}

	res, _ := json.MarshalIndent(v1, "", "  ")
	fmt.Println(res)
	fmt.Println(string(res))
}
