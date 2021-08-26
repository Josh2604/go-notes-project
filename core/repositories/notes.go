package repositories

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
)

type NoteRepository interface {
	Get(ctx context.Context, noteID string) (*entities.Note, error)
	GetAll(ctx context.Context) (*[]entities.Note, error)
	Create(ctx context.Context, note *entities.Note) error
	Update(ctx context.Context, note *entities.ShortNote) error
	Delete(ctx context.Context, noteID string) error
}
