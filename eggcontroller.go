package main

import (
	"net/http"
	"encoding/json"
)

func GetEggs(w http.ResponseWriter, r *http.Request)    {
	json.NewEncoder(w).Encode(RetrieveEggs())
}
func GetEgg(w http.ResponseWriter, r *http.Request)    {}
func CreateEgg(w http.ResponseWriter, r *http.Request) {
	var egg Egg
	if err := json.NewDecoder(r.Body).Decode(&egg); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	StoreEgg(egg);
}
func DeleteEgg(w http.ResponseWriter, r *http.Request) {}