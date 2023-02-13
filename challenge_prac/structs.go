package main
import(
	"fmt"
)
type User struct{
	Name string
	Email string
	status bool
	Age int
}
func main(){
	heema:= User{"Heema","heema@simform",true,20}
	fmt.Println(heema)
	fmt.Printf("Heema's details are: %+v\n",heema)
	fmt.Printf("User name is %v and status is %v",heema.Name,heema.status)
}