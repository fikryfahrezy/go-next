package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fikryfahrezy/go-next/feature/user/handler"
	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/feature/user/service/servicefakes"
	"github.com/fikryfahrezy/go-next/internal/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_ListUsers_Success(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	expectedUsers := []service.ListUsersResponse{
		{
			ID:        uuid.New(),
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockService.ListUsersReturns(expectedUsers, 2, nil)

	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ListUsers)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called with default pagination
	assert.Equal(t, 1, mockService.ListUsersCallCount())
	_, paginationReq := mockService.ListUsersArgsForCall(0)
	assert.Equal(t, 1, paginationReq.Page)
	assert.Equal(t, 10, paginationReq.PageSize)
}

func TestUserHandler_ListUsers_WithPagination(t *testing.T) {
	mockService := &servicefakes.FakeUserService{}
	expectedUsers := []service.ListUsersResponse{
		{
			ID:        uuid.New(),
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockService.ListUsersReturns(expectedUsers, 1, nil)

	userHandler := handler.NewUserHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users?page=2&page_size=5", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ListUsers)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called with custom pagination
	assert.Equal(t, 1, mockService.ListUsersCallCount())
	_, paginationReq := mockService.ListUsersArgsForCall(0)
	assert.Equal(t, 2, paginationReq.Page)
	assert.Equal(t, 5, paginationReq.PageSize)
}
