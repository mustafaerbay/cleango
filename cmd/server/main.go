package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mustafaerbay/cleango/internal/api"
	"github.com/mustafaerbay/cleango/internal/data/database"
	"github.com/mustafaerbay/cleango/internal/data/repository"
	"github.com/mustafaerbay/cleango/internal/biz/user"
	"github.com/mustafaerbay/cleango/pkg/config"
	"github.com/mustafaerbay/cleango/pkg/server"
)

func main() {
	// Load configuration values
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	// Initialize database connection
	db, err := database.NewMySQL(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}

	// Initialize user repository
	userRepo := repository.NewUserRepository(db)

	// Initialize user service
	userService := user.NewService(userRepo)

	// Initialize API handlers
	apiHandlers := api.NewHandlers(userService)

	// Initialize routes
	routes := api.Routes(apiHandlers)

	// Create server and start listening
	addr := fmt.Sprintf(":%s", cfg.Port)
	srv := server.NewServer(addr, routes)
	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
