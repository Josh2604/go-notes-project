package dependecies

import (
	"github.com/Josh2604/go-notes-project/core/providers/auth"
	"github.com/Josh2604/go-notes-project/core/providers/mongo"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/Josh2604/go-notes-project/infraestructure/clients/mongodb"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint/handlers"
	"github.com/spf13/viper"
)

type HandlerContainer struct {
	NoteCreate entrypoint.Handler
	NoteUpdate entrypoint.Handler
	NoteGetAll entrypoint.Handler
	SignUp     entrypoint.Handler
}

func Start() *HandlerContainer {
	dbConnection := mongodb.Start()

	// ---- REPOSITORIES ----
	user := &mongo.UserRepositoryImplementation{
		Db: dbConnection.Collection(viper.GetString("mongo.user_collection")),
	}

	notes := mongo.NotesRepositoryImplementation{
		Db: dbConnection.Collection(viper.GetString("mongo.notes_collection")),
	}

	auth := auth.AuthRepositoryImplementation{
		UserRepo:       user,
		HashSalt:       viper.GetString("auth.hash_salt"),
		SigningKey:     []byte(viper.GetString("auth.signing_key")),
		ExpireDurarion: viper.GetDuration("auth.token_ttl"),
	}

	// ---- USE CASES ----
	noteCreate := &usecases.CreateNoteImplementation{
		Note: &notes,
	}

	noteUpdate := &usecases.UpdateNoteImplementation{
		Note: &notes,
	}

	noteGetAll := &usecases.GetAllNotesImplemetation{
		Note: &notes,
	}

	signUp := &usecases.SingUpImplementation{
		Auth: &auth,
	}

	// ---- HANDLERS ----
	handlersApp := HandlerContainer{}

	// Auth
	handlersApp.SignUp = &handlers.AuthSignUp{
		SignUp: signUp,
	}

	// Notes handlers
	handlersApp.NoteCreate = &handlers.NoteCreate{
		Note: noteCreate,
	}

	handlersApp.NoteUpdate = &handlers.NoteUpdate{
		Note: noteUpdate,
	}

	handlersApp.NoteGetAll = &handlers.NoteGetAll{
		Note: noteGetAll,
	}

	return &handlersApp
}
