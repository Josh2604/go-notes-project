package usecases

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/repositories"
)

type NoteGetAll interface {
	Exec(ctx context.Context) (*[]entities.Note, error)
}

type GetAllNotesImplemetation struct {
	Note repositories.NoteRepository
}

func (gn *GetAllNotesImplemetation) Exec(ctx context.Context) (*[]entities.Note, error) {
	allNotes, err := gn.Note.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return allNotes, nil
}
