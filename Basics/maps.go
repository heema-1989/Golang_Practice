package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//maps and slice can be declared using make--maps are key-valued
	map1 := make(map[string]string)
	map1["JS"] = "Javascript"
	map1["PY"] = "Python"
	map1["Go"] = "Golang"
	map1["RB"] = "Ruby"
	fmt.Println(map1, reflect.TypeOf(map1))
	fmt.Println(map1["JS"])
	//delete can be used in both map and slice
	delete(map1, "Go")
	fmt.Println(map1)
	//looping in maps
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	str := "1010"
	intStr, err := strconv.ParseInt(str, 8, 32) //here 2 means the mentioned string is converted to binary
	fmt.Println(intStr, err, reflect.TypeOf(intStr))
	map2 := map[string]int{}
	//map2["heema"] = 6566
	fmt.Println(map2)
	var map3 map[string]int = map[string]int{}
	map3["djff"] = 555
	fmt.Println(map3)
	map4 := new(map[string]int)
	fmt.Println(map4, reflect.TypeOf(map4))
	var map5 = map[string]int{}
	map5["sjdd"] = 777
	fmt.Println(map5)
}
