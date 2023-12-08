package middleware

import (
	"net/http"
	"strings"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	services *service.Service
}

func NewMiddleware(service *service.Service) *Middleware {
	return &Middleware{services: service}
}

const (
	authorizationHeader = "Authorization"
	UserCtx             = "user_id"
)

func (m *Middleware) UserIdentity(c *gin.Context) {
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

	userID, err := m.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "can't parse token")
		c.Abort()
		return
	}

	c.Set(UserCtx, userID)
}
