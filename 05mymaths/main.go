package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths")
	// var mynumberNums int = 12
	// var myFloatNum float64 = 4
	// fmt.Println("Sum is:- ", mynumberNums+int(myFloatNum))

	// random number from maths/rand
	// fmt.Println(rand.Intn(9))

	// random number from crypto/rand
	randomNum, err := rand.Int(rand.Reader, big.NewInt(5)) // here 5 is max limit of random number generated
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Random Number is :- ", randomNum)

}
