package dependecies

import (
	"github.com/Josh2604/go-notes-project/core/providers/auth"
	"github.com/Josh2604/go-notes-project/core/providers/mongo"
	"github.com/Josh2604/go-notes-project/core/providers/postgres"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/Josh2604/go-notes-project/infraestructure/clients/mongodb"
	"github.com/Josh2604/go-notes-project/infraestructure/clients/postgresql"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint/handlers"
	"github.com/Josh2604/go-notes-project/infraestructure/entrypoint/middlewares"
	"github.com/Josh2604/go-notes-project/utils/zlog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HandlerContainer struct {
	NoteCreate gin.HandlerFunc
	NoteUpdate gin.HandlerFunc
	NoteGetAll gin.HandlerFunc
	NoteGet    gin.HandlerFunc
	NoteDelete gin.HandlerFunc
	SignUp     entrypoint.Handler
	SignIn     entrypoint.Handler
}

func Start() *HandlerContainer {
	dbConnection := mongodb.Start()
	pgConnection := postgresql.StartGorm()

	logger := zlog.New(true)

	// ---- REPOSITORIES ----
	user := &mongo.UserRepositoryImplementation{
		DB: dbConnection.Collection(viper.GetString("mongo.user_collection")),
	}

	postgres := &postgres.PostgresImplementation{
		DB: pgConnection,
	}
	/**
	MongoDB implementation, just substitute in usescases ej.
		noteCreate := &usecases.CreateNoteImplementation{
			Note: postgres,
		}
		to
		noteCreate := &usecases.CreateNoteImplementation{
			Note: notesMongoImpl,
		}
	*/
	// notesMongoImpl := mongo.NotesRepositoryImplementation{
	// 	Db: dbConnection.Collection(viper.GetString("mongo.notes_collection")),
	// }

	auth := auth.AuthRepositoryImplementation{
		UserRepo:       user,
		HashSalt:       viper.GetString("auth.hash_salt"),
		SigningKey:     []byte(viper.GetString("auth.signing_key")),
		ExpireDuration: viper.GetDuration("auth.token_ttl"),
	}

	// ---- USE CASES ----
	noteCreate := &usecases.CreateNoteImplementation{
		Note: postgres,
	}

	noteUpdate := &usecases.UpdateNoteImplementation{
		Note: postgres,
	}

	noteGet := &usecases.GetNoteImplementation{
		Note: postgres,
	}

	noteGetAll := &usecases.GetAllNotesImplementation{
		Note: postgres,
	}

	noteDelete := &usecases.DeleNoteImplementation{
		Note: postgres,
	}

	signUp := &usecases.SingUpImplementation{
		Auth: &auth,
	}

	signIn := &usecases.SignInImplentation{
		Auth: &auth,
	}

	// ---- HANDLERS ----
	handlersApp := HandlerContainer{}

	// Auth
	handlersApp.SignUp = &handlers.AuthSignUp{
		SignUp: signUp,
	}

	handlersApp.SignIn = &handlers.AuthSignIn{
		SignIn: signIn,
		Logger: logger,
	}

	// Notes handlers
	handlersApp.NoteCreate = middlewares.NewAuthMiddleware(
		&handlers.NoteCreate{
			Note:   noteCreate,
			Logger: logger,
		}, &auth,
	)

	handlersApp.NoteUpdate = middlewares.NewAuthMiddleware(
		&handlers.NoteUpdate{
			Note: noteUpdate,
		}, &auth,
	)

	handlersApp.NoteGetAll = middlewares.NewAuthMiddleware(
		&handlers.NoteGetAll{
			Note: noteGetAll,
		}, &auth,
	)

	handlersApp.NoteGet = middlewares.NewAuthMiddleware(
		&handlers.NoteGet{
			Note: noteGet,
		}, &auth,
	)

	handlersApp.NoteDelete = middlewares.NewAuthMiddleware(
		&handlers.NoteDelete{
			Note:   noteDelete,
			Logger: logger,
		}, &auth,
	)

	return &handlersApp
}
