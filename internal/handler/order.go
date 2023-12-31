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

	id, errGet := c.Get(userCtx)
	if !errGet {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "can't get userID")
		return
	}

	errUserNumber := h.services.Order.CheckUserOrder(id.(int), string(body))
	if errUserNumber != nil {
		entity.NewErrorResponse(c, http.StatusOK, "order already accepted")
		return
	}

	errNumber := h.services.Order.CheckNumber(string(body))
	if errNumber != nil {
		entity.NewErrorResponse(c, http.StatusConflict, "number already exist")
		return
	}

	order := entity.Order{
		UserID: id.(int),
		Number: string(body),
		Status: "NEW",
	}

	err = h.services.Order.CreateOrder(order)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, "err to create order")
		return
	}

	responseOrder := ResponseOrder{
		Number:     order.Number,
		Status:     order.Status,
		Accrual:    float64(order.Accrual) / 100,
		UploadedAt: order.UploadedAt,
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
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

	orders := make([]ResponseOrder, 0)
	for _, v := range res {
		order := ResponseOrder{
			Number:     v.Number,
			Status:     v.Status,
			Accrual:    float64(v.Accrual) / 100,
			UploadedAt: v.UploadedAt,
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}
