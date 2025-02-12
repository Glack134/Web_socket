package web_socket

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required, username"`
	Password string `json:"password" binding:"required, password"`
	Email    string `json:"email" binding:"required"`
}
