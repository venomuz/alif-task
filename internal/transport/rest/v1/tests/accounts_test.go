package tests

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/service"
	mock_service "github.com/venomuz/alif-task/internal/service/mocks"
	v1 "github.com/venomuz/alif-task/internal/transport/rest/v1"
	"github.com/venomuz/alif-task/pkg/logger"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_accountSingUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAccounts, user models.SignUpAccountInput)

	testTable := []struct {
		name                string
		inputBody           string
		inputSingUpAccount  models.SignUpAccountInput
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"birthday": "2011-01-11T00:00:00Z","lastName": "Farkhadov","name": "Davron","password": "admin123","phoneNumber": "998903456789","pinCode": 1111}`,
			inputSingUpAccount: models.SignUpAccountInput{
				Name:        "Davron",
				LastName:    "Farkhadov",
				PhoneNumber: "998903456789",
				Password:    "admin123",
				PinCode:     1111,
				Birthday: func() *time.Time {
					now, _ := time.Parse(time.RFC3339, "2011-01-11T00:00:00Z")
					return &now
				}(),
			},
			mockBehavior: func(r *mock_service.MockAccounts, input models.SignUpAccountInput) {
				ctx := context.Background()
				now, _ := time.Parse(time.RFC3339, "2011-01-11T00:00:00Z")
				r.EXPECT().SingUp(ctx, input).Return(models.AccountOut{
					ID:          uuid.MustParse("ea463003-b188-45e2-a033-6abd1c6e9148"),
					Name:        "Davron",
					LastName:    "Farkhadov",
					PhoneNumber: "998903456789",
					Password:    "admin123",
					Birthday:    &now,
					LastVisit:   nil,
					CreatedAt:   &now,
					UpdatedAt:   nil,
				}, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":"ea463003-b188-45e2-a033-6abd1c6e9148","name":"Davron","lastName":"Farkhadov","phoneNumber":"998903456789","password":"admin123","birthday":"2011-01-11T00:00:00Z","lastVisit":null,"createdAt":"2011-01-11T00:00:00Z","updatedAt":null}`,
		},
		{},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			account := mock_service.NewMockAccounts(c)

			testCase.mockBehavior(account, testCase.inputSingUpAccount)

			logger.New("debug", "app")

			cfg, err := config.Init("../../../../../configs")
			if err != nil {
				logger.Zap.Fatal("error while load configs", logger.Error(err))
				return
			}

			services := &service.Services{Accounts: account}

			handler := v1.NewHandler(services, cfg)

			g := gin.New()

			g.POST("/api/v1/accounts/sing-up", handler.AccountSingUp)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/api/v1/accounts/sing-up", bytes.NewBufferString(testCase.inputBody))

			g.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
