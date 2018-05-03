package main

import (
	"hatchery/data"
	"hatchery/engine"
	"hatchery/routing"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {

	//starts the function that will evaluate when to alert users
	data.InitializeRepository()
	engine.NewIncubator()
	//creates a new router
	router := mux.NewRouter()
	//set the routes
	routing.InitializeRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
