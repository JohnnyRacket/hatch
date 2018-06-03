package controllers

import (
	"encoding/json"
	"hatch/hatchery-service/data"
	"hatch/hatchery-service/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//EggController servers to manage the functions we open up acces to for eggs
type EggController struct {
	repository data.EggRepository
}

//NewEggController creates a new EggController with the given repo
func NewEggController(repo data.EggRepository) *EggController {
	return &EggController{repository: repo}
}

//RegisterRoutes creates the routes for the application that will be available via web api
func (c *EggController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/eggs", c.GetEggs).Methods("GET")
	router.HandleFunc("/eggs/{id}", c.GetEgg).Methods("GET")
	router.HandleFunc("/eggs", c.CreateEgg).Methods("POST")
	router.HandleFunc("/eggs/{id}", c.DeleteEgg).Methods("DELETE")
}

//GetEggs returns all eggs as json
func (c *EggController) GetEggs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.repository.RetrieveEggs())
}

//GetEgg returns a single egg as json
func (c *EggController) GetEgg(w http.ResponseWriter, r *http.Request) {}

//CreateEgg will create an egg and insert it into the data storage after jittering its time
func (c *EggController) CreateEgg(w http.ResponseWriter, r *http.Request) {
	var egg models.Egg
	if err := json.NewDecoder(r.Body).Decode(&egg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	c.jitterEggTime(&egg)
	c.repository.StoreEgg(egg)
}

//DeleteEgg will delete and egg
func (c *EggController) DeleteEgg(w http.ResponseWriter, r *http.Request) {}

//JitterEggTime jitters the time the egg will hatch across a minute to lower clustering
func (c *EggController) jitterEggTime(egg *models.Egg) {
	egg.HatchTime = egg.HatchTime.Truncate(time.Minute)                  // make sure egg time starts at 0 seconds
	rand := rand.Intn(60) - 30                                           //get a value between -30 and 30 randomly
	egg.HatchTime = egg.HatchTime.Add(time.Second * time.Duration(rand)) //jitter the eggs hatch time 30 secnds either way
}
