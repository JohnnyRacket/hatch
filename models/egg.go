package models

import "time"

type Egg struct {
	Id int `json:"id"`
	To string `json:"to"`
	From  string `json:"from"`
	Message  string `json:"message"`
	Layed time.Time   `json:"layed"`
	HatchTime time.Time `json:"hatchTime"`
}