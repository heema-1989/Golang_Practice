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

	s1 := make([]int, 2, 3)
	s1[0] = 3
	// s1 = append(s1, 1)
	// s1 = append(s1, 7)
	// s1 = append(s1, 8)
	// s1 = append(s1, 10, 11)
	//s1 = append(s1, 11)
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Println(s1, len(s1), cap(s1))
	s2 := []int{}
	fmt.Println(s2, len(s2))
	s3 := new([2]int)
	s3[0] = 1
	s3[1] = 2
	fmt.Println(s3)
}
