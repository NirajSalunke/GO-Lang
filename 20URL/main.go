package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=uiaaf9134713"

func main() {
	fmt.Println("Welcome to URL")
	fmt.Println(myURL)
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	// To save the params/queries
	qparams := result.Query()
	fmt.Printf("Type of query is %T\n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["paymentid"])
	for _, val := range qparams {
		fmt.Println(val)
	}

	// To construct URl
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev:3000",
		Path:    "/learn",
		RawPath: "user=niraj",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
