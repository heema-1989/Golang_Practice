package main

import "fmt"

func Printer(i interface{}) {
	fmt.Println(i)
}
func main() {
	var i interface{}
	i = "Go language"
	fmt.Println(i)
	var shape = []string{"circle", "triangle"}
	Printer(shape)
	var animal = map[string]string{"NAME": "dog", "type": "domestic"}
	Printer(animal)
	var name = []string{"heema", "dhatri"}
	Printer(name)
	var per = [3]int{1, 2, 3}
	Printer(per)
}
