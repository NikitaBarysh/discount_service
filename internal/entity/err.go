package entity

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	ErrNotUniqueLogin       = errors.New(`login is busy`)
	ErrToGenerateToken      = errors.New("error to generate token")
	ErrInvalidLoginPassword = errors.New("invalid login or password")
	ErrNotEnoughMoney       = errors.New("not enough money")
	ErrTooManyRequest       = errors.New("too many request")
)

type errResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errResponse{message})
}
