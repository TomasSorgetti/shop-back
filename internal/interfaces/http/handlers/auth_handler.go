package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "Logged in")
}

func (h *AuthHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, "Registered")
}
