package main

import (
	"log"
	"log/slog"
	"server/internal/config"
	"server/internal/logging"
	"server/internal/server"
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

	// server start
	err = server.ServerStartup(cfg.Server)
	if err != nil {
		return
	}
}
