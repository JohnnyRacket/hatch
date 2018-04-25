package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"hatchery/controllers"
)


// our main function
func main() {

	//start ticking
	eggs := StartIncubation()

	//create a pool of harvester to harvest plants
	for f := 1; f <= 5; f++ {
        go Farmer(f, eggs)
    }

	router := mux.NewRouter()
	router.HandleFunc("/beans", controllers.GetEggs).Methods("GET")
	router.HandleFunc("/beans/{id}", controllers.GetEgg).Methods("GET")
	router.HandleFunc("/beans", controllers.CreateEgg).Methods("POST")
	router.HandleFunc("/beans/{id}", controllers.DeleteEgg).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}