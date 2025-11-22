package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fikryfahrezy/go-next/feature/user/repository"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	s.log.Info("Creating new user",
		slog.String("email", req.Email),
		slog.String("name", req.Name),
	)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Failed to hash password",
			slog.String("error", err.Error()),
		)
		return CreateUserResponse{}, fmt.Errorf("%w: %w", ErrFailedToHashPassword, err)
	}

	user := repository.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return CreateUserResponse{}, err
	}

	response := ToCreateUserResponse(user)
	s.log.Info("User created successfully",
		slog.String("user_id", user.ID.String()),
		slog.String("email", user.Email),
	)

	return response, nil
}
