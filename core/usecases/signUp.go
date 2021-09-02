package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/repositories"
)

type AuthSignUp interface {
	Exec(ctx context.Context, username, password string) error
}

type SingUpImplementation struct {
	Auth repositories.AuthRepository
	// Logger *zlog.Logger
}

func (ev *SingUpImplementation) Exec(ctx context.Context, username, password string) error {

	err := ev.Auth.SingUp(ctx, username, password)
	if err != nil {
		return err
	}

	return nil
}
