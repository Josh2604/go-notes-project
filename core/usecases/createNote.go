package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/repositories"
)

type NoteCreate interface {
	Exec(ctx context.Context, note *entities.Note) error
}

type CreateNoteImplementation struct {
	Note repositories.NoteRepository
}

func (ev *CreateNoteImplementation) Exec(ctx context.Context, note *entities.Note) error {
	err := ev.Note.Create(ctx, note)
	if err != nil {
		return err
	}

	return nil
}
