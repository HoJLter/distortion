package main

import (
	"log"
	"log/slog"
	"server/internal/config"
	"server/internal/handler"
	"server/internal/logging"
	"server/internal/repository"
)

func main() {
	// config setup
	cfg := config.LoadConfig()
	log.Println("Configs successfully loaded")

	// logging setup
	err := logging.ConfigureLogger(cfg.Logger)
	if err != nil {
		log.Println("Error while configuring logger")
		return
	}
	slog.Debug("Logger successfully configured")

	// repository startup

	db, err := repository.SetupRepositories(cfg.Database)
	if err != nil {
		return
	}
	defer db.Close()
	slog.Debug("Repository setup success")

	h := handler.NewHandler(db)

	// server start
	err = handler.RunServer(cfg.Server, h)
	if err != nil {
		return
	}
}
