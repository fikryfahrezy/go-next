package service

//counterfeiter:generate -o servicefakes/fake_user_service.go . UserService

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (GetUserResponse, error)
	ListUsers(ctx context.Context, req ListUsersRequest) ([]ListUsersResponse, int64, error)
}
