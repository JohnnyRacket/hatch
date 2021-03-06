package models

import (
	"regexp"

	"github.com/google/uuid"
)

var rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")
var rxName = regexp.MustCompile("^[a-zA-Z][\\w|-]{0,17}$")

//User represents a user in our system
type User struct {
	Id                   uuid.UUID
	Name                 string
	Email                string
	NotificationEndpoint string
}

//ValidateUser validates that the user struct is not malformed
func (u *User) Validate() bool {

	if u.Email == "" || !rxEmail.MatchString(u.Email) {
		return false
	}
	if u.Name == "" || !rxName.MatchString(u.Name) {
		return false
	}
	return true
}
