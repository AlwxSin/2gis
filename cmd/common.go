package cmd

import (
	"log/slog"
	"os"
)

type Type int

const (
	DefaultType Type = iota
	TextType
	JSONType
)

const serviceName = "rooms"

func SetDefaultLogger(loggerType Type, serviceName string, opts *slog.HandlerOptions) {
	var loggerHandler slog.Handler
	switch loggerType {
	case TextType:
		loggerHandler = slog.NewTextHandler(os.Stdout, opts)
	case JSONType, DefaultType:
		loggerHandler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(loggerHandler).With(slog.String("serviceName", serviceName))
	slog.SetDefault(logger)
}

func setupLogger() {
	loggerType := JSONType
	loggerOpts := &slog.HandlerOptions{}

	debug := os.Getenv("DEBUG") == "true"
	if debug {
		loggerType = TextType
		loggerOpts = &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	}

	SetDefaultLogger(loggerType, serviceName, loggerOpts)
}
