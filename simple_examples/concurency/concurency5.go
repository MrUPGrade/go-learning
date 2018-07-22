package main

import (
	"time"
	"fmt"
)

func main() {
	numbers := make(chan int, 32)
	waitFor := make(chan int)

	go func() {
		for x := 0; x < 100; x += 10 {
			for y := 0; y <= 9; y++ {

				numbers <- y + x
			}
			time.Sleep(time.Millisecond*5)
		}
		close(numbers)
	}()

	var nums []int

	go func() {
		for {
			select {
			case i, ok := <-numbers:
				if !ok {
					fmt.Println("Not ok")
					goto l
				}
				nums = append(nums, i)
				fmt.Println(i)

			default:
				fmt.Println("waiting...")
				time.Sleep(time.Millisecond * 10)
			}
		}
	l:
		waitFor <- 0
	}()

	<-waitFor
	fmt.Println(nums)
	fmt.Println("Len:", len(nums))

}
