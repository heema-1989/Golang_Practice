package main

import "fmt"

type Vehicle interface {
	Structure() []string
	Speed() string
}
type Human interface {
	Structure() []string
	Performance() string
}
type Car string
type Man string

func (c Car) Structure() []string {
	name := []string{"Gasfilter", "Wiper", "Tank", "Engine"}
	fmt.Println("Car name is: ", c)
	return name
}
func (c Car) Speed() string {
	return "800km/hr"
}
func (m Man) Structure() []string {
	name := []string{"john", "joe", "jake", "harry"}
	//fmt.Println("Person name is: ", m)
	return name
}
func (m Man) Performance() string {
	return "8hrs/day"
}
func main() {
	var car Vehicle
	car = Car("Breeza")
	fmt.Println(car.Structure())
	var per Human
	per = Man("Labours")
	fmt.Println(per.Structure())
	for i, val := range car.Structure() {
		fmt.Printf("%-15s<=====================>%15s\n", val, per.Structure()[i])
	}
}
