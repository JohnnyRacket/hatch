package main

import (
	"fmt"
	"hatch/hatchery-service/controllers"
	"hatch/hatchery-service/data"
	"hatch/hatchery-service/engine"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {

	//creates a new router
	router := mux.NewRouter()
	//initialize repositories
	eggRepository := data.NewPostgresRepository()
	//have controllers register their routes
	eggController := controllers.NewEggController(&eggRepository)
	eggController.RegisterRoutes(router)

	//starts the function that will evaluate when to alert users
	engine.NewIncubator(&eggRepository)

	fmt.Println("hi there from hatchery service")

	log.Fatal(http.ListenAndServe(":8080", router))
}
