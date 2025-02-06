package repository

import (
	"github.com/Glack134/web_socket"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user web_socket.User) (int, error)
	GetUser(username, password string) (web_socket.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
