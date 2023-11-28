package entity

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	NotUniqueLogin       = errors.New(`login is busy`)
	ErrToGenerateToken   = errors.New("error to generate token")
	InvalidLoginPassword = errors.New("invalid login or password")
	NotEnoughMoney       = errors.New("not enough money")
)

type errResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errResponse{message})
}
