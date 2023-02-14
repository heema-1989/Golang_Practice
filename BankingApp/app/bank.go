package main

import (
	"BankingApp/app/functions"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	users := []functions.Users{{"Heema Goratela", 6351545, 1905, "Savings", 40000},
		{"Dhatri Goratela", 9428319, 2006, "Fixed Deposit", 50000},
		{"Khushi Rajpal", 4567834, 1876, "Recurring", 30000},
		{"Riya Sonara", 3498744, 2060, "Savings", 40000}}
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("Welcome to our ATM.Please enter your credentials as asked.")
	fmt.Println("--------------------------------------------------------------------------------------")
	functions.ValidateUsers(&users)
	for {
		fmt.Println("Please choose any one from the below mentioned options..")
		fmt.Println("--------------------------------------------")
		fmt.Println("1. Withdraw money from ATM.")
		fmt.Println("2. Deposit money to ATM.")
		fmt.Println("3. Check Account Balance.")
		fmt.Println("4. Exit.")
		inp, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(inp))
		if err != nil {
			fmt.Println("Please enter valid input. Enter digits only")

		}

		switch choice {
		case 1:
			functions.WithdrawMoney(&users)
		case 2:
			functions.DepositMoney(&users)
		case 4:
			break
		}
	}

	fmt.Println("Thank you for choosing our ATM")
	return
}
