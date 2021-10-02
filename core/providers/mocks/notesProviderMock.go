package mocks

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/stretchr/testify/mock"
)

type NotesPostgresImplementationMock struct {
	mock.Mock
}

func (mock *NotesPostgresImplementationMock) Get(ctx context.Context, noteID string) (*entities.Note, error) {
	responseArgs := mock.Called(ctx, noteID)
	response := responseArgs.Get(0)
	err := responseArgs.Error(1)
	if response != nil {
		return response.(*entities.Note), err
	}

	return nil, err
}

func (mock *NotesPostgresImplementationMock) GetAll(ctx context.Context) (*[]entities.Note, error) {
	responseArgs := mock.Called(ctx)
	response := responseArgs.Get(0)
	err := responseArgs.Error(1)
	if err != nil {
		return nil, err
	}

	return response.(*[]entities.Note), nil
}

func (mock *NotesPostgresImplementationMock) Create(ctx context.Context, note *entities.Note) error {
	responseArgs := mock.Called(ctx, note)
	err := responseArgs.Error(0)
	if err != nil {
		return err
	}
	return nil
}

func (mock *NotesPostgresImplementationMock) Update(ctx context.Context, note *entities.ShortNote) error {
	responseArgs := mock.Called(ctx, note)
	err := responseArgs.Error(0)
	if err != nil {
		return err
	}
	return nil
}

func (mock *NotesPostgresImplementationMock) Delete(ctx context.Context, noteID string) error {
	responseArgs := mock.Called(ctx, noteID)
	err := responseArgs.Error(0)
	if err != nil {
		return err
	}
	return nil
}
