package data

import (
	"hatch/user-service/models"

	"github.com/google/uuid"
)

//RefreshTokenRepository interface provides functions to Refresh users acces tokens
type RefreshTokenRepository interface {
	GetRefreshToken(userId uuid.UUID) (models.RefreshToken, error)
	AddRefreshToken(userId uuid.UUID, token string) error
	UpdateRefreshToken(userId uuid.UUID, token string) error
	RemoveRefreshCode(userId uuid.UUID) error
}
