package models

import (
	"time"
)

//AccessToken represents a token granted to a user in our system
type AccessToken struct {
	UserId     int
	Token      string
	Expiration time.Time
}
