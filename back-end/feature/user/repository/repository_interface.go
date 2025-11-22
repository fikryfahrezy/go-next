package repository

//counterfeiter:generate -o repositoryfakes/fake_user_repository.go . UserRepository

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user User) error
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	List(ctx context.Context, limit, offset int) ([]User, int64, error)
}
