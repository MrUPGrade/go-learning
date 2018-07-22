package main

import (
	"time"
	"fmt"
)

func main() {
	numbers := make(chan int, 10)

	go func() {
		for x := 0; x < 100; x += 10 {
			for y := 0; y <= 9; y++ {

				numbers <- y + x
			}
			time.Sleep(time.Millisecond*2)
		}
		close(numbers)
	}()

	var nums [512]int
	idx := 0
	for {
		select {
		case i, ok := <-numbers:
			nums[idx] = i
			idx++
			fmt.Println(i)
			if !ok {
				fmt.Println("Not ok")
				goto endfor
			}

		default:
			fmt.Println("waiting...")
			time.Sleep(time.Millisecond * 10)
		}
	}

endfor:
	fmt.Println(nums)

}
