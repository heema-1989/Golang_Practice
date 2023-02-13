package main
import (
	"fmt"
	"reflect"
)
func main(){
	//maps and slice can be declared using make--maps are key-valued
	map1:= make(map[string]string)
	map1["JS"]="Javascript"
	map1["PY"]="Python"
	map1["Go"]="Golang"
	map1["RB"]="Ruby"
	fmt.Println(map1,reflect.TypeOf(map1))
	fmt.Println(map1["JS"])
	//delete can be used in both map and slice
	delete(map1,"Go")
	fmt.Println(map1)
	//looping in maps
	for key,value := range map1{
		fmt.Println(key,value)
	}
}