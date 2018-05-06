package models

import (
	"time"

	"github.com/google/uuid"
)

//EmailCode represents a magical code that will be sent to the user to register in our system
type EmailCode struct {
	UserId     uuid.UUID
	Code       string
	Expiration time.Time
}
