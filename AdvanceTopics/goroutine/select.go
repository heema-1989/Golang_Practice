package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	quit := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "hey"

	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hello"

	}()
	go func() {
		time.Sleep(4 * time.Second)
		quit <- 0
	}()
	for {
		select {
		case res1 := <-ch1:
			fmt.Println("I received from channel 1 ", res1)

		case res2 := <-ch2:
			fmt.Println("I received from channel 2 ", res2)

		case <-quit:
			fmt.Println("Quitting")
			return
		}

	}

	//time.Sleep(4 * time.Second)
}
