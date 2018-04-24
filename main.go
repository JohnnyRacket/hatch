package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/beans", GetEggs).Methods("GET")
	router.HandleFunc("/beans/{id}", GetEgg).Methods("GET")
	router.HandleFunc("/beans", CreateEgg).Methods("POST")
	router.HandleFunc("/beans/{id}", DeleteEgg).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}