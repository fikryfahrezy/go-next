package repository

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	users, ok := r.db.V["users"]
	if !ok {
		return User{}, ErrFailedToGetUser
	}

	userValue, ok := users[id.String()]
	if !ok {
		return User{}, ErrUserNotFound
	}

	user, ok := userValue.(User)
	if !ok {
		r.log.Error("Failed to get user by ID",
			slog.String("user_id", id.String()),
		)
		return User{}, ErrFailedToGetUser
	}

	return user, nil
}
