package main

import (
	"fmt"
	"net/url"
	"reflect"
)

const myurl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=dufi44"

func main() {
	purl, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(myurl)
	fmt.Println(purl.Scheme)
	fmt.Println(purl.Host)
	fmt.Println(purl.Port())
	fmt.Println(purl.Path)
	fmt.Println(purl.RawPath)
	fmt.Println("Raw query: ", purl.RawQuery)
	fmt.Println("Query: ", purl.Query())
	fmt.Println(reflect.TypeOf(purl.Query()))
	for _, val := range purl.Query() {
		fmt.Println("Values are: ", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "golangprograms.com",
		Path:    "/go-language.interface.html",
		RawPath: "user=heema",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
