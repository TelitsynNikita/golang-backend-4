package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("signIn")
}

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("signUp")
}
