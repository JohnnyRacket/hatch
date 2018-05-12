package models

//NotificationDetail represents how we store the details of how to notify a user
type NotificationDetail struct {
	Auth     string
	P256SH   string
	Endpoint string
}
