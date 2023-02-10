package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := os.Args
	if len(inp) != 2 {
		fmt.Println("Please enter a number")
		return
	}
	no, err := strconv.Atoi(inp[1])
	if err != nil {
		fmt.Printf("%s is not a number", inp[1])
	} else {
		if no%8 == 0 {
			fmt.Printf("\nIt is an even number divisible by 8")
		} else if no%2 == 0 {
			fmt.Printf("\nIt is an even number")
		} else {
			fmt.Printf("\nIt is an odd number")
		}
	}
}
