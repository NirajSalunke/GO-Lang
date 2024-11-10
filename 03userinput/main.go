package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Program"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter ur age:- ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
