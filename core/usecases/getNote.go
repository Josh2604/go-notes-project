package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/repositories"
)

type NoteGet interface {
	Exec(ctx context.Context, noteID string) (*entities.Note, error)
}

type GetNoteImplementation struct {
	Note repositories.NoteRepository
}

func (gn *GetNoteImplementation) Exec(ctx context.Context, noteID string) (*entities.Note, error) {
	note, err := gn.Note.Get(ctx, noteID)
	if err != nil {
		return nil, err
	}

	return note, nil
}
