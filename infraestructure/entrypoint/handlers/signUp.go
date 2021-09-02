package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/dto/requests"
	"github.com/Josh2604/go-notes-project/core/dto/response"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/gin-gonic/gin"
)

type AuthSignUp struct {
	SignUp usecases.AuthSignUp
	// Logger *zlog.Logger
}

func (h *AuthSignUp) Handle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "SingUp User")

	inp := new(requests.UserSingInRequest)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, response.CustomMessage{
			Code:          http.StatusBadRequest,
			Message:       err.Error(),
			CustomMessage: "Error SingUp User",
		})
		return
	}

	if err := h.SignUp.Exec(ctx, inp.UserName, inp.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
