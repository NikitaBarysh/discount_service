package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_setOrder(t *testing.T) {
	type mockBehaviour func(s *service.MockOrder, order entity.Order, number string)

	testTable := []struct {
		name                string
		inputBody           string
		order               entity.Order
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Ok",
			inputBody: `378282246310005`,
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order, inputBody string) {
				num, _ := strconv.Atoi(order.Number)
				s.EXPECT().LuhnAlgorithm(num).Return(true)
				s.EXPECT().CheckUserOrder(1, order.Number).Return(nil)
				s.EXPECT().CheckNumber(order.Number).Return(nil)
				s.EXPECT().CreateOrder(order).Return(nil)
			},
			expectedStatusCode:  202,
			expectedRequestBody: `{"order":{"number":"378282246310005","status":"NEW","uploaded_at":"0001-01-01T00:00:00Z"}}`,
		},
		{
			name: "Err to convert to int ",
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order, inputBody string) {
				num, _ := strconv.Atoi(inputBody)
				fmt.Println(num)

			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"can't convert to int"}`,
		},
		{
			name:      "Err to pass Luhn algorithm",
			inputBody: `378282246310005`,
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order, inputBody string) {
				num, _ := strconv.Atoi(inputBody)
				s.EXPECT().LuhnAlgorithm(num).Return(false)
			},
			expectedStatusCode:  422,
			expectedRequestBody: `{"message":"don't pass luhn algorithm check"}`,
		},
		{
			name:      "Check if order already in progress",
			inputBody: `378282246310005`,
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order, inputBody string) {
				num, _ := strconv.Atoi(inputBody)
				s.EXPECT().LuhnAlgorithm(num).Return(true)
				s.EXPECT().CheckUserOrder(1, order.Number).Return(errors.New("already exist"))
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"message":"order already accepted"}`,
		},
		{
			name:      "Order already added by other user",
			inputBody: `378282246310005`,
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order, inputBody string) {
				num, _ := strconv.Atoi(inputBody)
				s.EXPECT().LuhnAlgorithm(num).Return(true)
				s.EXPECT().CheckUserOrder(1, order.Number).Return(nil)
				s.EXPECT().CheckNumber(order.Number).Return(errors.New("already exist"))
			},
			expectedStatusCode:  409,
			expectedRequestBody: `{"message":"number already exist"}`,
		},
		{
			name:      "Err to create order",
			inputBody: `378282246310005`,
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order, inputBody string) {
				num, _ := strconv.Atoi(inputBody)
				s.EXPECT().LuhnAlgorithm(num).Return(true)
				s.EXPECT().CheckUserOrder(1, order.Number).Return(nil)
				s.EXPECT().CheckNumber(order.Number).Return(nil)
				s.EXPECT().CreateOrder(order).Return(errors.New("err to create order"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"err to create order"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			order := service.NewMockOrder(c)
			testCase.mockBehaviour(order, testCase.order, testCase.inputBody)

			services := &service.Service{Order: order}
			handler := NewHandler(services)

			r := gin.Default()
			r.Use(func(c *gin.Context) {
				c.Set(userCtx, 1)
			})
			r.POST("/order", handler.setOrder)

			rw := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/order",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(rw, req)

			assert.Equal(t, testCase.expectedStatusCode, rw.Code)
			assert.Equal(t, testCase.expectedRequestBody, rw.Body.String())
		})
	}
}

func TestHandler_getOrders(t *testing.T) {
	type mockBehaviour func(s *service.MockOrder, order entity.Order)

	testTable := []struct {
		name                string
		order               entity.Order
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "Ok",
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order) {
				s.EXPECT().GetOrders(1).Return([]entity.Order{order}, nil)
				res := []entity.Order{order}
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
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"number":"378282246310005","status":"NEW","uploaded_at":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name: "Err get orders",
			order: entity.Order{
				UserID: 1,
				Number: `378282246310005`,
				Status: "NEW",
			},
			mockBehaviour: func(s *service.MockOrder, order entity.Order) {
				s.EXPECT().GetOrders(1).Return([]entity.Order{}, errors.New("err to get orders"))
			},
			expectedStatusCode: 204,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			order := service.NewMockOrder(c)
			testCase.mockBehaviour(order, testCase.order)

			services := &service.Service{Order: order}
			handler := NewHandler(services)

			r := gin.Default()
			r.Use(func(c *gin.Context) {
				c.Set(userCtx, 1)
			})
			r.GET("/order", handler.getOrders)

			rw := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order", nil)

			r.ServeHTTP(rw, req)

			assert.Equal(t, testCase.expectedStatusCode, rw.Code)
			assert.Equal(t, testCase.expectedRequestBody, rw.Body.String())
		})
	}
}
