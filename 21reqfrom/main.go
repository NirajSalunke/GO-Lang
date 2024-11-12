package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to our web")
	// PerformGetReq("http://localhost:8000/")
	// PerformPostJSONReq("http://localhost:8000/post")
	PerformPostFormReq("http://localhost:8000/postform")
}

func PerformGetReq(myURl string) {
	// const myURL string = "http://localhost:8000/"
	response, err := http.Get(myURl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // closing
	fmt.Println("Status Code:- ", response.StatusCode)
	fmt.Println("Content length:- ", response.ContentLength)
	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content in the response body:- ", string(content)) // string()
	// instead of using string(), we can get more controll over the data by using strings.Builder
	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Byte COunt", byteCount)
	fmt.Println("String", responseString.String()) // this will give original string
}

func PerformPostJSONReq(myUrl string) {
	// fake JSON Payload
	requestBody := strings.NewReader(`
	{
		"coursename":"Lets GO with GO lang",
		"price":0,
		"platform":"LCO"
	}`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormReq(myURl string) {
	// formdata

	data := url.Values{}
	data.Add("coursename", "Lets GO with GO lang")
	data.Add("name", "Niraj")
	data.Add("email", "nirajsalunke@gmail.com")

	response, err := http.PostForm(myURl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}
