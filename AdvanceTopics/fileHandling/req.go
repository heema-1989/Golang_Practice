package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/example", context)
	http.ListenAndServe(":3000", nil)
}
func context(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "example run")
	case <-context.Done():
		//err := context.Err()
		fmt.Fprintf(w, "server error")

	}
}
