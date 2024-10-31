package main

import (
	"fmt"

	"github.com/rwx03/Pastebin/backend/internal/handler"
	"github.com/rwx03/Pastebin/backend/internal/repository"
	"github.com/rwx03/Pastebin/backend/internal/services"
	"github.com/rwx03/Pastebin/backend/pkg/config"
	"github.com/rwx03/Pastebin/backend/pkg/logger"
)

func main() {
	cfg := config.GetConfig()

	pool, err := repository.NewPostgresDB(repository.Config{
		Host:         cfg.Database.Host,
		Port:         cfg.Database.Port,
		Username:     cfg.Database.Username,
		Password:     cfg.Database.Password,
		DatabaseName: cfg.Database.DatabaseName,
	})
	if err != nil {
		logger.Log.Fatalf("could not connect to database, error: %v", err)
	}

	repo := repository.NewRepository(pool)
	service := services.NewService(repo)

	fmt.Print(repo)

	h := handler.NewHandler(service)
	srv := h.InitRoutes()

	srv.Run(fmt.Sprintf("127.0.0.1:%d", cfg.Server.Port))
}
