package handler

import (
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/go-next/internal/http_server"
	"github.com/google/uuid"
)

// GetUser retrieves a user by ID
// @Summary Get a user by ID
// @Description Retrieve a user by their unique identifier
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} http_server.APIResponse{result=service.GetUserResponse}
// @Failure 400 {object} http_server.APIResponse
// @Failure 404 {object} http_server.APIResponse
// @Failure 500 {object} http_server.APIResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.log.Warn("Invalid user ID parameter",
			slog.String("id", idParam),
		)
		http_server.BadRequestResponse(w, "Invalid user UUID format", err)
		return
	}

	user, err := h.userService.GetUserByID(r.Context(), id)
	if err != nil {
		h.translateServiceError(w, err, "Failed to get user")
		return
	}

	http_server.SuccessResponse(w, "User retrieved successfully", user)
}
