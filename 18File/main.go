package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	writeFile("./Mytext.txt", "my name is niraj")
	readFile("./Mytext.txt")
}
func writeFile(filePath string, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:- ", length)
	defer file.Close()
}
func readFile(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// file returns buffer
	fmt.Println(string(file))
	// defer file.Close()
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
