package main

import (
	//"bufio"

	"fmt"
	//"log"
	//"os"
	//"reflect"
	//"strconv"
)

var println = fmt.Println

func Hello() {
	/*println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		println("Hello", name)
	} else {
		log.Fatal(err)
	}*/
	/*println(reflect.TypeOf(5.5))
	println(reflect.TypeOf(5))
	println(reflect.TypeOf(true))
	println(reflect.TypeOf("true"))*/
	/*var (
		a string = "500000"
		d string = "255.5"
	)*/
	/*b, err := strconv.Atoi(a)
	println(a, err, reflect.TypeOf(b))
	c := strconv.Itoa(b)
	println(reflect.TypeOf(b), reflect.TypeOf(c))*/
	/*d = "NaN"
	//d = "Inf"
	d = "Infinity"
	d = "heema"
	e, err := strconv.ParseFloat(d, 32)
	println(d, err, reflect.TypeOf(e), e)*/
	const (
		November = 11 - iota
		October
		September
	)
	fmt.Println(November, October, September)
}
func main() {
	//Hello()
}
