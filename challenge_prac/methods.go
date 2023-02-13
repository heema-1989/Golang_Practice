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
	user1.getDetails()
}
