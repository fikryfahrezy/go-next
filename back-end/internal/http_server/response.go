package http_server

import (
	"net/http"

	"github.com/fikryfahrezy/go-next/internal/app_error"
)

// APIResponse represents a standard API response
type APIResponse struct {
	Message     string         `json:"message"`
	Error       string         `json:"error"`
	ErrorFields map[string]any `json:"error_fields"`
	Result      any            `json:"result"`
}

// ListAPIResponse represents a paginated API response
type ListAPIResponse struct {
	Message     string              `json:"message"`
	Error       string              `json:"error"`
	ErrorFields map[string]any      `json:"error_fields"`
	Result      any                 `json:"result"`
	Pagination  *PaginationResponse `json:"pagination,omitempty"`
}

// PaginationResponse represents pagination metadata
type PaginationResponse struct {
	TotalData  int64 `json:"total_data"`
	TotalPages int64 `json:"total_pages"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
}

// PaginationRequest represents pagination input parameters
type PaginationRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func SuccessResponse(w http.ResponseWriter, message string, data any) {
	JSON(w, http.StatusOK, APIResponse{
		Message: message,
		Error:   "",
		Result:  data,
	})
}

func CreatedResponse(w http.ResponseWriter, message string, data any) {
	JSON(w, http.StatusCreated, APIResponse{
		Message: message,
		Error:   "",
		Result:  data,
	})
}

func ListSuccessResponse(w http.ResponseWriter, message string, data any, pagination PaginationResponse) {
	JSON(w, http.StatusOK, ListAPIResponse{
		Message:    message,
		Error:      "",
		Result:     data,
		Pagination: &pagination,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	errorCode := http.StatusText(statusCode)
	errorMessage := message
	var errorFields map[string]any

	if err != nil {
		// Extract error code if it's an AppError
		code := app_error.GetCode(err)
		if code != "" {
			errorCode = code
		}

		// Extract error message if it's an AppError
		message := app_error.GetMessage(err)
		if message != "" {
			errorMessage = message
		}

		// You can add error fields parsing here if needed
		errorFields = map[string]any{}
	}

	JSON(w, statusCode, APIResponse{
		Message:     errorMessage,
		Error:       errorCode,
		ErrorFields: errorFields,
	})
}

func BadRequestResponse(w http.ResponseWriter, message string, err error) {
	ErrorResponse(w, http.StatusBadRequest, message, err)
}

func NotFoundResponse(w http.ResponseWriter, message string, err error) {
	ErrorResponse(w, http.StatusNotFound, message, err)
}

func InternalServerErrorResponse(w http.ResponseWriter, message string, err error) {
	ErrorResponse(w, http.StatusInternalServerError, message, err)
}

// ValidationErrorResponse creates a response for validation errors with field-specific errors
func ValidationErrorResponse(w http.ResponseWriter, message string, errorFields map[string]any) {
	JSON(w, http.StatusUnprocessableEntity, APIResponse{
		Message:     message,
		Error:       http.StatusText(http.StatusUnprocessableEntity),
		ErrorFields: errorFields,
	})
}

// CreatePaginationResponse creates a pagination response object
func CreatePaginationResponse(totalData, totalPages int64, page, limit int) PaginationResponse {
	return PaginationResponse{
		TotalData:  totalData,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}
}
