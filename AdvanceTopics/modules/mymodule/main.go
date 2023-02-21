package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("Welcome to golang series")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}
