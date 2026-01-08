package logger

import (
	"log/slog"
	"os"
)

// Package-level shared logger
var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}
