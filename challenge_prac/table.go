package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := os.Args

	if len(inp) != 2 {
		fmt.Printf("Wrong size")
		return
	}
	tsize, _ := strconv.Atoi(inp[1])
	if tsize < 0 {
		fmt.Printf("Wrong size")
	} else {
		fmt.Printf("%5s", "X")
		for i := 0; i <= tsize; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for i := 0; i <= tsize; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= tsize; j++ {
				fmt.Printf("%5d", i*j)
			}
			fmt.Println()
		}
	}
	var (
		sum int
		i   int = 1
	)
	for {
		if i > 10 {
			break
		}
		fmt.Printf("%d", i)
		sum += i
		if i == 10 {
			i++
			continue
		}
		fmt.Print("+")
		i++
	}
	fmt.Println("= ", sum)
}
