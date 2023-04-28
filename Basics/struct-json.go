package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type employee struct {
	Name string `json:"name"`
	City string `json:"city"`
}

// assign default value to struct field
func (e *employee) assignDefault() {
	if e.Name == "" {
		e.Name = "default"
	}
}
func main() {
	jsonStr := `{ "name":"Heema", "city": "Jmanagar"}`
	var emp = new(employee)
	jsonErr := json.Unmarshal([]byte(jsonStr), &emp)
	if jsonErr != nil {
		log.Fatal("Error parsing json")
	}
	fmt.Println("Employee 1: ", emp)
	emp2 := new(employee)
	emp2.Name = "Heema"
	emp2.City = "Jmanagar"
	jsonBody, err := json.MarshalIndent(emp2, "", "\t")
	if err != nil {
		log.Fatal("Error encoding json")
	}
	fmt.Println("Employee 2: ", string(jsonBody))
	emp3 := employee{City: "Jamnagar"}
	emp3.assignDefault()
	fmt.Println("Employee 3: ", emp3)
	fmt.Printf("\nEmp type is %v: ", reflect.TypeOf(emp))
	fmt.Printf("\nEmp type is %v: ", reflect.ValueOf(emp).Kind())
	fmt.Printf("\nEmp type is %v: ", reflect.ValueOf(emp2).Kind())
	fmt.Printf("\nEmp3 type is %v: ", reflect.TypeOf(emp3))
	fmt.Printf("\nEmp3 type is %v: \n", reflect.ValueOf(emp3).Kind())
	fmt.Printf("\nEmp3 type is %d: \n", reflect.ValueOf(emp3).Kind())
	fmt.Printf("\nEmp3 type is %+v: \n", emp3)
	fmt.Printf("\nEmp3 type is %#v: \n", emp3)
	if emp == emp2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	var emp4 = new(employee)
	emp5 := &employee{}
	if emp4 == emp5 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
