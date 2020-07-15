package models

type User struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
}
