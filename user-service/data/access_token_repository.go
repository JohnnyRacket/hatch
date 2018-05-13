package data

import (
	"hatch/user-service/models"

	"github.com/google/uuid"
)

//AccessTokenRepository interface provides functions to access users acces tokens
type AccessTokenRepository interface {
	GetAccessToken(userId uuid.UUID) (models.AccessToken, error)
	AddAccessToken(userId uuid.UUID, token string) error
	UpdateAccessToken(userId uuid.UUID, token string) error
	RemoveAccessCode(userId uuid.UUID) error
}
