package dependecies

import (
	"github.com/Josh2604/go-notes-project/core/providers/mongo"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/Josh2604/go-notes-project/infraestructure/clients/mongodb"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint/handlers"
	"github.com/spf13/viper"
)

type HandlerContainer struct {
	NoteCreate entrypoint.Handler
}

func Start() *HandlerContainer {
	dbConnection := mongodb.Start()

	// ---- REPOSITORIES ----
	notes := mongo.NotesRepositoryImplementation{
		Db: dbConnection.Collection(viper.GetString("mongo.notes_collection")),
	}

	// ---- USE CASES ----
	noteCreate := &usecases.CreateNoteImplementation{
		Note: &notes,
	}

	// ---- HANDLERS ----
	handlersApp := HandlerContainer{}

	// Notes handlers
	handlersApp.NoteCreate = &handlers.NoteCreate{
		Note: noteCreate,
	}

	return &handlersApp
}
