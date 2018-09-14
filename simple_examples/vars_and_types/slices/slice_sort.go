package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	s := make([]int, 10)
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}

	fmt.Println("S:", s)

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	fmt.Println("S:", s)
}
