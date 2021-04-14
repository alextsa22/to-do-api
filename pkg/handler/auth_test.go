package handler

import (
	"bytes"
	"errors"
	"github.com/alextsa22/to-do-api/pkg/model"
	"github.com/alextsa22/to-do-api/pkg/service"
	mock_service "github.com/alextsa22/to-do-api/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_singUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user model.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		expectedStatusCod   int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "username":"test", "password":"qwerty"}`,
			inputUser: model.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCod:   200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{"username":"test", "password":"qwerty"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user model.User) {},
			expectedStatusCod:   400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"Test", "username":"test", "password":"qwerty"}`,
			inputUser: model.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCod:   500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCod, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
