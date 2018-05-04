package models

import (
	"time"
)

//RefreshToken represents a users session in our system, and their ability to refresh their access tokens
type RefreshToken struct {
	UserId     int
	Token      string
	Expiration time.Time
}
