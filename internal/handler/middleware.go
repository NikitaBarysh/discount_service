package handler

import (
	"net/http"
	"strings"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[1] == "" {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "can't parse token")
		//c.Abort()
		return
	}

	c.Set(userCtx, userID)
}
