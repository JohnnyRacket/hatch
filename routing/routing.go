package routing

import (
	"hatchery/controllers"

	"github.com/gorilla/mux"
)

//InitializeRoutes creates the routes for the application that will be available via web api
func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/beans", controllers.GetEggs).Methods("GET")
	router.HandleFunc("/beans/{id}", controllers.GetEgg).Methods("GET")
	router.HandleFunc("/beans", controllers.CreateEgg).Methods("POST")
	router.HandleFunc("/beans/{id}", controllers.DeleteEgg).Methods("DELETE")
}
