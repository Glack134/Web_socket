package handler

import (
	"github.com/Glack134/web_socket/pkg/service"
	"github.com/Glack134/web_socket/pkg/websocket"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ws", websocket.WsPage)

	auth := router.Group("/login")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		chat := api.Group("/chat")
		{
			chat.POST("/", h.createChat)
			chat.GET("/", h.getAllChat)
			chat.GET("/:id", h.getAllChatById)
			chat.PUT("/:id", h.updateChat)
			chat.DELETE("/:id", h.deleteChat)
		}
	}
	// chat := router.Group("/chat")
	// {
	// 	chat.POST("/my_chat", h.MyChat)
	// 	chat.POST("/group_chat", h.GroupChat)
	// }
	return router
}
