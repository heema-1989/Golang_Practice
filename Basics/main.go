package main

import (
	"fmt"
	//"github.com/go_practice/mascot"
	//"path"
	//"os"
	//"unicode/utf8"
	"strings"
)

// "path"
// "os"
//var Args []string

func main() {
	//fmt.Println(mascot.GetName())
	//var  file string
	//dir, file = path.Split("js/main.js")
	//fmt.Println("Directory name is: ", dir)

	//_, file := path.Split("css/main.css")
	//fmt.Println("File name is: ", file)
	/*speed := 5
	force := 5.23
	speed = int(float64(speed) * force)
	fmt.Println("Speed is: ", speed)*/

	//var Args []string //this args is a string slice which can store multiple string values. HERE STRING VALUES CAN BE ACCESSED USING INDEXING
	//here args is a variable and it stores command line arguments and its path
	/*fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path of Args: ", os.Args[0])
	fmt.Println("1st argument: ", os.Args[1])
	fmt.Println("2nd argument: ", os.Args[2])
	fmt.Println("Length of arguments is: ", len(os.Args))*/

	/*color := "blue"
	color = "red"*/
	/*color := "green"
	color = "Dark" + color*/
	/*n := 0.
	n = 3.14 * 2
	fmt.Println(n)*/
	/*width, height := 5, 6
	var perimeter int
	perimeter = 2 * (width + height)*/
	/*var (
		lang    string
		version int
	)
	lang, version = "Go", 2*/
	/*var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 13.8
	fmt.Println("It is cool on ", planet)
	fmt.Println("Value is ", isTrue)
	fmt.Println("Temperature is: ", temp, " degrees")*/
	/*mascot.Hello()
	var string1 string = "heema"
	string1 = " dhatri"
	fmt.Println(string1)*/
	/*fmt.Println(multi())
	_, int2 := multi()
	fmt.Println(int2)*/
	/*color, color1 := "red", "blue"
	color, color1 = "orange", "green"
	fmt.Println(color, color1)*/
	/*red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)*/
	/*dir, _ := path.Split("secret/file.txt")
	fmt.Println("The directory of file is: ", dir)*/
	/*a, b := 10, 5.5
	fmt.Println(float64(a) + b)*/
	/*a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)*/
	/*age := 2
	fmt.Println(float64(7.5) + float64(age))*/
	/*min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))*/

	/*count := 0
	fmt.Printf("%#v\n", os.Args)
	count = len(os.Args)
	fmt.Printf("There are %d names", count)*/
	//fmt.Printf("Hi %s ", os.Args[1])
	//fmt.Println("hOW ARE YOU!")
	/*var s string
	s = "<html>\n\t<body>\t\"Hello\"\n\t</body>\n</html>"
	s = `	<html>
		<body> "hello" </body>
	</html>`
	s = "c:\\heema\\downloads"
	s = `c:\heema\downlaods`

	fmt.Println(s)*/
	/*var s string
	msg := os.Args[1]
	l := len(msg)
	s = msg + strings.Repeat("!", l)
	fmt.Println(s)*/
	/*json := `{
		"Items": [{
						"Item": {
								"name": "Teddy Bear"
						}
				}]
	}`
	fmt.Println(json)*/
	/*name := os.Args[1]
	s := "Hello" + name + `How are you!`
	fmt.Println(s)*/
	/*input := "İNANÇ"
	l := utf8.RuneCountInString(input)
	fmt.Println(l)*/
	/*s := os.Args[1]
	sl := strings.ToLower(s)
	fmt.Println(sl)*/
	/*msg := `

		The weather looks good.
	I should go and play.
		`
		trim := strings.Trim(msg, "\n\t")
		fmt.Println(trim)*/
	msg := "Hello        "
	tirm := strings.TrimRight(msg, " ")
	fmt.Println(len(tirm))
}

/*func multi() (int, int) {
	return 5, 4
}*/
