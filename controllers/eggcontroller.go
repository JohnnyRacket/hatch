package controllers

import (
	"net/http"
	"encoding/json"
	"hatchery/data"
	"hatchery/models"
)

func GetEggs(w http.ResponseWriter, r *http.Request)    {
	json.NewEncoder(w).Encode(data.RetrieveEggs())
}
func GetEgg(w http.ResponseWriter, r *http.Request)    {}
func CreateEgg(w http.ResponseWriter, r *http.Request) {
	var egg models.Egg
	if err := json.NewDecoder(r.Body).Decode(&egg); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	data.StoreEgg(egg);
}
func DeleteEgg(w http.ResponseWriter, r *http.Request) {}