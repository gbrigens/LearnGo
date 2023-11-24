package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	println("Files")

	content := "This need to be stored in a file"

	file, err := os.Create("./myFile.txt")

	checkNillErr(err)

	length, err := io.WriteString(file, content)

	checkNillErr(err)

	fmt.Println("length is ", length)
	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(fileName string) {

	dataByte, err := os.ReadFile(fileName)

	checkNillErr(err)

	fmt.Println("Text in the file is  ", string(dataByte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
