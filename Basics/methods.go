package main

import (
	"fmt"
)

type User struct {
	Name   string
	status bool
	email  string
}

func (u User) getDetails() {
	fmt.Println(u.Name, u.status, u.email)
}
func main() {
	user1 := User{"Heema", true, "heema@simform"}

	//while we divide by zero in go we get an integer divide by zero error
	user1.getDetails()
	a := 14.7899
	fmt.Printf("%8.2f\n", a)
	fmt.Printf("%g\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%e\n", a)
	fmt.Printf("%o\n", a)
}
