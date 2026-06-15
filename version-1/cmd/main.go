package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/carloscfgos1980/pizza-tracker/internal/models"
)

func main() {
	config := LoadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(config.DBPath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully", "db_path", config.DBPath)

	RegisterCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()

	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}
	slog.Info("Templates loaded successfully")

	setupRoutes(router, h)

	slog.Info("Serving starting", "url", "http://localhost:"+config.Port)
	if err := router.Run(":" + config.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
	router.Run(":" + config.Port)
}
