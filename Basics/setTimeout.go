package main

import (
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get("http://golangprograms.com")
	if err != nil {
		panic(err)
	}
}
