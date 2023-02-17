package main

import (
	"fmt"
	"sync"
)

func main() {
	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Go routine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Go routine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Go routine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
