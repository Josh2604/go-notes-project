package server

import (
	"github.com/Josh2604/go-notes-project/infraestructure/dependecies"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine, handlers *dependecies.HandlerContainer) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
