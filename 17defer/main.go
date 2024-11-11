package main

import "fmt"

func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
	myDefer()

}

func myDefer() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
