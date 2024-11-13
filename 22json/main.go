package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"` // this name will be used when JSON data is made using this!
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // Wont get this in JSON, no need of passwords there!
	Tags     []string `json:"tags,omitempty"` // here omitempty means
	// here omitempty means if val null then it won't include in JSON
}

func main() {
	fmt.Println("Welcome")
	locoCourse := []course{
		{
			Name:     "Go Programming",
			Price:    299,
			Platform: "Udemy",
			Password: "golang123",
			Tags:     []string{"programming", "golang", "backend"},
		},
	}
	// EncodeJSON(locoCourse)
	// DecodeJSON(locoCourse)
	CreatingMap(locoCourse)
}

func EncodeJSON(locoCourse []course) []byte {
	// fmt.Println(locoCourse)
	finalJSON, err := json.MarshalIndent(locoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(finalJSON))
	return finalJSON
}

func DecodeJSON(locoCourse []course) []course {
	finalJSON := EncodeJSON(locoCourse)
	fmt.Println(string(finalJSON))
	var decodedJSON []course
	if json.Valid(finalJSON) {
		fmt.Println("JSON was Valid")
		json.Unmarshal(finalJSON, &decodedJSON)
	} else {
		fmt.Println("JSON was not Valid")
	}
	return decodedJSON
}

func CreatingMap(locoCourse []course) {
	finalJSON, err := json.MarshalIndent(locoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	var onlineData []map[string]interface{}
	json.Unmarshal(finalJSON, &onlineData)
	fmt.Println(onlineData[0]["price"])
}
