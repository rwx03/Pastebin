package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rwx03/Pastebin/backend/internal/models"
	"github.com/rwx03/Pastebin/backend/pkg/utils"
)

type PasteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *Handler) newPaste(c *gin.Context) {
	var input PasteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	uniqueID, err := utils.GenerateUniqueID()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	paste := models.Paste{
		PasteID:   uniqueID,
		Title:     input.Title,
		Content:   input.Content,
		CreatorID: 0,
		CreatedAt: time.Now(),
		Views:     0,
	}

	if _, err := h.services.Paste.Create(paste); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) getPaste(c *gin.Context) {
	pasteID := c.Query("id")

	if pasteID == "" {
		newErrorResponse(c, http.StatusBadRequest, "missing field id")
		return
	}

	paste, err := h.services.Paste.GetPasteByID(pasteID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"id":      paste.PasteID,
			"title":   paste.Title,
			"content": paste.Content,
		},
	)
}
