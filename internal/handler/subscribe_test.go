package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/AndrewMislyuk/payments-api/internal/domain"
	"github.com/AndrewMislyuk/payments-api/internal/service"
	mock_service "github.com/AndrewMislyuk/payments-api/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_productSubscribe(t *testing.T) {
	type mockBehavior func(s *mock_service.MockPayments)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           domain.Product
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"id":"dca9128a-2c42-4495-ab4e-229c0a323328"}`,
			inputUser: domain.Product{
				Id: "dca9128a-2c42-4495-ab4e-229c0a323328",
			},
			mockBehavior: func(s *mock_service.MockPayments) {
				s.EXPECT().ProductSubscription("dca9128a-2c42-4495-ab4e-229c0a323328").Return("https://api.stripe.com/subscribe/dca9128a-2c42-4495-ab4e-229c0a323328", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"url":"https://api.stripe.com/subscribe/dca9128a-2c42-4495-ab4e-229c0a323328"}`,
		},

		{
			name:                "Empty Fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_service.MockPayments) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Key: 'Product.Id' Error:Field validation for 'Id' failed on the 'required' tag"}`,
		},

		{
			name:      "Service Failure",
			inputBody: `{"id":"dca9128a-2c42-4495-ab4e-229c0a323328"}`,
			inputUser: domain.Product{
				Id: "dca9128a-2c42-4495-ab4e-229c0a323328",
			},
			mockBehavior: func(s *mock_service.MockPayments) {
				s.EXPECT().ProductSubscription("dca9128a-2c42-4495-ab4e-229c0a323328").Return("", errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			payments := mock_service.NewMockPayments(c)
			testCase.mockBehavior(payments)

			services := &service.Service{Payments: payments}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/subscribe", handler.productSubscribe)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/subscribe", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
