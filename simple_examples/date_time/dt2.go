package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	fmt.Printf("Time: %02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
}
