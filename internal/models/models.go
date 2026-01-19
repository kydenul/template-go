package models

type User struct {
	Username string
	Password string `json:"-"`
	FullName string
	Email    string
}
