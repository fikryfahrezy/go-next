package repository

import (
	"log/slog"

	"github.com/fikryfahrezy/go-next/internal/database"
)

type userRepository struct {
	db  *database.DB
	log *slog.Logger
}

func NewUserRepository(log *slog.Logger, db *database.DB) *userRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}
