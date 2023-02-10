package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := os.Args
	if len(inp) == 2 {
		feet, err := strconv.Atoi(os.Args[1])
		meter := 0.
		if err != nil {
			fmt.Printf("ERROR: PLease enter a number")
			return
		}
		meter = (float64(feet) * 0.3048)
		fmt.Printf("%d feet is equal to %f meters", feet, meter)
	}
}
