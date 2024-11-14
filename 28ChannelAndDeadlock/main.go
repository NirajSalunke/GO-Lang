package main

import "fmt"

func main() {
	fmt.Println("Channels in golang")

	myChannel := make(chan int)
	myChannel <- 5
	fmt.Println(<-myChannel)

}
