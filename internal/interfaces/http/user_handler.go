// internal/interfaces/http/user_handler.go
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuhmanudheent/go-microservices/internal/usecase"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: u}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.useCase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
