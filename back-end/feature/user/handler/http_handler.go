package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/go-next/feature/user/repository"
	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/internal/http_server"
)

type UserHandler struct {
	userService service.UserService
	log         *slog.Logger
}

func NewUserHandler(log *slog.Logger, userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log,
	}
}

// translateServiceError converts service errors to appropriate HTTP responses
func (h *UserHandler) translateServiceError(w http.ResponseWriter, err error, defaultMessage string) {
	if errors.Is(err, service.ErrUserAlreadyExists) {
		http_server.BadRequestResponse(w, "Email address is already taken", err)
	}
	if errors.Is(err, service.ErrFailedToHashPassword) {
		http_server.InternalServerErrorResponse(w, "Password processing failed", err)
	}
	if errors.Is(err, repository.ErrUserNotFound) {
		http_server.NotFoundResponse(w, "User not found", err)
	}

	// Log unexpected errors
	h.log.Error("Service error",
		slog.String("error", err.Error()),
		slog.String("operation", defaultMessage),
	)
	http_server.InternalServerErrorResponse(w, defaultMessage, err)
}

// SetupRoutes configures all versioned API routes for users
func (h *UserHandler) SetupRoutes(server *http_server.Server) {
	server.HandleFunc("POST /api/v1/users", h.CreateUser)
	server.HandleFunc("GET /api/v1/users", h.ListUsers)
	server.HandleFunc("GET /api/v1/users/detail/{id}", h.GetUser)
}
