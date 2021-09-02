package repositories

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
)

const CtxUserKey = "auth_reader"

type AuthRepository interface {
	SingUp(ctx context.Context, username, password string) error
	SingIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entities.User, error)
}
