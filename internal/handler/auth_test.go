package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(s *service.MockAuthorization, user entity.User, token string)
	tokenAnswer := "test"

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           entity.User
		token               string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			token: "token",
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().ValidateLogin(user).Return(nil)
				s.EXPECT().CreateUser(user).Return(1, nil)
				s.EXPECT().GenerateToken(1).Return(tokenAnswer, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"user created","token":"test"}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{"login":"test"}`,
			mockBehaviour:       func(s *service.MockAuthorization, user entity.User, token string) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Server Error",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().ValidateLogin(user).Return(nil)
				s.EXPECT().CreateUser(user).Return(0, errors.New("service error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error, can't do registration"}`,
		},
		{
			name:      "Login is exist",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().ValidateLogin(user).Return(entity.ErrNotUniqueLogin)
			},
			expectedStatusCode:  409,
			expectedRequestBody: `{"message":"create new login, this is busy"}`,
		},
		{
			name:      "Generate token error",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			token: "token",
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().ValidateLogin(user).Return(nil)
				s.EXPECT().CreateUser(user).Return(1, nil)
				s.EXPECT().GenerateToken(1).Return("", entity.ErrToGenerateToken)
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"can't generate token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.Default()
			r.POST("/register", handler.signUp)

			rw := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/register",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(rw, req)

			assert.Equal(t, testCase.expectedStatusCode, rw.Code)
			assert.Equal(t, testCase.expectedRequestBody, rw.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehaviour func(s *service.MockAuthorization, user entity.User, token string)
	tokenAnswer := "tokenAnswer"

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		token                string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			token: "token",
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().CheckData(user).Return(1, nil)
				s.EXPECT().GenerateToken(1).Return(tokenAnswer, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"logined","token":"tokenAnswer"}`,
		},
		{
			name:                 "Empty body",
			inputBody:            `{"login":"test""}`,
			mockBehaviour:        func(s *service.MockAuthorization, user entity.User, token string) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Wrong data",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().CheckData(user).Return(0, entity.ErrInvalidLoginPassword)
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid login or password"}`,
		},
		{
			name:      "Generate token error",
			inputBody: `{"login":"test", "password":"qwerty"}`,
			inputUser: entity.User{
				Login:    "test",
				Password: "qwerty",
			},
			token: "token",
			mockBehaviour: func(s *service.MockAuthorization, user entity.User, token string) {
				s.EXPECT().CheckData(user).Return(1, nil)
				s.EXPECT().GenerateToken(1).Return("", entity.ErrToGenerateToken)
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"can't generate token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/login", handler.signIn)

			rw := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/login",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(rw, req)

			assert.Equal(t, testCase.expectedStatusCode, rw.Code)
			assert.Equal(t, testCase.expectedResponseBody, rw.Body.String())
		})
	}
}
