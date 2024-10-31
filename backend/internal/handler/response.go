package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rwx03/Pastebin/backend/pkg/logger"
)

type errorResponse struct {
	Message string `json:"mesage"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Log.Printf("Response Error: %s", message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
