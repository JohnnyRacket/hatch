package main

import (
	"net/http"

	"hatch/rpc/user"
	"hatch/user-service/controllers"
)

func main() {
	server := &controllers.Server{} // implements Haberdasher interface
	twirpHandler := user.NewUserServiceServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
