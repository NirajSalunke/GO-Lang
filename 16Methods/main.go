package main

import (
	"fmt"
	"strings"
)

type User struct { // First  Letter Capital for naming struct
	// Each have capital letter so they can be accessed, exported
	Name   string
	Email  string
	Age    int
	Status bool
}

func (myUser User) getUser() { // this is a method of struct
	fmt.Println("Name is ", myUser.Name)
}
func (myUser *User) NewMailChangePerman() { // this is a method of struct
	myUser.Email = strings.Split(myUser.Email, "@")[0] + "@go.dev"
	fmt.Println("New Email is ", myUser.Email)
}
func (myUser User) NewMailChangeTemp() { // this is a method of struct
	myUser.Email = strings.Split(myUser.Email, "@")[0] + "@go.dev"
	fmt.Println("New Email is ", myUser.Email)
}

func main() {
	fmt.Println("Methods  in GO Lang")
	niraj := User{
		"niraj",
		"nirajsalunke@gmail.com",
		12,
		true,
	}

	niraj.getUser() // using the method of struct
	niraj.NewMailChangePerman()
	fmt.Println(niraj.Email)

}
