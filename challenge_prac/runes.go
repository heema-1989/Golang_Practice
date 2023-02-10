package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	rStr := "ajsjdjf"
	fmt.Println("Rune count: ", utf8.RuneCountInString(rStr))
	for i, rvalue := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, rvalue, rvalue)
	}
	fmt.Println(time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println(time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
