package sloghelper

import (
	"context"
	"os"

	"log/slog"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Fatal(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, slog.Level(slog.LevelError+4), msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	logger.ErrorContext(ctx, msg, args...)
}
