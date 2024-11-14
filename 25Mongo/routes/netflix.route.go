package routes

import (
	"github.com/NirajSalunke/modules/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/create", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", controllers.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controllers.DeleteAll).Methods("DELETE")

	return router
}
