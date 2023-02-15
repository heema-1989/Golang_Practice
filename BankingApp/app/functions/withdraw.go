package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Users struct {
	Username    string
	CardNo      int
	PinNo       int
	AccountType string
	Balance     float64
}

var reader = bufio.NewReader(os.Stdin)

func WithdrawMoney(u *[]Users) {
WITHDRAW:
	fmt.Println("Please enter amount to be withdrawn.")
	inp, _ := reader.ReadString('\n')
	withAmt, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		fmt.Println("Please enter valid input. Enter digits only")
		goto WITHDRAW
	}
	fmt.Println("Your current balance is: ", usr.Balance)
	if withAmt > usr.Balance {
		fmt.Println("Withdraw amount succeeded the balance amount.Not enough balance.")
		fmt.Println("Please enter lower amount")
		goto WITHDRAW
	} else if withAmt>0 {
		usr.Balance -= withAmt
		fmt.Printf("You have successfully withdrawn %.2f\n", withAmt)
		fmt.Printf("Your new balance is: %.2f\n", usr.Balance)
	}else{
		fmt.Println("Enter only positive amount")
		goto WITHDRAW
	}
}
