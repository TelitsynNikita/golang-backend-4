package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxId           = "userId"
	userCtxRole         = "userRole"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorMessage(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorMessage(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, userRole, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtxId, userId)
	c.Set(userCtxRole, userRole)
}
