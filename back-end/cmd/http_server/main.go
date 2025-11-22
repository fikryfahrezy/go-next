// @title Go Simple API
// @description A simple API writte in Go
// @version 1.0
// @host localhost:8080
// @BasePath /api
// @schemes http https
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fikryfahrezy/go-next/config"
	_ "github.com/fikryfahrezy/go-next/docs"
	userHandler "github.com/fikryfahrezy/go-next/feature/user/handler"
	userRepository "github.com/fikryfahrezy/go-next/feature/user/repository"
	userService "github.com/fikryfahrezy/go-next/feature/user/service"
	"github.com/fikryfahrezy/go-next/internal/database"
	healthHandler "github.com/fikryfahrezy/go-next/internal/health/handler"
	server "github.com/fikryfahrezy/go-next/internal/http_server"
	"github.com/fikryfahrezy/go-next/internal/logger"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildTime = "unknown"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.Logger)

	log.Info("Starting application",
		slog.String("version", version),
		slog.String("commit", commit),
		slog.String("build_time", buildTime),
		slog.String("server_host", cfg.Server.Host),
		slog.Int("server_port", cfg.Server.Port),
	)

	db, err := database.NewDB()
	if err != nil {
		log.Error("Failed to connect to database",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// Create server configuration
	serverConfig := server.Config{
		Host: cfg.Server.Host,
		Port: cfg.Server.Port,
	}

	// Initialize feature dependencies
	userRepo := userRepository.NewUserRepository(log, db)
	userService := userService.NewUserService(log, userRepo)
	userHandlerInstance := userHandler.NewUserHandler(log, userService)

	// Initialize health handler
	healthHandlerInstance := healthHandler.NewHealthHandler(db, version, commit, buildTime)

	// Create and initialize server
	srv := server.New(serverConfig)

	routeHandlers := []server.RouteHandler{
		healthHandlerInstance,
		userHandlerInstance,
	}

	// Start server in goroutine
	go func() {
		if err := srv.Start(routeHandlers); err != nil {
			log.Error("Server error",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down gracefully...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop server
	if err := srv.Stop(ctx); err != nil {
		log.Error("Failed to shutdown server gracefully",
			slog.String("error", err.Error()),
		)
	}

	log.Info("Application shutdown complete")
}
