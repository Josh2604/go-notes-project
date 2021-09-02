package server

import (
	"github.com/Josh2604/go-notes-project/infraestructure/dependecies"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine, handlers *dependecies.HandlerContainer) {
	group := router.Group("/api/v1/notes")

	group.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	group.POST("/create", handlers.NoteCreate.Handle)
	group.PUT("/update/:id", handlers.NoteUpdate.Handle)
	group.GET("/all", handlers.NoteGetAll.Handle)
	group.POST("/signup", handlers.SignUp.Handle)
}
