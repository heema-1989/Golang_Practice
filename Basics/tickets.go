package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	age int
	err error
)

func main() {
	inp := os.Args
	if len(inp) != 2 {
		fmt.Println("PLease enter user age")
		return
	}
	if age, err = strconv.Atoi(inp[1]); err != nil {
		fmt.Printf("Please enter a number")
		return
	} else {
		if age >= 17 {
			fmt.Println("R-Rated")
		} else if age >= 13 && age <= 17 {
			fmt.Println("PG-13")
		} else if age <= 0 {
			fmt.Println("PLease enter positive number")
		} else {
			fmt.Println("PG-Rated")
		}
	}

}
