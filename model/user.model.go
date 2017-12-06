package model

//User is a struct representing an app user
type User struct {
	ID             int `gorm:"primary_key"`
	Login          string
	PasswordHash   string
	FacebookUserID string
	VKUserID       string
}

//StandardAuth Auth info
type StandardAuth struct {
	Login    string
	Password string
}
