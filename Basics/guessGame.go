package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	randSecs := time.Now().Unix()
	rand.Seed(randSecs)
	randInt := rand.Intn(51)
	for true {
		fmt.Print("Guess a number between 0 and 50")
		fmt.Printf("Random number is: %d", randInt)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iguess, e := strconv.Atoi(guess)

		if e != nil {
			log.Fatal(e)
		}
		if iguess > randInt {
			fmt.Println("Choose a lower number")
		} else if iguess < randInt {
			fmt.Println("Choose a larger number")
		} else {
			fmt.Println("Correct guess")
			break
		}

	}
}
