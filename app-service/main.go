package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"hatch/app-service/controllers"
	"hatch/app-service/middleware"
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
	middlewareManager := middleware.NewManager(sessionManager)
	var authenticationController controllers.Controller = controllers.NewAuthenticationController(sessionManager)

	authenticationController.RegisterRoutes(r)
	r.PathPrefix("/").Handler(middlewareManager.ValidateAuthenticationForHandler(http.StripPrefix("/", http.FileServer(http.Dir("hatch/build")))))

	log.Println(fmt.Sprintf("App Service listening at port %s", os.Getenv("PUBLIC_PORT")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
