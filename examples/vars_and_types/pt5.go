package main

import "fmt"

func main() {
	type User struct {
		name       string
		secondName *string
	}

	f := func() *User {
		secondName := "Doe"
		u := &User{
			name:       "John",
			secondName: &secondName,
		}
		return u
	}

	u := f()

	fmt.Println("user.name", u.name)
	fmt.Println("&user.name", &u.name)

	fmt.Println("user.secondName", u.secondName)
	fmt.Println("&user.secondName", *u.secondName)
}
