package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/gin-gonic/gin"
)

type NoteDelete struct {
	Note usecases.DeleteNote
}

func (h *NoteDelete) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "delete_note")

	noteID := c.Param("id")
	if noteID == "" {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	err := h.Note.Exec(ctx, noteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	c.JSON(http.StatusOK, new(interface{}))
}
