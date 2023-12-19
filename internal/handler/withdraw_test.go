package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_getBalance(t *testing.T) {
	type mockBehaviour func(s *service.MockWithdraw, balance entity.Balance, user interface{})

	testTable := []struct {
		name                string
		user                interface{}
		balance             entity.Balance
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "Ok",
			user: 1,
			balance: entity.Balance{
				Money: 10000,
				Bonus: 1000,
			},
			mockBehaviour: func(s *service.MockWithdraw, balance entity.Balance, user interface{}) {
				s.EXPECT().GetBalance(user).Return(balance, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"current":100,"withdrawn":10}`,
		},
		{
			name: "Err to get balance",
			user: 1,
			balance: entity.Balance{
				Money: 10000,
				Bonus: 1000,
			},
			mockBehaviour: func(s *service.MockWithdraw, balance entity.Balance, user interface{}) {
				s.EXPECT().GetBalance(user).Return(entity.Balance{}, errors.New("err to get balance"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"err to get balance"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			withdraw := service.NewMockWithdraw(c)
			testCase.mockBehaviour(withdraw, testCase.balance, testCase.user)

			services := &service.Service{Withdraw: withdraw}
			handler := NewHandler(services)

			r := gin.Default()
			r.Use(func(c *gin.Context) {
				c.Set(userCtx, 1)
			})
			r.GET("/balance", handler.getBalance)

			rw := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/balance", nil)

			r.ServeHTTP(rw, req)

			assert.Equal(t, testCase.expectedStatusCode, rw.Code)
			assert.Equal(t, testCase.expectedRequestBody, rw.Body.String())
		})
	}
}
