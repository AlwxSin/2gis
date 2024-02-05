package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/rest"
)

func RunRestServer() {
	setupLogger()

	portStr := os.Getenv("SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		slog.Error("SERVER_PORT is not integer", "port", portStr)
		os.Exit(1)
	}

	s := rest.NewServer(&rest.ServerOptions{
		DB: models.NewInMemoryDB(),
	})

	slog.Info(fmt.Sprintf("Server listening on http://localhost:%d", port))
	err = s.ListenAndServe(port)
	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("Server closed")
	} else if err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
