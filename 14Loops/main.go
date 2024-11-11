package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ { // there is ++d
		fmt.Println(days[d])
	}
	// does the same as above
	for i := range days {
		fmt.Println(days[i])
	}

	// Like forEach
	for _, day := range days { // as i dont need index here i replaced _
		fmt.Println(day)
	}

	// while loops kinda
	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 2 {
			fmt.Println("Continuing")
			rogueValue++ // hence a incereamnt, if not given
			continue     // here it go to infinite loop as continue skips all below and rogueValue will be 2 forever

		} else if rogueValue == 4 {
			goto lco
		} else if rogueValue == 7 {
			fmt.Println("BReaking here")
			break
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Yeahhhh goto is here")
}
