package handler

import (
	"net/http"

	"github.com/Glack134/web_socket"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createChat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input web_socket.WebsocketChat
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllChat(c *gin.Context) {

}

func (h *Handler) getAllChatById(c *gin.Context) {

}
func (h *Handler) updateChat(c *gin.Context) {

}
func (h *Handler) deleteChat(c *gin.Context) {

}
