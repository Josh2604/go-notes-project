package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/dto/requests"
	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/gin-gonic/gin"
)

type NoteCreate struct {
	Note usecases.NoteCreate
}

func (h *NoteCreate) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "create_note")

	inp := new(requests.NoteRequest)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	note := entities.NewNoteFromRequest(*inp)
	err := h.Note.Exec(ctx, &note)
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
	}

	c.JSON(http.StatusOK, new(interface{}))
}
