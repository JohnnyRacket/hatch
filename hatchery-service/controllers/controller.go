package controllers

import (
	"github.com/gorilla/mux"
)

//Controller interface ensures all controller can register the routes they wish to broadcast
type Controller interface {
	RegisterRoutes(router *mux.Router)
}
