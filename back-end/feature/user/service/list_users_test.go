package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikryfahrezy/go-next/feature/user/repository"
	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/internal/database"
	"github.com/fikryfahrezy/go-next/internal/http_server"
	"github.com/fikryfahrezy/go-next/internal/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService_ListUsers_Success(t *testing.T) {
	db, err := database.NewDB()
	db.V["users"] = map[string]any{}

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	userRepo := repository.NewUserRepository(log, db)
	userService := service.NewUserService(log, userRepo)
	ctx := context.Background()

	users := db.V["users"]

	firstUser := repository.User{
		ID:        uuid.New(),
		Name:      "User 1",
		Email:     "user1@example.com",
		Password:  "hashedpassword1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users[firstUser.ID.String()] = firstUser

	secondUser := repository.User{
		ID:        uuid.New(),
		Name:      "User 2",
		Email:     "user2@example.com",
		Password:  "hashedpassword2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users[secondUser.ID.String()] = firstUser

	paginationReq := service.ListUsersRequest{
		PaginationRequest: http_server.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	}

	result, totalCount, err := userService.ListUsers(ctx, paginationReq)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(2), totalCount)

	// Verify first user
	assert.Equal(t, firstUser.ID, result[0].ID)
	assert.Equal(t, firstUser.Name, result[0].Name)
	assert.Equal(t, firstUser.Email, result[0].Email)
}

func TestUserService_ListUsers_WithCustomPagination(t *testing.T) {
	db, err := database.NewDB()
	db.V["users"] = map[string]any{}

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	userRepo := repository.NewUserRepository(log, db)
	userService := service.NewUserService(log, userRepo)
	ctx := context.Background()

	users := db.V["users"]

	firstUser := repository.User{
		ID:        uuid.New(),
		Name:      "User 1",
		Email:     "user1@example.com",
		Password:  "hashedpassword1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users[firstUser.ID.String()] = firstUser

	paginationReq := service.ListUsersRequest{
		PaginationRequest: http_server.PaginationRequest{
			Page:     3,
			PageSize: 5,
		},
	}

	result, totalCount, err := userService.ListUsers(ctx, paginationReq)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), totalCount)
}

func TestUserService_ListUsers_EmptyResult(t *testing.T) {
	db, err := database.NewDB()
	db.V["users"] = map[string]any{}

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	userRepo := repository.NewUserRepository(log, db)
	userService := service.NewUserService(log, userRepo)
	ctx := context.Background()

	paginationReq := service.ListUsersRequest{
		PaginationRequest: http_server.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	}

	result, totalCount, err := userService.ListUsers(ctx, paginationReq)

	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), totalCount)
}
