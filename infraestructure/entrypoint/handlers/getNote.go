package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/gin-gonic/gin"
)

type NoteGet struct {
	Note usecases.NoteGet
}

func (h *NoteGet) Handle(c *gin.Context) {
	ctx := context.Context(c)

	noteID := c.Param("id")
	if noteID == "" {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	note, err := h.Note.Exec(ctx, noteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
	}

	c.JSON(http.StatusOK, note)
}
