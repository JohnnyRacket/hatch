package repository

type UserRepository interface {
	GetUser()
	GetUsers()
	AddUser()
	RemoveUser()
}
