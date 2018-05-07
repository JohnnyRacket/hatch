package data

import "github.com/google/uuid"

//EmailRepository interface provides functions to access email codes
type EmailRepository interface {
	GetEmailCode(code uuid.UUID) error
	AddEmailCode(userId uuid.UUID, code uuid.UUID) error
	RemoveEmailCode(code uuid.UUID) error
}
