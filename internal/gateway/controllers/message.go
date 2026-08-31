package controllers

import (
	"net/http"
	"polaris/internal/db"
	"polaris/internal/gateway/utils"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ID      uint   `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	FromAI  bool   `json:"from_ai,omitempty"`
	ChatID  uint   `json:"chat_id,omitempty"`
}

func messageModelToResponse(model db.Message) Message {
	return Message{
		ID:      model.ID,
		Content: model.Content,
		FromAI:  model.FromAI,
		ChatID:  model.ChatID,
	}
}

func messagesModelToResponse(model []db.Message) []Message {
	messages := make([]Message, 0, len(model))
	for _, m := range model {
		messages = append(messages, messageModelToResponse(m))
	}
	return messages
}

type CreateMessageRequest struct {
	ChatID  uint   `json:"chat_id"`
	Content string `json:"content"`
	FromAI  bool   `json:"from_ai"`
}

// CreateMessage godoc
// @Summary      Create a new message
// @Description  Creates a new message within a chat and returns its data
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        request  body      CreateMessageRequest  true  "Message data"
// @Success      201  {object}  Message
// @Failure      400  {object}  map[string]any  "invalid request body"
// @Failure      500  {object}  map[string]any
// @Router       /messages [post]
func (s *Service) CreateMessage(c *gin.Context) {
	var request CreateMessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.BadRequest(c, err, "request body")
		return
	}

	message, err := s.dbService.CreateMessage(
		c.Request.Context(),
		request.ChatID,
		request.Content,
		request.FromAI,
	)
	if err != nil {
		utils.InternalServerError(c, err, "failed to create message")
		return
	}

	c.JSON(http.StatusCreated, messageModelToResponse(*message))
}
