package controllers

import (
	"net/http"
	"polaris/internal/db"
	"polaris/internal/gateway/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Chat struct {
	ID       uint      `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}

func chatModelToResponse(model db.Chat) Chat {
	return Chat{
		ID:       model.ID,
		Title:    model.Title,
		Messages: messagesModelToResponse(model.Messages),
	}
}

// CreateChat godoc
// @Summary      Create a new chat
// @Description  Creates a new chat and returns its data
// @Tags         chats
// @Accept       json
// @Produce      json
// @Success      201  {object}  map[string]interface{}  "msg: created chat, chat: Chat"
// @Failure      500  {object}  map[string]any
// @Router       /chats [post]
func (s *Service) CreateChat(c *gin.Context) {
	chat, err := s.dbService.CreateChat(c.Request.Context())
	if err != nil {
		utils.InternalServerError(c, err, "failed create chat")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "created chat",
		"chat": chatModelToResponse(*chat),
	})
}

// GetChat godoc
// @Summary      Get a chat by ID
// @Description  Returns chat data, including messages, by its identifier
// @Tags         chats
// @Produce      json
// @Param        id   path      int  true  "Chat ID"
// @Success      200  {object}  Chat
// @Failure      400  {object}  map[string]any  "invalid path param"
// @Failure      500  {object}  map[string]any
// @Router       /chats/{id} [get]
func (s *Service) GetChat(c *gin.Context) {
	chatIDParam := c.Param("id")

	chatID, err := strconv.ParseUint(chatIDParam, 10, 64)
	if err != nil {
		utils.BadRequest(c, err, "path param")
		return
	}

	chat, err := s.dbService.GetChat(c.Request.Context(), uint(chatID))
	if err != nil {
		utils.InternalServerError(c, err, "failed get chat")
		return
	}

	c.JSON(http.StatusOK, chatModelToResponse(*chat))
}

type ListChatsRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit" binding:"required"`
}

// ListChats godoc
// @Summary      List chats
// @Description  Returns a paginated list of chats
// @Tags         chats
// @Produce      json
// @Param        offset  query     int  true  "Offset"
// @Param        limit   query     int  true  "Limit"
// @Success      200  {array}   Chat
// @Failure      400  {object}  map[string]any  "invalid query parameters"
// @Failure      500  {object}  map[string]any
// @Router       /chats [get]
func (s *Service) ListChats(c *gin.Context) {
	var request ListChatsRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		utils.BadRequest(c, err, "query")
		return
	}

	chats, err := s.dbService.ListChats(c.Request.Context(), request.Limit, request.Offset)
	if err != nil {
		utils.InternalServerError(c, err, "failed list chats")
		return
	}

	response := make([]Chat, 0, len(chats))
	for _, chat := range chats {
		response = append(response, chatModelToResponse(chat))
	}

	c.JSON(http.StatusOK, response)
}
