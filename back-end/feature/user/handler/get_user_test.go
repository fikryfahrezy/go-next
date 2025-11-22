package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fikryfahrezy/go-next/feature/user/handler"
	"github.com/fikryfahrezy/go-next/feature/user/repository"
	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/feature/user/service/servicefakes"
	"github.com/fikryfahrezy/go-next/internal/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_GetUser_Success(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	userID := uuid.New()
	expectedResponse := service.GetUserResponse{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockService.GetUserByIDReturns(expectedResponse, nil)

	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/{id}", nil)
	req.SetPathValue("id", userID.String())
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetUser)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called with correct ID
	assert.Equal(t, 1, mockService.GetUserByIDCallCount())
	_, actualID := mockService.GetUserByIDArgsForCall(0)
	assert.Equal(t, userID, actualID)
}

func TestUserHandler_GetUser_InvalidUUID(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/invalid-uuid", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetUser)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Service should not be called on invalid UUID
	assert.Equal(t, 0, mockService.GetUserByIDCallCount())
}

func TestUserHandler_GetUser_NotFound(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	userID := uuid.New()
	mockService.GetUserByIDReturns(service.GetUserResponse{}, repository.ErrUserNotFound)

	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/{id}", nil)
	req.SetPathValue("id", userID.String())
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetUser)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.GetUserByIDCallCount())
}
