package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/apierrors"
	"github.com/Josh2604/go-notes-project/core/apimessages"
	"github.com/Josh2604/go-notes-project/core/dto/requests"
	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/Josh2604/go-notes-project/utils/zlog"
	"github.com/gin-gonic/gin"
)

type NoteCreate struct {
	Note   usecases.NoteCreate
	Logger *zlog.Logger
}

func (h *NoteCreate) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "create_note")

	inp := new(requests.NoteRequest)

	if err := c.BindJSON(inp); err != nil {
		h.Logger.Error(apimessages.ErrorBindingNote.GetMessage(), err, zlog.Tags{})
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequest())
		return
	}

	note := entities.NewNoteFromRequest(*inp)
	err := h.Note.Exec(ctx, &note)
	if err != nil {
		h.Logger.Error(apimessages.ErrorCreatingNote.GetMessage(), err, zlog.Tags{})
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	c.JSON(http.StatusCreated, apimessages.CreateSuccessMessage())
}
