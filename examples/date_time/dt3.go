package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	t2 := t.Format(time.RFC3339)
	fmt.Println(t2)
}
