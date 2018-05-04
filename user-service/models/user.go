package models

//User represents a user in our system
type User struct {
	Id                   int
	Name                 string
	Email                string
	NotificationEndpoint string
}
