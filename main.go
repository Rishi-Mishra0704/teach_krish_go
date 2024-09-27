package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

type name string

func (u User) String() {
	fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
}
func (n name) String() {
	fmt.Println(n)
}

func print(c Converter) {
	c.String()
}

func main() {
	user := User{
		Name:  "Krish",
		Email: "something@xyz",
	}

	print(user)
	print(name("Krish"))
}
