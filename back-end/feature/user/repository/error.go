package repository

import "github.com/fikryfahrezy/go-next/internal/app_error"

// Repository errors
var (
	// User not found errors
	ErrUserNotFound = app_error.New("USER-USER_NOT_FOUND", "user not found")

	// Database operation errors
	ErrFailedToCreateUser = app_error.New("USER-FAILED_TO_CREATE_USER", "failed to create user")
	ErrFailedToGetUser    = app_error.New("USER-FAILED_TO_GET_USER", "failed to get user")
)
