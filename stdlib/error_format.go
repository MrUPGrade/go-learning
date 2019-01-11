package main

import (
	"fmt"
	"os"
)

func main() {
	err := fmt.Errorf("this is an error %s", "ERROR")
	_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
}
