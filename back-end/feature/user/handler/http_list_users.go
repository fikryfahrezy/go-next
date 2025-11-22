package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/internal/http_server"
)

// ListUsers retrieves a list of users with pagination
// @Summary List users
// @Description Retrieve a paginated list of users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Success 200 {object} http_server.ListAPIResponse{result=[]service.GetUserResponse}
// @Failure 500 {object} http_server.APIResponse
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pageParam := queryParams.Get("page")
	pageSizeParam := queryParams.Get("page_size")

	page := 1
	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}

	pageSize := 10
	if pageSizeParam != "" {
		if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	paginationReq := http_server.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}
	users, totalCount, err := h.userService.ListUsers(r.Context(), service.ListUsersRequest{
		PaginationRequest: paginationReq,
	})
	if err != nil {
		h.translateServiceError(w, err, "Failed to list users")
		return
	}

	totalPages := int64(math.Ceil(float64(totalCount) / float64(pageSize)))
	pagination := http_server.CreatePaginationResponse(totalCount, totalPages, page, pageSize)

	http_server.ListSuccessResponse(w, "Users retrieved successfully", users, pagination)
}
