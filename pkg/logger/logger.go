package logger

import (
	"log/slog"
	"lostx/go-fiber-hw/config"
	"lostx/go-fiber-hw/pkg/logger/tint"
	"os"
	"strings"
)

func New(cfg *config.LoggerConfig) *slog.Logger {
	var handler slog.Handler

	switch strings.ToLower(cfg.Format) {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.Level(cfg.Level),
		})
	default:
		handler = tint.NewHandler(os.Stderr, &tint.Options{
			AddSource: true,
			Level:     slog.Level(cfg.Level),
		})
	}

	return slog.New(handler)
}
