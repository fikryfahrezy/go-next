package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/internal/http_server"
)

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body service.CreateUserRequest true "User creation request"
// @Success 201 {object} http_server.APIResponse{result=service.GetUserResponse}
// @Failure 400 {object} http_server.APIResponse
// @Failure 500 {object} http_server.APIResponse
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req service.CreateUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.log.Error("Failed to bind request",
			slog.String("error", err.Error()),
		)
		http_server.BadRequestResponse(w, "Invalid request format", err)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		h.translateServiceError(w, err, "Failed to create user")
		return
	}

	http_server.CreatedResponse(w, "User created successfully", user)
}
