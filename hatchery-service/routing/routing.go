package routing

import (
	"hatch/hatchery-service/controllers"

	"github.com/gorilla/mux"
)

//InitializeRoutes creates the routes for the application that will be available via web api
func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/eggs", controllers.GetEggs).Methods("GET")
	router.HandleFunc("/eggs/{id}", controllers.GetEgg).Methods("GET")
	router.HandleFunc("/eggs", controllers.CreateEgg).Methods("POST")
	router.HandleFunc("/eggs/{id}", controllers.DeleteEgg).Methods("DELETE")
}
