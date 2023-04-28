package main

import (
	"fmt"
	"reflect"
	"time"
	"unicode/utf8"
)

func main() {
	rStr := "ajsjdjf"
	fmt.Println(reflect.TypeOf(rStr))
	input := "İNANÇ"
	fmt.Println(reflect.TypeOf(input))
	fmt.Println("Input length: ", len((input)))
	fmt.Println("Rune count: ", utf8.RuneCountInString(input))
	fmt.Println("Rune count: ", utf8.RuneCountInString(rStr))
	for i, rvalue := range input {
		fmt.Printf("%d : %#U : %c\n", i, rvalue, rvalue)
	}
	a := 5
	fmt.Printf("\n%U\n", a)
	fmt.Printf("\n%#U\n", a)
	fmt.Println(time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println(time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
