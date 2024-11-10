package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"hello", "Niraj"}
	// fmt.Printf("%T", fruitList)
	fruitList = append(fruitList, "banana", "mango")
	fmt.Println(fruitList)
	// here 0 index will be omitted ,also  3rd and above 3 will be omitted
	fruitList = fruitList[1:3]
	fmt.Println(fruitList)

	highScore := make([]int, 4) // first arg is type of data and second its length
	highScore[0] = 100
	highScore[1] = 300
	highScore[2] = 200
	highScore[3] = 400

	// highScore[4] = 500 , this will cause error as size is 4 but,
	highScore = append(highScore, 500, 700) // this will fulfill

	sort.Ints(highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) // gives bool isSorted or not
	fmt.Println(highScore)

	// How to remove a value from slice based on index

	var courses = []string{"Go", "React", "JavaScript", "Swift", "Python"}

	var indexOfItemToBeRemoved int = 2
	courses = append(courses[:indexOfItemToBeRemoved], courses[indexOfItemToBeRemoved+1:]...)
	fmt.Println(courses)

}
