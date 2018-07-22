package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Compiler:", runtime.Compiler)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
	fmt.Println("CPU num:", runtime.NumCPU())
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Println("Mem:", memStats)
	fmt.Println("GoRoutines num:", runtime.NumGoroutine())
}
