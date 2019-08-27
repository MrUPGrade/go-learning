package main

import "fmt"

type Value struct{ value int }

// Try removing the '*'
func (v *Value) Print() {
	fmt.Printf("Value: '%v', type: '%T'\n", v.value, v)
}

// Try removing the '*'
func (v *Value) ChangeValue(newvalue int) {
	v.value = newvalue
}

func main() {
	v1 := Value{1}
	v2 := v1

	v1.ChangeValue(3)

	v1.Print()
	v2.Print()
}
