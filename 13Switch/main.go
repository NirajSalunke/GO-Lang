package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case")
	dice := rand.Intn(6) + 1
	switch dice {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough // this statement if there will make compiler  exceute all the below case as true
		// works like no break in switch here!
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("Default")
	}
}
