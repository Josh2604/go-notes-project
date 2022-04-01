package server

import (
	"github.com/Josh2604/go-notes-project/infraestructure/dependecies"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine, handlers *dependecies.HandlerContainer) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	group := router.Group("/api/v1/notes")
	group.POST("/", handlers.NoteCreate)
	group.PUT("/:id", handlers.NoteUpdate)
	group.GET("/:id", handlers.NoteGet)
	group.GET("/all", handlers.NoteGetAll)
	group.DELETE("/:id", handlers.NoteDelete)
	group.POST("/signup", handlers.SignUp.Handle)
	group.POST("/signin", handlers.SignIn.Handle)
}
