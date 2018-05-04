package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello from app-service~ 📱")
	// Create new router instance
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("hatch/build"))))

	// TODO: handle non-html routes here

	log.Println("Listening at port 8080")
	http.ListenAndServe(":8080", r)
}
