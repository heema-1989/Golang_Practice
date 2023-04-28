package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage = "Usage: [username] [password]"
	user  = "Jack"
	pwd   = "1888"
	user1 = "Heema"
	pwd1  = "1989"
)

func main() {
	input := os.Args
	if len(input) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := os.Args[1], os.Args[2]
	if u != strings.ToLower(user) {
		fmt.Printf("\nAccess denied for %s", u)
	} else if p != pwd {
		fmt.Printf("\nInvalid password for %s", u)
	} else {
		fmt.Printf("\nAccess granted to user %q", u)
	}
}
