package main

import (
	"fmt"
	"time"
)

func main() {
	nubmers := make(chan int, 32)

	go func() {
		for x := 0; x <= 1024; x ++ {

			nubmers <- x
		}
		nubmers <- -1
		time.Sleep(time.Microsecond*100000)
		// Use close
		//close(nubmers)
	}()

	for r := range nubmers {
		fmt.Println(r)
	}

}
