package service

import (
	"github.com/Glack134/web_socket"
	"github.com/Glack134/web_socket/pkg/repository"
)

type WebSocketChatService struct {
	repo repository.WedsocketChat
}

func NewWedsocketChatService(repo repository.WedsocketChat) *WebSocketChatService {
	return &WebSocketChatService{repo: repo}
}

func (s *WebSocketChatService) Create(UserId int, list web_socket.WebsocketChat) (int, error) {
	return s.repo.Create(UserId, list)
}
