package functions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var usr Users

func ValidateUsers(u *[]Users) {
	fmt.Println("Please enter your card number")

	inp, _ := reader.ReadString('\n')

	card_no, err := strconv.Atoi(strings.TrimSpace(inp))
	if err != nil {
		fmt.Println("Please enter valid input. Enter digits only")
		return
	}
	for _, val := range *u {
		if card_no == val.CardNo {
			usr = val
			fmt.Println(val)
			break
		}
	}
	if (usr == Users{}) {
		log.Fatal("User  not found")
	}

	count := 0
PIN:
	for {
		if count == 4 {
			fmt.Println("Maximum limit for entering pin exceeded.Please try again after some time.")
			return
		}

		fmt.Println("Please enter your pin number.")
		inp, _ = reader.ReadString('\n')
		pin_no, err1 := strconv.Atoi(strings.TrimSpace(inp))
		if err1 != nil {
			fmt.Println("Please enter valid input. Enter digits only")
			return
		}
		if pin_no != usr.PinNo {
			fmt.Println("Please enter valid pin number...")
			fmt.Println("Only 3 attempts left.Enter valid pin.")
			count++
			goto PIN
		} else {
			fmt.Println("--------------------------------------------")
			fmt.Println("Successfully logged in to your account..")
			fmt.Println("--------------------------------------------")
			fmt.Printf("Welcome %v\n", usr.Username)
			fmt.Println("--------------------------------------------")
			break
		}
	}
}
