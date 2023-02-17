package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func responseSize(url string) {
	fmt.Println("Step 1: ", url)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Step2: ", url)
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer wg.Done()
	fmt.Println("Step 3: ", url)
	defer response.Body.Close()
	fmt.Println("Step 4: Length is ", len(databyte))
}
func main() {
	fmt.Println("Goroutines starting....")
	wg.Add(3)
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	wg.Wait()
	//time.Sleep(10 * time.Second) //this will make the main goroutine sleep for 10seconds
	fmt.Println("Goroutines terminating....")
}
