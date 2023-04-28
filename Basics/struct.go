package main

import "fmt"

type Rectangle struct {
	length, breadth int
	color           string
	geometry        Geometry
}
type Geometry struct {
	area, perimeter int
}

func main() {
	rect := Rectangle{}
	rect.length = 3
	rect.breadth = 4
	rect.geometry.area = rect.length * rect.breadth
	fmt.Printf("%#v\n", rect)
	var rect1 = new(Rectangle)
	fmt.Print(rect1)
	rect2 := new(Rectangle)
	fmt.Println("\n", rect2)
	rect3 := &Rectangle{}
	(*rect3).length = 4
	rect3.breadth = 3
	fmt.Println(rect3)
}
