package main

import (
	"fmt"
	"os"
)

func main() {

	ch := os.Args
	l := len(ch)

	if l == 1 {
		fmt.Println("Give me a letter")
	} else if os.Args[1] == "a" || os.Args[1] == "e" || os.Args[1] == "i" || os.Args[1] == "o" || os.Args[1] == "u" {
		fmt.Printf("%s is a vowel", os.Args[1])
	} else if os.Args[1] == "w" || os.Args[1] == "y" {
		fmt.Printf("%s is sometimes vowel and sometimes consonant", os.Args[1])
	} else {
		fmt.Printf("%s is a consonant", os.Args[1])
	}

}
