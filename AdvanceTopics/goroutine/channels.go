package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var w sync.WaitGroup

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
func main() {
	ch := make(chan int)
	/*go func() {
		ch <- 789
	}() //adding data to channel
	n := <-ch
	fmt.Printf("THe channel number is: %d", n)*/

	go func() {
		w.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				ch <- doWork()
			}
			close(ch)
			defer w.Done()
		}()
		w.Wait()
	}()
	for i := range ch {
		fmt.Printf("%d\n", i)
	}

}
