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
	fmt.Printf("\nEmployee full name is %v %v\n", e.Firstname, e.Lastname)
	fmt.Printf("\nEmployee email is %v", e.Email)
	fmt.Printf("\nEmployee age is %v", e.Age)
	fmt.Printf("\nEmployee monthly salary is: ")
	for _, value := range e.MonthlySalary {
		fmt.Printf("\n---------------------\n")
		fmt.Println(value.Basic)
		fmt.Println(value.HRA)
		fmt.Println(value.DA)
	}
	return "----------------------------"
}
func main() {
	emp1 := []Employee{
		{"Heema", "Goratela", "heema@simform", 20,
			[]Salary{{
				20000, 500, 200}, {4000, 200, 600}}},
		{"Dhatri", "Goratela", "dhatrig@", 16, []Salary{
			{20000, 500, 700}, {5000, 600, 200}}}}
	for _, val := range emp1 {
		fmt.Println(emp_info(val))
	}
}
