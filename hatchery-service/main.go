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
	//init db
	db, err := data.Init()
	if err != nil {
		panic(err)
	}
	//initialize repositories
	eggRepository := data.NewPostgresRepository(db)
	//have controllers register their routes
	eggController := controllers.NewEggController(&eggRepository)
	eggController.RegisterRoutes(router)

	//starts the function that will evaluate when to alert users
	engine.NewIncubator(&eggRepository)

	fmt.Println("hi there from hatchery service")

	log.Fatal(http.ListenAndServe(":8080", router))
}
