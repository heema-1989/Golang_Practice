package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	printDefer()
	//defer works on LIFO characteristics
}
func printDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
