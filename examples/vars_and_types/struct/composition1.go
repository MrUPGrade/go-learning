package main

import "fmt"

type A struct {
	Field1 string
}

type B struct {
	A

	// Uncomment to see difference
	Field1 string
	Field2 string
}

func testComposition(aField A) {
	fmt.Println("struct of type A", aField)

}

func main() {
	a := A{Field1: "text1"}

	b := B{Field2: "text2", A: a}

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("b.Field2", b.Field2, "b.Field1", b.A.Field1)
	fmt.Println("b.Field2", b.Field2, "b.Field1", b.Field1)

	testComposition(a)

	// Will note compile
	//testComposition(b)
}
