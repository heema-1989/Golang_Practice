package main

import (
	"fmt"
)

func main() {
	var ptr *int
	fmt.Print("Value of pointer is: ", ptr)
	no := 7
	var rptr = &no
	fmt.Println("Address of pointer is: ", rptr)
	fmt.Println("Value of pointer is: ", *rptr)
}
