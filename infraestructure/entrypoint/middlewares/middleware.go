package middlewares

import (
	"net/http"
	"strings"

	"github.com/Josh2604/go-notes-project/core/apierrors"
	"github.com/Josh2604/go-notes-project/core/providers/auth"
	"github.com/Josh2604/go-notes-project/core/repositories"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	usecases entrypoint.Handler
	auth     auth.AuthRepositoryImplementation
}

func NewAuthMiddleware(usecase entrypoint.Handler, auth *auth.AuthRepositoryImplementation) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecases: usecase,
		auth:     *auth,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.auth.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusUnauthorized
		if err == apierrors.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatus(status)
		return
	}

	c.Set(repositories.CtxUserKey, user)
	m.usecases.Handle(c)

}
