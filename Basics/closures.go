package main

import "fmt"

func Adder() func(int) int {
	n := 0
	return func(no int) int {
		n = n + no
		return n
	}
}
func main() {
	a := Adder()
	fmt.Println(a(10))
	fmt.Println(a(20))
	fmt.Println(a(30))
}
