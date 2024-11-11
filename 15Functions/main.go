package main

import "fmt"

func main() {
	// fmt.Println("")
	gretter("niraj")
	prt := gretter // stores address of the function if written without ()
	fmt.Println(prt)
	fmt.Println(adder(2, 3))
	fmt.Println(proAdder(2, 3, 4, 4, 4, 4, 2, 2, 3))
}
func gretter(name string) {
	fmt.Println("Namaste from GO Lang", name)
}

func proAdder(values ...int) (int, string) { // this ensures arg as slice and they can add infinite  args
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Yeah the ans is"
}

func adder(a int, b int) int {
	return a + b
}
