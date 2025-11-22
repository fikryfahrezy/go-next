package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

func (r *userRepository) Create(ctx context.Context, user User) error {
	users, ok := r.db.V["users"]
	if !ok {
		return ErrFailedToCreateUser
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Generate UUIDv7 for the user ID
	userID := uuid.Must(uuid.NewV7())
	user.ID = userID

	if err := r.db.BeginTx(); err != nil {
		r.log.Error("Failed to begin transaction", "error", err.Error())
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	users[userID.String()] = user

	if err := r.db.Commit(); err != nil {
		r.log.Error("Failed to commit transaction", "error", err.Error())
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	r.log.Info("User created successfully",
		slog.String("user_id", user.ID.String()),
		slog.String("email", user.Email),
	)

	return nil
}
