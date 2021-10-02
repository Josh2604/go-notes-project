package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/providers/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateNoteProccessSuccess(t *testing.T) {

}

func TestCreateNoteProccessWithError(t *testing.T) {
	ctx := context.Background()

	testData := []struct {
		name        string
		note        entities.Note
		expectedErr interface{}
		err         bool
	}{
		{
			name: "Error when trying to create a note",
			note: entities.Note{
				Name:        "Test note",
				Description: "Test description",
			},
			expectedErr: errors.New("DB Error: creating record"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			notePostgresImplemetationMock := new(mocks.NotesPostgresImplementationMock)

			createNoteImplementation := CreateNoteImplementation{
				Note: notePostgresImplemetationMock,
			}

			notePostgresImplemetationMock.On("Create", ctx, &tt.note).Return(tt.expectedErr)

			err := createNoteImplementation.Exec(ctx, &tt.note)

			assert.NotNil(t, err)
			notePostgresImplemetationMock.AssertNumberOfCalls(t, "Create", 1)

		})
	}
}
