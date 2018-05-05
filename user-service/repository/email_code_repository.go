package repository

import "github.com/google/uuid"

type EmailRepository interface {
	GetEmailCode(code uuid.UUID)
	AddEmailCode(email string, code uuid.UUID)
	RemoveEmailCode(code uuid.UUID)
}
