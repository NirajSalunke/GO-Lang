package main

import "fmt"

type User struct { // First  Letter Capital for naming struct
	// Each have capital letter so they can be accessed, exported
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Structs in GO Lang")
	// No inheritance
	niraj := User{
		"niraj",
		"nirajsalunke@gmail.com",
		12,
		true,
	}
	// Trying to print whole
	fmt.Printf("Name of User: %v\nEmail: %v\nAge: %v\nLogged in Status: %v\n  ", niraj.Name, niraj.Email, niraj.Age, niraj.Status)
	fmt.Printf("Niraj's details:- %v\n", niraj)
	fmt.Printf("Niraj's details:- %+v", niraj)

}
