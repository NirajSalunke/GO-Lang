package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/gorilla/mux"
)

// model for course -file
type Course struct {
	CourseID    string  `json:"-"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	AuthorName  *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
// Seeding of data
var courses = []Course{
	{
		CourseID:    "1",
		CourseName:  "Learn Go Language",
		CoursePrice: 499,
		AuthorName: &Author{
			FullName: "Hitesh Choudhary",
			Website:  "https://hiteshchoudhary.com",
		},
	},
	{
		CourseID:    "2",
		CourseName:  "Master ReactJS",
		CoursePrice: 999,
		AuthorName: &Author{
			FullName: "Traversy Media",
			Website:  "https://www.traversymedia.com",
		},
	},
	{
		CourseID:    "3",
		CourseName:  "Intro to Web Development",
		CoursePrice: 299,
		AuthorName: &Author{
			FullName: "Angela Yu",
			Website:  "https://www.udemy.com/user/angela-yu/",
		},
	},
	{
		CourseID:    "4",
		CourseName:  "Advanced Python Programming",
		CoursePrice: 699,
		AuthorName: &Author{
			FullName: "Corey Schafer",
			Website:  "https://coreyms.com",
		},
	},
	{
		CourseID:    "5",
		CourseName:  "Data Structures and Algorithms",
		CoursePrice: 799,
		AuthorName: &Author{
			FullName: "Abdul Bari",
			Website:  "https://www.udemy.com/user/abdul-bari-35/",
		},
	},
}

// middlewares/helpers/utility functions -file
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/create", createOneCourse).Methods("POST")
	router.HandleFunc("/update/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}/delete", deleteCourse).Methods("DELETE")
	router.HandleFunc("/delete", deleteAll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5173", router))

}

// controllers -file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my worldddd</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course ")
	w.Header().Set("Content-Type", "application/json")
	// grab id from req
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given id")
	return
}

// Creating Course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course ")
	w.Header().Set("Content-Type", "application/json")
	// if body empty'
	fmt.Println(r.Body)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Pls send some Data")
		return
	}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		panic(err)
	}

	// is req.body has no name!
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Empty JSON")
		return
	}
	for _, val := range courses {
		if course.CourseName == val.CourseName {
			json.NewEncoder(w).Encode("ALready exists")
			return
		}
	}
	// creating unique ID
	course.CourseID = string(rand.IntN(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course ")
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, course := range courses {
		if course.CourseID == id {

			// we searched for a course, removed it , updated and added again!
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			err := json.NewDecoder(r.Body).Decode(&newCourse)
			if err != nil {
				panic(err)
			}
			newCourse.CourseID = id
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No Course Found with given id")
		return
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course ")
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, course := range courses {
		if course.CourseID == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course have being Deleted Successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given id")
	return

}

func deleteAll(w http.ResponseWriter, r *http.Request) {
	courses = []Course{}
	json.NewEncoder(w).Encode("Successfully deleted all!")
}
