package logging

import (
	"io"
	"log"
	"log/slog"
	"os"
	"server/internal/config"
)

func ConfigureLogger(cfg config.LoggingConfig) error {

	var logger *slog.Logger
	var level slog.Level
	var output io.Writer = os.Stdout

	err := level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		log.Println("Error while configuring logging level")
		return err
	}

	switch cfg.Type {
	case config.TextLogger:
		{
			logger = slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
				AddSource: cfg.AddSource,
				Level:     level,
			}))
		}
	case config.JsonLogger:
		{
			logger = slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
				AddSource: cfg.AddSource,
				Level:     level,
			}))
		}
	default:
		{
			logger = slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
				AddSource: cfg.AddSource,
				Level:     level,
			}))
		}
	}

	slog.SetDefault(logger)
	return nil
}
