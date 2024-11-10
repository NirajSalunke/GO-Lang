package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go lang")
	var fruitlist [5]string
	fruitlist[0] = "hello"
	fruitlist[1] = "my"
	// fruitlist[2] = 3
	fruitlist[3] = "is"
	fruitlist[4] = "Niraj"
	fmt.Println(len(fruitlist))
	var vegList = [3]string{"Veggies", "potates", "mushroom"}
	fmt.Println(vegList)
}
