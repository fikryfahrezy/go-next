package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fikryfahrezy/go-next/feature/user/handler"
	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/feature/user/service/servicefakes"
	"github.com/fikryfahrezy/go-next/internal/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_CreateUser_Success(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	userID := uuid.New()
	expectedResponse := service.CreateUserResponse{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockService.CreateUserReturns(expectedResponse, nil)

	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	requestBody := service.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	body, err := json.Marshal(requestBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CreateUser)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.CreateUserCallCount())
	_, actualReq := mockService.CreateUserArgsForCall(0)
	assert.Equal(t, requestBody.Name, actualReq.Name)
	assert.Equal(t, requestBody.Email, actualReq.Email)
	assert.Equal(t, requestBody.Password, actualReq.Password)
}

func TestUserHandler_CreateUser_ServiceError(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	mockService.CreateUserReturns(service.CreateUserResponse{}, service.ErrUserAlreadyExists)

	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	requestBody := service.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	body, err := json.Marshal(requestBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CreateUser)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.CreateUserCallCount())
}
