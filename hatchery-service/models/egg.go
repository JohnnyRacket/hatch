package models

import "time"

//Egg is the messages/reminders that user will create for themslves or other users
type Egg struct {
	Id        int       `json:"id"`
	Target    string    `json:"target"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	Layed     time.Time `json:"layed"`
	HatchTime time.Time `json:"hatchTime"`
	Picture   string    `json:"picture"`
}
