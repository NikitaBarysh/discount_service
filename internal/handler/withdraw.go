package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getBalance(c *gin.Context) {
	id, errGet := c.Get(userCtx)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	balance, err := h.services.Withdraw.GetBalance(id.(int))
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to get balance")
		return
	}

	responseBalance := ResponseBalance{
		Current:  float64(balance.Money) / 100,
		Withdraw: float64(balance.Bonus) / 100,
	}

	c.JSON(http.StatusOK, responseBalance)
}

func (h *Handler) useWithdraw(c *gin.Context) {
	var withdraw entity.Withdraw

	err := c.BindJSON(&withdraw)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, "err to read ")
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

	id, errGet := c.Get(userCtx)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	err = h.services.Withdraw.SetWithdraw(withdraw, id.(int))
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
	id, errGet := c.Get(userCtx)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}

	withdraw, err := h.services.Withdraw.GetWithdraw(id.(int))
	if err != nil {
		entity.NewErrorResponse(c, http.StatusNoContent, "history is empty")
		return
	}

	responseWithdraw := make([]ResponseWithdraw, 0)

	for _, v := range withdraw {
		res := ResponseWithdraw{
			OrderNumber: v.Number,
			Sum:         float64(v.Sum) / 100,
			UploadedAt:  v.UploadedAt,
		}
		responseWithdraw = append(responseWithdraw, res)
	}

	c.JSON(http.StatusOK, responseWithdraw)
}
