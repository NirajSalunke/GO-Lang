package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web req")
	const url string = "https://hiteshchoudhary.com"
	reponse, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	// fmt.Println(reponse.Body)
	defer reponse.Body.Close()

	data, err := io.ReadAll(reponse.Body)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
