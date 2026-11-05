package main

import (
	"log"
	"log/slog"
	"server/internal/config"
	"server/internal/logging"
)

func main() {
	// config setup
	cfg := config.LoadConfig()
	log.Println("Configs successfully loaded")

	// logging setup
	err := logging.ConfigureLogger(cfg.Logger)
	if err != nil {
		slog.Error("Error while configuring logger: ", "err", err)
		return
	}
	slog.Debug("Logger successfully configured")
	// repository startup

	// server start

}
