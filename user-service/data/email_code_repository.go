package data

import (
	"hatch/user-service/models"

	"github.com/google/uuid"
)

//EmailRepository interface provides functions to access email codes
type EmailRepository interface {
	GetEmailCode(code uuid.UUID) (models.EmailCode, error)
	AddEmailCode(userId uuid.UUID, code uuid.UUID) error
	RemoveEmailCode(code uuid.UUID) error
}
