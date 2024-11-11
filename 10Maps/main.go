package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in GO lang")
	lang := make(map[string]string) // using make to create to avoid errors
	lang["js"] = "JavaScript"
	lang["rb"] = "Ruby"
	lang["ts"] = "TypeScript" // adding values
	lang["py"] = "Python"

	fmt.Println(lang)
	fmt.Println(lang["js"])

	// To delete a key value
	delete(lang, "js")
	fmt.Println(lang)
	// Loops in the Maps
	for key, value := range lang {
		fmt.Printf("Key is %v, and value is %v\n", key, value)
		// here %T gives Type of it,
		// and %v gives value of it.
	}
	// If there is no need of key in the program
	for _, value := range lang { // just replace with key with _
		fmt.Printf("Value is %v\n", value)
	}
}
