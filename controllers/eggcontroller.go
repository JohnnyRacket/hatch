package controllers

import (
	"encoding/json"
	"hatchery/data"
	"hatchery/models"
	"math/rand"
	"net/http"
	"time"
)

//TODO:  add wait groups to make it work async

//GetEggs returns all eggs as json
func GetEggs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.RetrieveEggs())
}

//GetEgg returns a single egg as json
func GetEgg(w http.ResponseWriter, r *http.Request) {}

//CreateEgg will create an egg and insert it into the data storage after jittering its time
func CreateEgg(w http.ResponseWriter, r *http.Request) {
	var egg models.Egg
	if err := json.NewDecoder(r.Body).Decode(&egg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jitterEggTime(&egg)
	data.StoreEgg(egg)
}

//DeleteEgg will delete and egg
func DeleteEgg(w http.ResponseWriter, r *http.Request) {}

//jitterEggTime jitters the time the egg will hatch across a minute to lower clustering
func jitterEggTime(egg *models.Egg) {
	egg.HatchTime.Truncate(time.Minute)                  // make sure egg time starts at 0 seconds
	rand := rand.Intn(60) - 30                           //get a value between -30 and 30 randomly
	egg.HatchTime.Add(time.Second * time.Duration(rand)) //jitter the eggs hatch time 30 secnds either way
}
