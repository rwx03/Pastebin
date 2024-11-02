package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// sanitize Bearer prefix
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := h.services.Auth.ValidateToken(tokenString, true)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid or expired token")
		c.Abort()
		return
	}

	c.Next()
}
