package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/dto/requests"
	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/gin-gonic/gin"
)

type NoteUpdate struct {
	Note usecases.NoteUpdate
}

func (h *NoteUpdate) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "update_note")
	inp := new(requests.UpdateNoteRequest)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	noteID := c.Param("id")
	if noteID == "" {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	note := entities.NewNoteToUpdate(*inp)
	note.ID = noteID

	err := h.Note.Exec(ctx, &note)
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	c.JSON(http.StatusOK, new(interface{}))
}
