package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/repositories"
)

type NoteUpdate interface {
	Exec(ctx context.Context, note *entities.ShortNote) error
}

type UpdateNoteImplementation struct {
	Note repositories.NoteRepository
}

func (ev *UpdateNoteImplementation) Exec(ctx context.Context, note *entities.ShortNote) error {
	err := ev.Note.Update(ctx, note)
	if err != nil {
		return err
	}
	return nil
}
