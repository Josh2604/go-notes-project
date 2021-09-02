package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/gin-gonic/gin"
)

type NoteGetAll struct {
	Note usecases.NoteGetAll
}

func (h *NoteGetAll) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "get_all_notes")

	allNotes, err := h.Note.Exec(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
	}

	c.JSON(http.StatusOK, allNotes)
}
