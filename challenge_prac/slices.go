package main

import (
	"fmt"
	"reflect"
	"sort"
)
func main(){
	var l1 = [4]int{1,2,3,4} //array
	fmt.Println(l1,reflect.TypeOf(l1))
	//l1= append(l1,6,7)// hence you cannot append in array becoz it has fixed size
	var l2 = []int{2,3,5,6} // here the size of array not defined so its a slice and use append in slice only
	fmt.Println(l2)
	l2= append(l2,8,9)
	fmt.Println(l2)
	//defining slice using make()--or allocating memory using make
	 l3:=make([]int,2)
	 l3[0]=1
	 l3[1]=2
	 //here using make we cane use append--this make reallocates the memory
	 l4:= make([]int,4)
	 l4[0]=34
	 l4[1]=4
	 l4[2]=15
	 l4[3]=2
	 l4= append(l4,67,89)
	//this append will append into l4 list and overwrite them with starting from index 1 to last. so it will remove 34 from list
	 l4= append(l4[1:])
	 //in slice we can use sort method to sort the slice. we cant use this in array
	sort.Ints(l4)
	fmt.Println(l4)
	fmt.Println(sort.IntsAreSorted(l4))
	fmt.Println(sort.IntsAreSorted(l2))// here if slice is bydefault sorted then also it will give true
	//remove value from slice based on index in golang
	index:=2
	l2= append(l2[:index],l2[index+1:]...)
	fmt.Println(l2)

}
