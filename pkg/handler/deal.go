package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllNewDeals(c *gin.Context) {
	fmt.Println("getAllNewDeals")
}

func (h *Handler) getOneDeal(c *gin.Context) {
	fmt.Println("getOneDeal")
}

func (h *Handler) createOneDeal(c *gin.Context) {
	fmt.Println("createOneDeal")
}

func (h *Handler) deleteOneDeal(c *gin.Context) {
	fmt.Println("deleteOneDeal")
}

func (h *Handler) updateOneDeal(c *gin.Context) {
	fmt.Println("updateOneDeal")
}

func (h *Handler) updateSomeDeals(c *gin.Context) {
	fmt.Println("updateSomeDeals")
}
