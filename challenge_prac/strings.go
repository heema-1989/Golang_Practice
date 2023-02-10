package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "A word"
	s2 := strings.NewReplacer("A", "Another").Replace(s1)
	fmt.Println("Replaced string is: ", s2)
	fmt.Println("Contains ", strings.Contains(s2, "Another"))
	fmt.Println("To lower: ", strings.ToLower(s2))
	fmt.Println("To upper: ", strings.ToUpper(s2))
	fmt.Println("Has prefix: ", strings.HasPrefix("LavendarHaze", "Lavendar"))
	fmt.Println("Has suffix: ", strings.HasSuffix("LavendarHaze", "Haze"))
	s3 := "\nabcc\n"
	fmt.Println("Remove space: ", strings.TrimSpace(s3))
	fmt.Println("INdex of: ", strings.Index(s2, "o"))
	fmt.Println("Replaced string is: ", strings.Replace(s2, "o", "0", -1))
}
