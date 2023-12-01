package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setOrder(c *gin.Context) {
	defer c.Request.Body.Close()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	number, errConv := strconv.Atoi(string(body))
	if errConv != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, "can't convert to int")
		return
	}

	res := h.services.Order.LuhnAlgorithm(number)
	if !res {
		entity.NewErrorResponse(c, http.StatusUnprocessableEntity, "don't pass luhn algorithm check")
		return
	}

	errNumber := h.services.Order.CheckNumber(string(body))
	if errNumber != nil {
		entity.NewErrorResponse(c, http.StatusConflict, "number already exist")
		return
	}

	userID, errGet := c.Get(userCtx)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	order := entity.Order{
		UserID: userID.(int),
		Number: string(body),
		Status: "NEW",
	}

	err = h.services.Order.CreateOrder(order)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to create order")
		return
	}

	responseOrder := entity.ResponseOrder{
		Number:     order.Number,
		Status:     order.Status,
		Accrual:    order.Accrual,
		UploadedAt: order.UploadedAt,
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"order": responseOrder,
	})

}

func (h *Handler) getOrders(c *gin.Context) {
	userID, errGet := c.Get(userCtx)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	res, err := h.services.Order.GetOrders(userID.(int))
	if err != nil {
		entity.NewErrorResponse(c, http.StatusNoContent, "you don't have orders")
		return
	}

	orders := make([]entity.ResponseOrder, 0)
	for _, v := range res {
		order := entity.ResponseOrder{
			Number:     v.Number,
			Status:     v.Status,
			Accrual:    v.Accrual,
			UploadedAt: v.UploadedAt,
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}
