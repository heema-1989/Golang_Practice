package main

import (
	"fmt"
	"sync"
)

var x = 0

// solving race condition using mutex
func increment(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	defer w.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1) //ADDING ONE GOROUTINE IN EACH ITERATION
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("Final value of x: ", x)
}
