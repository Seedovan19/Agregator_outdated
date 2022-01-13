package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/service"
	mock_service "github.com/seedovan19/Agregator/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, user agregator.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            agregator.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{ "name": "Test", "email": "abc@mail.ru", "password": "qwerty"}`,
			inputUser: agregator.User{
				Name:     "Test",
				Email:    "abc@mail.ru",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user agregator.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty Fields",
			inputBody:            `{"name": "Test"}`,
			inputUser:            agregator.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user agregator.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(repo, testCase.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
