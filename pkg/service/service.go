package service

import (
	"github.com/Glack134/web_socket"
	"github.com/Glack134/web_socket/pkg/repository"
)

type Authorization interface {
	CreateUser(user web_socket.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
type WedsocketChat interface {
	Create(UserId int, list web_socket.WebsocketChat) (int, error)
}

type Service struct {
	Authorization
	WedsocketChat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		WedsocketChat: NewWedsocketChatService(repos.WedsocketChat),
	}
}
