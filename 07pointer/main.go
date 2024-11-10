package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")
	// Creation of Pointer but it point's to nothing
	// var ptr *int
	// fmt.Println(ptr)

	car := 10
	// Creating a pointer to the variable car using &
	ptr := &car
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr *= 10
	fmt.Println(car)

}
