package data

import (
	"hatch/user-service/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUser(id uuid.UUID) (models.User, error)
	GetUsers() ([]models.User, error)
	AddUser(email string, name string)
	RemoveUser(id uuid.UUID)
}
