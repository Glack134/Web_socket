package repository

import (
	"github.com/Glack134/web_socket"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user web_socket.User) (int, error)
	GetUser(username, password string) (web_socket.User, error)
}

type WedsocketChat interface {
	Create(userId int, list web_socket.WebsocketChat) (int, error)
}

type Repository struct {
	Authorization
	WedsocketChat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		WedsocketChat: NewWedsocketChatPostgres(db),
	}
}
