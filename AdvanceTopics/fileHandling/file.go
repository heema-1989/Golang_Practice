package main

func main() {
	/*content := "This is the content of the file"
	file, err := os.Create("./myfile.txt")
	checkNilErr(err)*/

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
