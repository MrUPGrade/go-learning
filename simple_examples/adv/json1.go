package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v1 := map[string]string{
		"test":    "value",
		"another": "value",
	}
	res, _ := json.MarshalIndent(v1, "", "  ")
	fmt.Println(res)
	fmt.Println(string(res))
}
