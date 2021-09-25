package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/repositories"
)

type DeleteNote interface {
	Exec(ctx context.Context, noteID string) error
}

type DeleNoteImplementation struct {
	Note repositories.NoteRepository
}

func (dn *DeleNoteImplementation) Exec(ctx context.Context, noteID string) error {
	err := dn.Note.Delete(ctx, noteID)
	if err != nil {
		return err
	}
	return nil
}
