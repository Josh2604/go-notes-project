package mocks

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/stretchr/testify/mock"
)

type CreateNoteImplementationMock struct {
	mock.Mock
}

func (mock *CreateNoteImplementationMock) Exec(ctx context.Context, note *entities.Note) error {
	response := mock.Called(ctx, note)
	err := response.Error(1)
	return err
}
