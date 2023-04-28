package main
import(
	"fmt"
)
func main(){
	result:= add(4,5)
	fmt.Println(result)
	prores,_:= proadder(2,3,4)
	fmt.Println(prores)
}
func add(no1 int,no2 int)int{
	return no1+no2
}
func proadder(values ...int)(int,string){
	total:=0
	for _,val:= range values{
		total+=val
	}
	return total,"hello"
}