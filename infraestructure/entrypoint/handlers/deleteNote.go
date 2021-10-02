package handlers

import (
	"context"
	"net/http"

	"github.com/Josh2604/go-notes-project/core/apierrors"
	"github.com/Josh2604/go-notes-project/core/apimessages"
	"github.com/Josh2604/go-notes-project/core/usecases"
	"github.com/Josh2604/go-notes-project/utils/zlog"
	"github.com/gin-gonic/gin"
)

type NoteDelete struct {
	Note   usecases.DeleteNote
	Logger *zlog.Logger
}

func (h *NoteDelete) Handle(c *gin.Context) {
	ctx := context.Context(c)
	ctx = context.WithValue(ctx, "action", "delete_note")

	noteID := c.Param("id")
	if noteID == "" {
		h.Logger.Error(apierrors.ErrParsingParam.Error(), apierrors.ErrParsingParam, zlog.Tags{})
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	err := h.Note.Exec(ctx, noteID)
	if err != nil {
		h.Logger.Error(apimessages.ErrorDeletingNote.GetMessage(), err, zlog.Tags{})
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	c.JSON(http.StatusOK, new(interface{}))
}
