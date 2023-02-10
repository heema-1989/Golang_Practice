package main

import (
	"fmt"
	"os"
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
	if u != user && u != user1 {
		fmt.Printf("Access denied for %q user", u)
	} else if p == pwd && u == user || p == pwd1 && u == user1 {
		fmt.Printf("Access granted for %q user", u)
	} else {
		fmt.Printf("INavlid password for %q user", u)
	}
}
