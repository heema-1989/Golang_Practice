package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := os.Args
	if len(inp) != 3 {
		fmt.Println("Usage: [year] [month]")
		fmt.Println("Please enter a year and month name")
		return
	}
	year, err := strconv.Atoi(inp[1])
	if err != nil {
		fmt.Printf("%q is not a year", inp[1])
	} else {

		if strings.ToLower(inp[2]) == "january" {
			fmt.Printf("%q has 31 days", inp[2])
		} else if strings.ToLower(inp[2]) == "february" {
			if year%4 == 0 {
				fmt.Printf("%q has 29 days", inp[2])
			} else {
				fmt.Printf("%q has 28 days", inp[2])
			}
		} else if strings.ToLower(inp[2]) == "march" {
			fmt.Printf("%q has 31 days", inp[2])
		} else if strings.ToLower(inp[2]) == "april" {
			fmt.Printf("%q has 30 days", inp[2])
		} else if strings.ToLower(inp[2]) == "may" {
			fmt.Printf("%q has 31 days", inp[2])
		} else if strings.ToLower(inp[2]) == "june" {
			fmt.Printf("%q has 30 days", inp[2])
		} else if strings.ToLower(inp[2]) == "july" {
			fmt.Printf("%q has 31 days", inp[2])
		} else if strings.ToLower(inp[2]) == "august" {
			fmt.Printf("%q has 31 days", inp[2])
		} else if strings.ToLower(inp[2]) == "september" {
			fmt.Printf("%q has 30 days", inp[2])
		} else if strings.ToLower(inp[2]) == "october" {
			fmt.Printf("%q has 31 days", inp[2])
		} else if strings.ToLower(inp[2]) == "november" {
			fmt.Printf("%q has 30 days", inp[2])
		} else if strings.ToLower(inp[2]) == "december" {
			fmt.Printf("%q has 31 days", inp[2])
		}
	}
}
