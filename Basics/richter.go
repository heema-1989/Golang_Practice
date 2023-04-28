package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := os.Args
	if len(inp) != 2 {
		fmt.Println("Give me the magnitude of earthquake")
		return
	}
	switch m, _ := strconv.ParseFloat(inp[1], 64); {
	case m <= 2:
		fmt.Println("minor")
	case m > 2.0 && m <= 3.0:
		fmt.Println("micro")
	case m > 3.0 && m <= 4.0:
		fmt.Println("very minor")
	case m > 4.0 && m <= 5.0:
		fmt.Println("minor")
	case m > 5.0 && m <= 6.0:
		fmt.Println("light")
	case m > 6.0 && m <= 7.0:
		fmt.Println("moderate")
	case m > 7.0 && m <= 8.0:
		fmt.Println("strong")
	case m > 8.0 && m <= 9.0:
		fmt.Println("major")
	case m > 9.0 && m < 10.0:
		fmt.Println("great")
	case m >= 10.0:
		fmt.Println("massive")
	default:
		fmt.Println("Icouldn't get that, sorry")
	}
}
