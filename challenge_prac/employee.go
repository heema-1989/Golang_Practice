package main

import (
	"fmt"
)

type Salary struct {
	Basic, HRA, DA float64
}
type Employee struct {
	Firstname, Lastname, Email string
	Age                        int
	MonthlySalary              []Salary
}

func emp_info(e Employee) string {
	fmt.Println("Employee details are: ")

}
func main() {
	emp1 := Employee{"Heema", "Goratela", "heema@simform", 20,
		[]Salary{Salary{
			20000, 500, 200}, Salary{4000, 200, 600}}}

}
