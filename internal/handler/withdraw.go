package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getBalance(c *gin.Context) {
	userLogin, errGet := c.Get(userLogin)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userLogin")
		return
	}

	userID, err := h.services.Authorization.GetUserIDByLogin(userLogin.(string))
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	balance, err := h.services.Withdraw.GetBalance(userID)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to get balance")
		return
	}

	responseBalance := entity.ResponseBalance{
		Current:  balance.Money,
		Withdraw: balance.Bonus,
	}

	c.JSON(http.StatusOK, responseBalance)
}

func (h *Handler) useWithdraw(c *gin.Context) {
	var withdraw entity.Withdraw

	err := c.BindJSON(&withdraw)

	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to read body")
		return
	}

	number, err := strconv.Atoi(withdraw.Number)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to conv number")
		return
	}

	errNum := h.services.Order.LuhnAlgorithm(number)
	if !errNum {
		entity.NewErrorResponse(c, http.StatusUnprocessableEntity, "err to pass LuhnAlgorithm")
		return
	}

	userLogin, errGet := c.Get(userLogin)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userLogin")
		return
	}

	userID, err := h.services.Authorization.GetUserIDByLogin(userLogin.(string))
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	err = h.services.Withdraw.SetWithdraw(withdraw, userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotEnoughMoney) {
			entity.NewErrorResponse(c, http.StatusPaymentRequired, "not enough money")
			return
		}
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to set withdraw")
		return
	}

	c.JSON(http.StatusOK, "successfully")

}

func (h *Handler) getWithdraw(c *gin.Context) {
	userLogin, errGet := c.Get(userLogin)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userLogin")
		return
	}

	userID, err := h.services.Authorization.GetUserIDByLogin(userLogin.(string))
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	withdraw, err := h.services.Withdraw.GetWithdraw(userID)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusNoContent, "history is empty")
		return
	}

	responseWithdraw := make([]entity.ResponseWithdraw, 0)

	for _, v := range withdraw {
		res := entity.ResponseWithdraw{
			OrderNumber: v.Number,
			Sum:         v.Sum,
			UploadedAt:  v.UploadedAt,
		}
		responseWithdraw = append(responseWithdraw, res)
	}

	c.JSON(http.StatusOK, responseWithdraw)
}
