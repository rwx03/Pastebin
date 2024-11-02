package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const refreshTokenTTL = 30 * 24 * 60 * 60 * 1000

func (h *Handler) register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	accessToken, refreshToken, err := h.services.Auth.Register(input.Email, input.Password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		newErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("error: %v", err))
		return
	}

	c.SetCookie("refreshToken", refreshToken, refreshTokenTTL, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"email":        input.Email,
	})
}

func (h *Handler) login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	accessToken, refreshToken, err := h.services.Auth.Login(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("error: %v", err))
		return
	}

	c.SetCookie("refreshToken", refreshToken, refreshTokenTTL, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"email":        input.Email,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	token, err := c.Cookie("refreshToken")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "no refresh token")
		return
	}

	accessToken, refreshToken, email, err := h.services.Auth.Refresh(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("error: %v", err))
		return
	}

	c.SetCookie("refreshToken", refreshToken, refreshTokenTTL, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"email":        email,
	})
}
