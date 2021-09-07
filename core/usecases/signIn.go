package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/repositories"
)

type AuthSignIn interface {
	Exec(ctx context.Context, username, password string) (string, error)
}

type SignInImplentation struct {
	Auth repositories.AuthRepository
}

func (si *SignInImplentation) Exec(ctx context.Context, username, password string) (string, error) {
	token, err := si.Auth.SingIn(ctx, username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
