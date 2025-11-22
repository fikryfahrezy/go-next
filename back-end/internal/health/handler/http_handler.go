package handler

import (
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/go-next/internal/database"
	"github.com/fikryfahrezy/go-next/internal/http_server"
)

type HealthHandler struct {
	db        *database.DB
	version   string
	commit    string
	buildTime string
}

func NewHealthHandler(db *database.DB, version, commit, buildTime string) *HealthHandler {
	return &HealthHandler{
		db:        db,
		version:   version,
		commit:    commit,
		buildTime: buildTime,
	}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	message := "Service is healthy"
	httpStatus := http.StatusOK

	dbCheck := map[string]any{"status": "unknown"}

	// Check database connection
	if err := h.db.Health(); err != nil {
		slog.Error("Database health check failed",
			slog.String("error", err.Error()),
		)
		status = "unhealthy"
		message = "Database connection failed"
		httpStatus = http.StatusServiceUnavailable
		dbCheck = map[string]any{
			"status": "unhealthy",
			"error":  err.Error(),
		}
	}

	response := map[string]any{
		"status":    status,
		"message":   message,
		"version":   h.version,
		"commit":    h.commit,
		"buildTime": h.buildTime,
		"checks": map[string]any{
			"database": dbCheck,
		},
	}

	http_server.JSON(w, httpStatus, response)
}

// SetupRoutes configures health check routes
func (h *HealthHandler) SetupRoutes(server *http_server.Server) {
	// Health check endpoint (no versioning needed)
	server.HandleFunc("GET /health", h.HealthCheck)
}
