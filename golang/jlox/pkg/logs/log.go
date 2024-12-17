package logs

import (
	"log/slog"
	"os"
)

func CreateLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}
