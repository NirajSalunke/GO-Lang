package main

import "fmt"

func main() {
	fmt.Println("if Else in go")
	loginAcc := 10
	var response string
	if loginAcc > 9 {
		response = "Yeah new user"
	} else if loginAcc == 10 {
		response = "Yeahh"
	} else {
		response = "doesn't matter"
	}
	fmt.Println(response)

	if num := 3; num > 10 {
		fmt.Println("Yeah")
	} else {
		fmt.Println("Noo")
	}
}
