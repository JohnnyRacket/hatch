package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"hatch/app-service/controllers"
	"hatch/app-service/session-manager"
)

func main() {
	// Create new router instance
	r := mux.NewRouter()
	authKey, authKeyErr := ioutil.ReadFile(os.Getenv("SESSION_AUTH_LOCATION"))
	cryptKey, cryptKeyErr := ioutil.ReadFile(os.Getenv("SESSION_CRYPT_LOCATION"))

	if authKeyErr != nil || cryptKeyErr != nil {
		log.Fatal("Could not read keys for session")
	}

	sessionManager := session.NewManager(authKey, cryptKey)
	var authenticationController controllers.Controller = controllers.NewAuthenticationController(sessionManager)

	authenticationController.RegisterRoutes(r)
	// TODO: figure out how to defer 404s here to index.html
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("hatch/build/static"))))
	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("hatch/build/assets"))))
	r.Path("/manifest.json").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hatch/build/manifest.json")
	})
	r.Path("/favicon.ico").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hatch/build/favicon.ico")
	})
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hatch/build/index.html")
	})

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	log.Println(fmt.Sprintf("App Service listening at port %s", os.Getenv("PUBLIC_PORT")))
	log.Fatal(http.ListenAndServe(":8080", handlers.CompressHandler(r)))
}
