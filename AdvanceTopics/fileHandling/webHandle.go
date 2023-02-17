package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url1 = "https://go.dev/doc/effective_go#interfaces_and_types"
)

func main() {
	response, err := http.Get(url1) //making the connection using http get method
	if err != nil {
		panic(err)
	}
	//defer response.Body.Close() //caller's responsibility to close the connection
	//databyte, err := ioutil.ReadAll(response.Body)
	//content := string(databyte)
	fmt.Printf("The type of response is: %T\n", response)
	//fmt.Println("The content of response is: ", content)
	head, err := ioutil.ReadAll(response.Header)
	if err != nil {
		panic(err)
	}
	headerContent := string(head)
	fmt.Println("The header content is: ", headerContent)
	defer response.Header.Close()

}
