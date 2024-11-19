package main

import "fmt"

func main() {
	fmt.Println("Channels in golang")

	channelmine := make(chan int)
	channelmine <- 5
	fmt.Println(<-channelmine)

}
