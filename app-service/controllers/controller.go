package controllers

import "github.com/gorilla/mux"

// Controller interface guarantees a controller can register its routes
type Controller interface {
	RegisterRoutes(mux *mux.Router)
}
