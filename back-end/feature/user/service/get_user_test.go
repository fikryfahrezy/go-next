package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikryfahrezy/go-next/feature/user/repository"
	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/internal/database"
	"github.com/fikryfahrezy/go-next/internal/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserByID_Success(t *testing.T) {
	db, err := database.NewDB()
	db.V["users"] = map[string]any{}

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	userRepo := repository.NewUserRepository(log, db)
	userService := service.NewUserService(log, userRepo)
	ctx := context.Background()

	userID := uuid.New()
	expectedUser := repository.User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users := db.V["users"]
	users[userID.String()] = expectedUser

	result, err := userService.GetUserByID(ctx, userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Name, result.Name)
	assert.Equal(t, expectedUser.Email, result.Email)
	assert.Equal(t, expectedUser.CreatedAt, result.CreatedAt)
	assert.Equal(t, expectedUser.UpdatedAt, result.UpdatedAt)
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	db, err := database.NewDB()
	db.V["users"] = map[string]any{}

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	userRepo := repository.NewUserRepository(log, db)
	userService := service.NewUserService(log, userRepo)
	ctx := context.Background()

	userID := uuid.New()

	result, err := userService.GetUserByID(ctx, userID)

	assert.Error(t, err)
	assert.Equal(t, repository.ErrUserNotFound, err)
	assert.Equal(t, service.GetUserResponse{}, result)
}
