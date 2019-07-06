package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		sig := <- sigs
		fmt.Println("Signal:", sig)
		done <- true
	}()

	res := <-done
	fmt.Println("Result:", res)
}