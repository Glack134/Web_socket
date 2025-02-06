package service

import (
	"github.com/Glack134/web_socket"
	"github.com/Glack134/web_socket/pkg/repository"
)

type Authorization interface {
	CreateUser(user web_socket.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
