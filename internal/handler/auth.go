package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/gin-gonic/gin"
)

const userID = "user_id"

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	errValidate := h.services.Authorization.ValidateLogin(input)
	if errors.Is(errValidate, entity.ErrNotUniqueLogin) {
		entity.NewErrorResponse(c, http.StatusConflict, "create new login, this is busy")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "server error, can't do registration")
		return
	}

	token, err := h.services.Authorization.GenerateToken(id)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't generate token")
		return
	}

	c.Header("Authorization", "Bearer "+token)
	c.IndentedJSON(http.StatusOK, token)
}

func (h *Handler) signIn(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, errData := h.services.Authorization.CheckData(input)
	fmt.Println("err to check data: ", errData)
	if errData != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, "invalid login or password")
		return
	}

	token, err := h.services.Authorization.GenerateToken(id)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't generate token")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "logined",
		"token":  token,
	})
}
