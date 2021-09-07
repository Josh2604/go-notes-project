package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/apierrors"
	"github.com/Josh2604/go-notes-project/core/dto/requests"
	"github.com/Josh2604/go-notes-project/core/dto/response"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/Josh2604/go-notes-project/utils/zlog"
	"github.com/gin-gonic/gin"
)

type AuthSignIn struct {
	SignIn usecases.AuthSignIn
	Logger *zlog.Logger
}

func (h *AuthSignIn) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "signin_user")

	inp := new(requests.UserSingInRequest)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.SignIn.Exec(c.Request.Context(), inp.UserName, inp.Password)
	if err != nil {
		if err == apierrors.ErrUserNotFound {
			h.Logger.Error("Error trying get token", err, zlog.Tags{"type": "token_error", "error": err})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response.SignInResponse{Token: token})
}
