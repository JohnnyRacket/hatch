package main

import (
	"net/http"

	"hatch/rpc/user"
	"hatch/user-service/controllers"
	"hatch/user-service/data"

	"github.com/gorilla/mux"
)

func main() {

	//init data
	db, err := data.Init()
	if err != nil {
		panic("db error")
	}
	//set up repo
	userRepo := data.NewPostgresUserRepository(db)
	emailcodeRepo := data.NewPostgresEmailRepository(db)

	//set up controllers
	userController := controllers.NewUserServiceServer(userRepo, emailcodeRepo)
	//server := &controllers.Server{} // implements Haberdasher interface
	twirpHandler := user.NewUserServiceServer(userController, nil)
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//should allow people to hit the twrip rpc
	router.Handle(user.UserServicePathPrefix, twirpHandler)

	http.ListenAndServe(":8080", twirpHandler)
}
