package models

import (
	"time"
)

//EmailCode represents a magical code that will be sent to the user to register in our system
type EmailCode struct {
	UserId     int
	Code       string
	Expiration time.Time
}
