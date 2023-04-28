package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())
	guess := 10
	for i := 0; i != guess; {
		i = rand.Intn(guess + 1)
		fmt.Printf("%d", i)
	}
	num := 7
	ptr := &num
	*ptr += 1
	fmt.Println(num)
}
