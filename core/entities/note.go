package entities

import (
	"time"

	"github.com/Josh2604/go-notes-project/core/dto/requests"
)

// Note -
type Note struct {
	ID          string
	Name        string
	Description string
	Deleted     bool
	DateCreated time.Time
	DateUpdated time.Time
	DateDeleted time.Time
}

type ShortNote struct {
	ID          string  `json:"_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Deleted     *bool   `json:"deleted"`
}

func NewNoteFromRequest(note requests.NoteRequest) Note {
	return Note{
		Name:        note.Name,
		Description: note.Description,
		Deleted:     false,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
		DateDeleted: time.Now(),
	}
}

func NewNoteToUpdate(note requests.UpdateNoteRequest) ShortNote {
	return ShortNote{
		Name:        note.Name,
		Description: note.Description,
		Deleted:     note.Deleted,
	}
}
