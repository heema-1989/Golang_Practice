package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func DepositMoney(u *[]Users) {

	fmt.Println("Please enter amount to be deposited.")
	inp, _ := reader.ReadString('\n')
	depAmt, _ := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	fmt.Println("Your current balance is: ", usr.Balance)

	usr.Balance += depAmt
	fmt.Printf("You have successfully withdrawn %.2f\n", depAmt)
	fmt.Printf("Your current balance is: %.2f\n", usr.Balance)

}
