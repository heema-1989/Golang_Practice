package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("How many lines you want in pyramid")
	reader := bufio.NewReader(os.Stdin)
	lines, _ := reader.ReadString('\n')
	lines = strings.TrimSpace(lines)
	linesNo, _ := strconv.Atoi(lines)
	flag := linesNo
	numFlag := 0
	for i := 1; i <= linesNo; i++ {
		num := i
		for j := 0; j <= flag; j++ {
			fmt.Print("  ")
		}
		k := 1
		for {
			if k > i {
				break
			} else {
				fmt.Printf("%d ", num)
				num++
				k++
			}
			numFlag = num - 1
		}
		m := 1
		for {
			if m >= i {
				break
			} else {
				fmt.Printf("%d ", numFlag-1)
				numFlag--
				m++
			}
		}
		flag--
		fmt.Print("\n")
	}
}
