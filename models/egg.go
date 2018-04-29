package models

import "time"

//Egg is the messages/reminders that user will create for themslves or other users
type Egg struct {
	Id        int       `json:"id"`
	To        string    `json:"to"`
	From      string    `json:"from"`
	Message   string    `json:"message"`
	Layed     time.Time `json:"layed"`
	HatchTime time.Time `json:"hatchTime"`
	Picture   string    `json:"picture"`
}
