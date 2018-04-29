package main

import (
	"hatchery/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {

	//start ticking
	//eggs := StartIncubation()

	//create a pool of harvester to harvest plants
	// for f := 1; f <= 5; f++ {
	//     go Farmer(f, eggs)
	// }
	StartFarmers(256)

	router := mux.NewRouter()
	router.HandleFunc("/beans", controllers.GetEggs).Methods("GET")
	router.HandleFunc("/beans/{id}", controllers.GetEgg).Methods("GET")
	router.HandleFunc("/beans", controllers.CreateEgg).Methods("POST")
	router.HandleFunc("/beans/{id}", controllers.DeleteEgg).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
