package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NirajSalunke/modules/routes"
)

func main() {
	fmt.Println("MongoDb APi")
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 5173")
	log.Fatal(http.ListenAndServe(":5173", routes.Router()))

}
