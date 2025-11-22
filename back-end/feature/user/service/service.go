package service

import (
	"log/slog"

	"github.com/fikryfahrezy/go-next/feature/user/repository"
)

type userService struct {
	userRepo repository.UserRepository
	log      *slog.Logger
}

func NewUserService(log *slog.Logger, userRepo repository.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
		log:      log,
	}
}
