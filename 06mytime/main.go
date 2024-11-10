package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.)
	present := time.Now()
	fmt.Println(present)
	fmt.Println(present.Format("02-01-2006 Monday"))
	fmt.Println(present.Format("02/01/2006 15:04:05 Monday"))

	// to create date
	createdDate := time.Date(2020, time.January, 11, 10, 10, 10, 10, time.UTC)
	fmt.Println(createdDate.Format("02/01/2006 Monday"))
}
