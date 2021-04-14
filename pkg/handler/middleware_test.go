package handler

import (
	"errors"
	"fmt"
	"github.com/alextsa22/to-do-api/pkg/service"
	mock_service "github.com/alextsa22/to-do-api/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCod    int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCod:    200,
			expectedResponseBody: "1",
		},
		{
			name:                 "No Header",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCod:    401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "Invalid Number Of Header Parts",
			headerName:           "Authorization",
			headerValue:          "Bearer test token",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCod:    401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Invalid Bearer",
			headerName:           "Authorization",
			headerValue:          "Bearr token",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCod:    401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Invalid Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCod:    401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Service Failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("failed to parse token"))
			},
			expectedStatusCod:    401,
			expectedResponseBody: `{"message":"failed to parse token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCod, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestGetUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type testContext func() *gin.Context

	testTable := []struct {
		name          string
		context       testContext
		expectedId    int
		expectedError error
	}{
		{
			name: "OK",
			context: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				c.Set(userCtx, 1)
				return c
			},
			expectedId:    1,
			expectedError: nil,
		},
		{
			name: "No Key",
			context: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				return c
			},
			expectedId:    0,
			expectedError: errors.New("user id not found"),
		},
		{
			name: "Invalid Key Type",
			context: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				c.Set(userCtx, "test")
				return c
			},
			expectedId:    0,
			expectedError: errors.New("user id is of invalid type"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			id, err := getUserId(testCase.context())

			assert.Equal(t, testCase.expectedId, id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
