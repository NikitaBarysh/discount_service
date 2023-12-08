package middleware

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/NikitaBarysh/discount_service.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehaviour func(s *service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(0, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "0",
		},
		{
			name:                 "No header",
			headerName:           "",
			token:                "token",
			mockBehaviour:        func(s *service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "Wrong header value",
			headerName:           "Authorization",
			headerValue:          "Bearer",
			mockBehaviour:        func(s *service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Invalid header value",
			headerName:           "Authorization",
			headerValue:          "Bear token",
			mockBehaviour:        func(s *service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Invalid token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehaviour:        func(s *service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Can't parse token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(0, errors.New("err"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"can't parse token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			//handler := handler2.NewHandler(services)
			middleware := NewMiddleware(services)

			r := gin.New()
			r.GET("/protected", middleware.UserIdentity, func(c *gin.Context) {
				id, _ := c.Get(UserCtx)
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			rw := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(rw, req)

			assert.Equal(t, testCase.expectedStatusCode, rw.Code)
			assert.Equal(t, testCase.expectedResponseBody, rw.Body.String())
		})
	}
}
