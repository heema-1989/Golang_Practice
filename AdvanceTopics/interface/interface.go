package main

import "fmt"

type Employee interface {
	PrintName(name string)
	GetSalary(basic int, tax int) int
}

// user defined datatype
type Emp int

func main() {
	var e Employee
	e = Emp(1)
	e.PrintName("Heema Goratela")
	fmt.Println("The in hand salary of employee is: ", e.GetSalary(25000, 2))
}
func (e Emp) PrintName(name string) {
	fmt.Println("The employee id is: ", e)
	fmt.Println("The name of employee is: ", name)
}
func (e Emp) GetSalary(basic int, tax int) int {
	salary := (basic * tax) / 100
	return basic - salary
}
